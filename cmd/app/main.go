// main.go
package main

import (
	"NewProjectGo/internal/database"
	"NewProjectGo/internal/handlers"
	"NewProjectGo/internal/taskService"
	"NewProjectGo/internal/userService"
	"NewProjectGo/internal/web/tasks"
	"NewProjectGo/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Репозиторий и сервис для задач
	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Репозиторий и сервис для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализация echo
	e := echo.New()

	// Используем логгер и recover middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация маршрутов для задач
	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	// Регистрация маршрутов для пользователей
	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	// Стартуем сервер
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
