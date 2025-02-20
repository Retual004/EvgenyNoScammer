package handlers

import (
	"NewProjectGo/internal/taskService"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Handler struct{
	Service *taskService.TaskService
}

//Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler{
	return &Handler{
		Service : service,
	}
}

func (h *Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request){
	tasks, err := h.Service.GetAllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)// что такое err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"] 
	id, err := strconv.ParseUint(idStr, 10, 32) 
	if err != nil {
		http.Error(w, "Invailed ID", http.StatusBadRequest)
		return
	}

	// Декодируем данные для обновления
	var requestBody taskService.Task
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invailed JSON", http.StatusBadRequest)
		return
	}

	updateTask, err := h.Service.UpdateTaskByID(uint(id), requestBody)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invailed ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	return 
	}

// Возвращаем статус 204 (No Content), так как тело ответа не нужно
	w.WriteHeader(http.StatusNoContent)
}
