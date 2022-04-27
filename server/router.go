package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

var store sessions.Store

func startServer(addr string, static string, sessionSecret string) error {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Static(static))

	store := sessions.NewCookieStore([]byte(sessionSecret))
	e.Use(session.Middleware(store))

	e.POST("/signup", signup)
	e.POST("/login", login)

	return e.Start(addr)
}

func signup(c echo.Context) error {
	var reqUser User
	err := c.Bind(&reqUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate hashed password")
	}

	reqUser.ID = uuid.New().String()
	reqUser.HashedPassword = string(hashedPassword)

	err = createUser(c.Request().Context(), &reqUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	return c.JSON(http.StatusCreated, reqUser)
}

func login(c echo.Context) error {
	var reqUser User
	err := c.Bind(&reqUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	user, err := getUserByName(c.Request().Context(), reqUser.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(reqUser.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	session.Values["userID"] = user.ID
	session.Values["userName"] = user.Name

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to save session")
	}

	return c.NoContent(http.StatusOK)
}
