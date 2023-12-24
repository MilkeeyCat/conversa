package handler

import (
	"github.com/MilkeeyCat/conversa/views/pages"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return render(c, pages.Index(c.Get("authed").(bool)))
}
