package main

import (
	"context"
	"log"

	"github.com/takatomatsumura/echo_todo/ent"
	"github.com/takatomatsumura/echo_todo/ent/migrate"
	"github.com/takatomatsumura/echo_todo/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	todoG := e.Group("/todos")
	accountsG := e.Group("/accounts")
	var err error
	handlers.Client, err = ent.Open("postgres", "host=localhost port=5432 user=takato dbname=takatodb password=thisistest sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	handlers.Ctx = context.Background()
	err = handlers.Client.Schema.Create(handlers.Ctx, migrate.WithDropColumn(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	todoG.POST("/create", handlers.TodoCreate)
	todoG.GET("/:id", handlers.TodoRetreive)
	todoG.GET("/todolist/len/:id", handlers.GetOverdueLength)
	todoG.GET("/todolist/:id", handlers.GetTodoList)
	todoG.PUT("/donebool/:pk", handlers.ChangeBool)
	todoG.PUT("/:id", handlers.TodoUpdate)
	todoG.DELETE("/:id", handlers.TodoDelete)

	accountsG.POST("/user", handlers.UserCreate)
	accountsG.GET("/user", handlers.UserList)
	accountsG.GET("/user/:uuid", handlers.UserReterive)
	accountsG.PUT("/user/:uuid", handlers.UsernameUpdate)

	e.Static("/media", "media")

	e.Logger.Fatal(e.Start(":8000"))
}
