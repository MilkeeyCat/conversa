package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MilkeeyCat/conversa/internal/database"
	handler "github.com/MilkeeyCat/conversa/internal/handlers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("couldn't parse variabled from .env file: %s", err)
	}

	err = database.InitDB()
	if err != nil {
		panic(err)
	}

	app := echo.New()

	app.GET("/", handler.Index)
	app.GET("/login", handler.Login)
	app.POST("/login", handler.LoginPOST)
	app.GET("/register", handler.Register)
	app.POST("/register", handler.RegisterPOST)
	app.GET("/ws", handler.WebsocketsHander)

	secret := os.Getenv("SECRET")
	if secret == "" {
		panic("cound't find the secret for JWT")
	}

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:6969"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	//its ugly as fuck but hey, guess what, it works
	app.Use(echojwt.WithConfig(echojwt.Config{
		ContinueOnIgnoredError: true,
		SigningKey:             []byte(secret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handler.JwtCustomClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return nil
		},
		TokenLookupFuncs: []middleware.ValuesExtractor{
			func(c echo.Context) ([]string, error) {
				token, err := c.Cookie("token")
				if err != nil {
					return []string{}, err
				}

				return []string{token.Value}, nil
			},
		},
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Start(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
