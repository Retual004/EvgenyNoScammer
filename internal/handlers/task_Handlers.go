package handlers

import (
	"NewProjectGo/internal/taskService"
	"NewProjectGo/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *taskService.TaskService
}



func NewTaskHandler(service *taskService.TaskService) *Handler {
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
func (h *Handler) PatchTasksTaskId(ctx context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	id := request.TaskId
	taskRequest := request.Body

	// Обновляем задачу через сервис
	updatedTask, err := h.Service.UpdateTaskByID(id, taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	})
	if err != nil {
		// Если ошибка при обновлении задачи, возвращаем ошибку
		return nil, err
	}

	// Возвращаем обновленную задачу в ответе
	// Здесь используем правильный тип ответа
	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем обновленную задачу в ответе
	return response, nil
}

// DeleteTask implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksTaskId(ctx context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	id := request.TaskId

	// Удаляем задачу через сервис
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		// Если задача не найдена или ошибка удаления, возвращаем ошибку
		return nil, err
	}

	// Возвращаем пустой ответ, так как задача успешно удалена
	return tasks.DeleteTasksTaskId204Response{}, nil
}
