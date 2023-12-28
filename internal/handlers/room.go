package handler

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func generateRoomToken() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, 15)

	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}

func CreateRoom(c echo.Context) error {
	roomName := c.FormValue("name")
	if roomName == "" {
		return errors.New("you fuckin bitch ass")
	}

	userId := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).Id
	roomToken := generateRoomToken()

	err := database.CreateRoom(userId, roomName, roomToken)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func JoinRoom(c echo.Context) error {
	roomToken := c.Param("token")

	userId := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).Id
	err := database.AddUserInRoom(roomToken, userId)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func RoomMessages(c echo.Context) error {
	roomToken := c.Param("token")

	room, err := database.FindRoomByToken(roomToken)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	_, err = database.GetRoomMessagesInRoom(room.Id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return nil
}
