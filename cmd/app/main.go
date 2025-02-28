package main

import (
	"NewProjectGo/internal/database"
	"NewProjectGo/internal/handlers"
	"NewProjectGo/internal/taskService"
	"NewProjectGo/internal/web/tasks"
	"log"
	"github.com/labstack/echo/v4"
  	"github.com/labstack/echo/v4/middleware"
	
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	// инициализируем echo
	e := echo.New() 

	// используем Logger и Recover 
	e.Use(middleware.Logger())// что это ?
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil) 
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

	// router := mux.NewRouter()
	// router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	// router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	// router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	// router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	// http.ListenAndServe(":8080", router)
}