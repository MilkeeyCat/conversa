package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"sync"

	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/MilkeeyCat/conversa/views/components"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type ConnectionPool struct {
	sync.RWMutex
	connections map[*websocket.Conn]struct {
		Id int
	}
}

var connectionPool ConnectionPool = ConnectionPool{
	connections: make(map[*websocket.Conn]struct{ Id int }),
}

func sendToAll(author, msg string, authorId int) error {
	connectionPool.RLock()
	defer connectionPool.RUnlock()

	for connection, data := range connectionPool.connections {
		var buf bytes.Buffer
		ctx := context.TODO()
		components.SwappingMessage(author, msg, data.Id == authorId).Render(ctx, &buf)

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
		connectionPool.connections[ws] = struct{ Id int }{
			Id: c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).Id,
		}
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

			err = json.Unmarshal(msg, &data)
			if err != nil {
				c.Logger().Error(err)
			}

			claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
			user, err := database.FindUserById(claims.Id)
			if err != nil {
				c.Logger().Error(err)
			}

			if err != nil {
				c.Logger().Error(err)
			}

			err = database.CreateMessage(claims.Id, data.Message, -1)
			if err != nil {
				c.Logger().Error(err)
			}

			err = sendToAll(user.Name, data.Message, claims.Id)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
