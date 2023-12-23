package main

import (
	"fmt"
	"os"

	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	err := database.InitDB()

	if err != nil {
		panic(err)
	}

	app := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Start(fmt.Sprintf(":%s", port))
}
