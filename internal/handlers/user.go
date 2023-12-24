package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MilkeeyCat/conversa/internal/database"
	"github.com/MilkeeyCat/conversa/views/pages"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func verifyPassword(password, hash string) bool {
	passwordBytes := argon2.IDKey([]byte(password), []byte("salt? why not sugar?"), 1, 64*1024, 4, 32)
	base64Password := base64.RawStdEncoding.EncodeToString(passwordBytes)

	return base64Password == hash
}

func Register(c echo.Context) error {
	return render(c, pages.Register())
}

func RegisterPOST(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	if name == "" || password == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := database.FindUser(name)
	if err != nil {
		_, ok := err.(*database.UserNotFound)
		if !ok {
			return err
		}
	}

	if user.Name == name {
		//cant use the same name twice, do smth bout it
		return errors.New("nono")
	}

	passwordBytes := argon2.IDKey([]byte(password), []byte("salt? why not sugar?"), 1, 64*1024, 4, 32)
	base64Password := base64.RawStdEncoding.EncodeToString(passwordBytes)

	err = database.CreateUser(name, base64Password)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func Login(c echo.Context) error {
	return render(c, pages.Login())
}

func LoginPOST(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	if name == "" || password == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := database.FindUser(name)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	fmt.Println(user)

	if !verifyPassword(password, user.Password) {
		//paswords are not the same
		fmt.Println("passwords arent the same")
		return errors.New("fock you")
	}

	id, err := strconv.Atoi(user.Id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	claims := JwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   t,
		Expires: time.Now().Add(time.Hour * 72),
	})

	c.Redirect(http.StatusTemporaryRedirect, "/")

	return nil
}
