package handler

import (
	"errors"
	"strconv"

	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/MilkeeyCat/conversa/views/components"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Chat(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("nonono")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	userId := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).Id
	err = database.IsUserInRoom(idInt, userId)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	messages, err := database.GetRoomMessagesInRoom(idInt)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return render(c, components.Chat(messages, userId))
}
