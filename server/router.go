package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

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
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  static,
		HTML5: true,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
	}))

	store = sessions.NewCookieStore([]byte(sessionSecret))
	e.Use(session.Middleware(store))

	apiGroup := e.Group("/api")
	{
		apiGroup.POST("/signup", signup)
		apiGroup.POST("/login", login)

		usersGroup := apiGroup.Group("/users")
		{
			usersGroup.GET("/me", getMe)
		}

		memoGroup := apiGroup.Group("/memos")
		{
			memoGroup.POST("", postMemo)
			memoGroup.GET("", getAllMemos)
			memoGroup.GET("/:memoID", getMemoByID)
		}
	}

	return e.Start(addr)
}

func signup(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate hashed password")
	}

	user.HashedPassword = string(hashedPassword)

	err = createUser(c.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
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

	user.Password = ""
	user.HashedPassword = ""

	return c.JSON(http.StatusCreated, user)
}

func login(c echo.Context) error {
	var reqUser User
	err := c.Bind(&reqUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	user, err := getUserByName(c.Request().Context(), reqUser.Name)
	if errors.Is(err, errNoUser) {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid user")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(reqUser.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid user")
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

	user.Password = ""
	user.HashedPassword = ""

	return c.JSON(http.StatusOK, user)
}

func getMe(c echo.Context) error {
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	iUserID, ok := session.Values["userID"]
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "not logged in")
	}
	userID, ok := iUserID.(int)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user id")
	}

	iUserName, ok := session.Values["userName"]
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "not logged in")
	}
	userName, ok := iUserName.(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user name")
	}

	return c.JSON(http.StatusOK, User{
		ID:   userID,
		Name: userName,
	})
}

func postMemo(c echo.Context) error {
	var reqMemo Memo
	err := c.Bind(&reqMemo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	userID, ok := session.Values["userID"]
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "not logged in")
	}

	reqMemo.UserID, ok = userID.(int)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user id")
	}
	reqMemo.CreatedAt = time.Now()

	err = createMemo(c.Request().Context(), &reqMemo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create memo")
	}

	return c.JSON(http.StatusCreated, reqMemo)
}

func getAllMemos(c echo.Context) error {
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}

	iUserID, ok := session.Values["userID"]
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "not logged in")
	}
	userID, ok := iUserID.(int)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user id")
	}

	memos, err := getMemos(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get memos")
	}

	return c.JSON(http.StatusOK, memos)
}

func getMemoByID(c echo.Context) error {
	memoID := c.Param("memoID")

	memo, err := getMemo(c.Request().Context(), memoID)
	if errors.Is(err, errNoMemo) {
		return echo.NewHTTPError(http.StatusNotFound, "no memo")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get memo")
	}

	return c.JSON(http.StatusOK, memo)
}
