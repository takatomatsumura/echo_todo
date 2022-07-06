package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/takatomatsumura/echo_todo/ent"
	"github.com/takatomatsumura/echo_todo/ent/todo"
	"github.com/takatomatsumura/echo_todo/ent/user"
)

var Client *ent.Client
var Ctx context.Context

func TodoCreate(c echo.Context) error {
	fmt.Println(c.FormValue("name"))
	title := c.FormValue("title")
	content := c.FormValue("content")
	ownerID := c.FormValue("owner")
	deadLineString := c.FormValue("deadline")
	deadLine, err := time.Parse("2006-01-02 15:04:05", deadLineString)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}

	image, err := c.FormFile("picture")
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}

	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("/media/images" + image.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	fmt.Println(image)
	fmt.Println(content)

	owners, err := Client.User.Query().Where(user.Myuuid(ownerID)).Only(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	todo, err := Client.Todo.Create().SetTitle(title).SetContent(content).SetDeadline(deadLine).SetImagePath(image.Filename).AddOwner(owners).Save(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	fmt.Println(todo)
	fmt.Println(c.FormParams())
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func TodoRetreive(c echo.Context) error {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	todo, err := Client.Todo.Query().Where(todo.ID(todoID)).WithOwner().Only(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	return c.JSON(http.StatusOK, map[string]*ent.Todo{"todo": todo})
}

func GetOverdueLength(c echo.Context) error {
	ownerID := c.Param("id")
	owners, err := Client.User.Query().Where(user.Myuuid(ownerID)).All(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	if len(owners) != 0 {
		owner := owners[0]
		todosLen, err := Client.User.QueryTodos(owner).Where(
			todo.DeadlineLT(time.Now()),
			todo.TodoComplete(false),
		).Count(Ctx)
		if err != nil {
			return fmt.Errorf("failed %w", err)
		}
		return c.JSON(http.StatusOK, map[string]int{"len": todosLen})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "user is not exist"})
	}
}

func GetTodoList(c echo.Context) error {
	boolString := c.QueryParam("complete")
	todoBool, err := strconv.ParseBool(boolString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
	ownerID := c.Param("id")
	owners, err := Client.User.Query().Where(user.Myuuid(ownerID)).All(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	if len(owners) != 0 {
		owner := owners[0]
		todos, err := Client.User.QueryTodos(owner).Where(todo.TodoComplete(todoBool)).Order(ent.Asc(todo.FieldDeadline)).WithOwner().All(Ctx)
		if err != nil {
			return fmt.Errorf("failed %w", err)
		}
		return c.JSONPretty(http.StatusOK, map[string][]*ent.Todo{"todoList": todos}, " ")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "user is not exist"})
	}
}

func ChangeBool(c echo.Context) error {
	TodoID, err := strconv.Atoi(c.Param("pk"))
	fmt.Println(TodoID)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	todo, err := Client.Todo.Query().Where(todo.ID(TodoID)).Only(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	err = Client.Todo.UpdateOneID(todo.ID).SetTodoComplete(!todo.TodoComplete).Exec(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func TodoUpdate(c echo.Context) error {
	TodoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed get id: %w", err)
	}
	title := c.FormValue("title")
	content := c.FormValue("content")
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	todoItem, err := Client.Todo.Query().Where(todo.ID(TodoID)).Only(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	err = todoItem.Update().SetTitle(title).SetContent(content).Exec(Ctx)
	fmt.Println(todoItem)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func TodoDelete(c echo.Context) error {
	TodoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	err = Client.Todo.DeleteOneID(TodoID).Exec(Ctx)
	if err != nil {
		return fmt.Errorf("failed %w", err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
