package handlers

import (
	"NewProjectGo/internal/taskService"
	"NewProjectGo/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}


func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// PatchTask implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	taskRequest := request.Body

	// Обновляем задачу через сервис
	updatedTask, err := h.Service.UpdateTaskByID(id, taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	})
	if err != nil {
		// Если ошибка при обновлении задачи, возвращаем ошибку
		return   nil , err 
	}

	// Возвращаем обновленную задачу в ответе
	// Здесь используем правильный тип ответа
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем обновленную задачу в ответе
	return response, nil
}

// DeleteTask implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	// Удаляем задачу через сервис
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		// Если задача не найдена или ошибка удаления, возвращаем ошибку
		return nil, err
	}

	// Возвращаем пустой ответ, так как задача успешно удалена
	return tasks.DeleteTasksId204Response{}, nil
}

//Нужна для создания структуры Handler на этапе инициализации приложения


// func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	idStr := vars["id"]
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		http.Error(w, "Invailed ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Декодируем данные для обновления
// 	var requestBody taskService.Task
// 	err = json.NewDecoder(r.Body).Decode(&requestBody)
// 	if err != nil {
// 		http.Error(w, "Invailed JSON", http.StatusBadRequest)
// 		return
// 	}

// 	updateTask, err := h.Service.UpdateTaskByID(uint(id), requestBody)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(updateTask)
// }

// func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	idStr := vars["id"]
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		http.Error(w, "Invailed ID", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.Service.DeleteTaskByID(uint(id))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Возвращаем статус 204 (No Content), так как тело ответа не нужно
// 	w.WriteHeader(http.StatusNoContent)
// }
