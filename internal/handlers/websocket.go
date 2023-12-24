package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/MilkeeyCat/conversa/views/components"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type ConnectionPool struct {
	sync.RWMutex
	connections map[*websocket.Conn]struct{}
}

var connectionPool ConnectionPool = ConnectionPool{
	connections: make(map[*websocket.Conn]struct{}),
}

func sendToAll(author, msg string) error {
	connectionPool.RLock()
	defer connectionPool.RUnlock()

	for connection := range connectionPool.connections {
		var buf bytes.Buffer
		ctx := context.TODO()
		components.Message(author, msg).Render(ctx, &buf)

		if err := websocket.Message.Send(connection, buf.String()); err != nil {
			return err
		}
	}

	return nil
}

func WebsocketsHander(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		defer func(connection *websocket.Conn) {
			connectionPool.Lock()
			delete(connectionPool.connections, connection)
			connectionPool.Unlock()
		}(ws)

		connectionPool.Lock()
		connectionPool.connections[ws] = struct{}{}
		connectionPool.Unlock()

		for {
			var msg []byte
			err := websocket.Message.Receive(ws, &msg)
			if err == io.EOF {
				break
			} else if err != nil {
				c.Logger().Error(err)
			}

			type WsRequest struct {
				Message string `json:"message"`
			}

			var data WsRequest
			fmt.Println(string(msg))

			err = json.Unmarshal(msg, &data)
			if err != nil {
				c.Logger().Error(err)
			}

			fmt.Println(data)

			name, err := c.Cookie("user")
			if err != nil {
				c.Logger().Error(err)
			}

			err = sendToAll(name.Value, data.Message)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
