package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takatomatsumura/echo_todo/ent"
	"github.com/takatomatsumura/echo_todo/ent/user"
)

func UserCreate(c echo.Context) error {
	name := c.FormValue("name")
	myuuid := c.FormValue("uuid")
	users, err := Client.User.Query().Where(user.Myuuid(myuuid)).Count(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	if users == 0 {
		user, err := Client.User.Create().SetName(name).SetMyuuid(myuuid).Save(Ctx)
		if err != nil {
			return fmt.Errorf("failed: %w", err)
		}
		fmt.Println(user)
		return c.JSON(http.StatusOK, map[string]string{"message": "success"})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "this user is already registerd"})
	}
}

func UserList(c echo.Context) error {
	users, err := Client.User.Query().All(Ctx)
	if err != nil {
		return fmt.Errorf("failed fetch user: %w", err)
	}
	return c.JSONPretty(http.StatusOK, map[string][]*ent.User{"users": users}, " ")
}

func UserReterive(c echo.Context) error {
	userID := c.Param("uuid")
	user, err := Client.User.Query().Where(user.Myuuid(userID)).Only(Ctx)
	if err != nil {
		return fmt.Errorf("failed fetch user: %w", err)
	}
	fmt.Println(user)
	return c.JSONPretty(http.StatusOK, map[string]*ent.User{"user": user}, " ")
}

func UsernameUpdate(c echo.Context) error {
	userID := c.Param("uuid")
	userName := c.FormValue("name")
	err := Client.User.Update().Where(user.Myuuid(userID)).SetName(userName).Exec(Ctx)
	if err != nil {
		return fmt.Errorf("failed fetch user: %w", err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
