package handler

import (
	"github.com/MilkeeyCat/conversa/views/pages"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	isAuthed := c.Get("user") != nil

	if isAuthed {
		return render(c, pages.Authed())
	}

	return render(c, pages.Guest())
}
