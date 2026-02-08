package main

import (
	"Goprojects/ASS-1/internal/handlers"
	"Goprojects/ASS-1/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", handlers.GetTasks)
	mux.HandleFunc("GET /tasks/{id}", handlers.GetTaskByID)
	mux.HandleFunc("POST /tasks", handlers.CreateTask)
	mux.HandleFunc("PATCH /tasks/{id}", handlers.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handlers.DeleteTask)

	handler := middleware.Logging(middleware.APIKey(mux))

	fmt.Println("🚀 Server started on :8080")
	http.ListenAndServe(":8080", handler)
}
