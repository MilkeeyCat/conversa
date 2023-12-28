package handler

import (
	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/MilkeeyCat/conversa/views/pages"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	claims := c.Get("user")

	if claims != nil {
		data := claims.(*jwt.Token).Claims.(*JwtCustomClaims)
		messages, err := database.GetRoomMessagesInRoom(-1)
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		rooms, err := database.UserRooms(data.Id)
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		return render(c, pages.Authed(rooms, messages, data.Id))
	}

	return render(c, pages.Guest())
}
