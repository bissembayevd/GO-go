package handlers

import (
	"Goprojects/ASS-1/internal/models"
	"Goprojects/ASS-1/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Done bool `json:"done"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var list []models.Task
	for _, t := range storage.Tasks {
		list = append(list, t)
	}
	respondJSON(w, http.StatusOK, list)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	task, ok := storage.Tasks[id]
	if !ok {
		respondError(w, http.StatusNotFound, "task not found")
		return
	}

	respondJSON(w, http.StatusOK, task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Title == "" {
		respondError(w, http.StatusBadRequest, "title is required")
		return
	}

	task := models.Task{
		ID:    storage.NextID,
		Title: req.Title,
		Done:  false,
	}

	storage.Tasks[storage.NextID] = task
	storage.NextID++

	respondJSON(w, http.StatusCreated, task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	task, ok := storage.Tasks[id]
	if !ok {
		respondError(w, http.StatusNotFound, "task not found")
		return
	}

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}

	task.Done = req.Done
	storage.Tasks[id] = task

	respondJSON(w, http.StatusOK, task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if _, ok := storage.Tasks[id]; !ok {
		respondError(w, http.StatusNotFound, "task not found")
		return
	}

	delete(storage.Tasks, id)
	respondJSON(w, http.StatusOK, map[string]string{"message": "task deleted"})
}
