package handler

import (
	"errors"
	"net/http"

	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateRoom(c echo.Context) error {
	roomName := c.FormValue("name")
	if roomName == "" {
		return errors.New("you fuckin bitch ass")
	}

	userId := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).Id

	err := database.CreateRoom(roomName, userId)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
