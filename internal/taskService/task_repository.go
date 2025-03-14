package taskService

 import (
	"gorm.io/gorm"
	"fmt"
)

type TaskRepository interface {
	// CreateTask - передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTask - возвращаем массив из всех задач в БД и ошибку
	GetAllTask() ([]Task, error)
	// UpdateTaskByID - передаем id и Task, возвращаем обновленный
	// Task и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - передаем id для управления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
	// 
	GetTasksUserUserId(id uint) ([]Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository{
	return &taskRepository{db : db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error){
	if task.UserID == 0 {
		return Task{}, fmt.Errorf("user_id cannot be zero")
	}

	result := r.db.Create(&task)
	if result.Error != nil{
		return Task{}, result.Error
	}
	return task, nil 
}

func (r *taskRepository) GetAllTask() ([]Task, error){
	var tasks []Task
err := r.db.Find(&tasks).Error
return tasks, err 
}

func (r *taskRepository) UpdateTaskByID(id uint, requestBody Task) (Task, error){
	var task Task
result := r.db.First(&task, id)
if result.Error != nil{
	return task, result.Error
}

if requestBody.Task != "" {
	task.Task = requestBody.Task
}
if requestBody.IsDone != task.IsDone { 
	task.IsDone = requestBody.IsDone
}

result = r.db.Save(&task)
if result.Error != nil{
	return task  , result.Error
}
return task, nil 
}

func (r *taskRepository) DeleteTaskByID(id uint) error{
	var task Task
result := r.db.First(&task, id)
if result.Error != nil{
	return result.Error
}
result = r.db.Delete(&task)
if result.Error != nil {
	return result.Error
}
return nil
}

func (r *taskRepository) GetTasksUserUserId(userID uint) ([]Task, error) {
	var tasks []Task
	// Предполагаем, что у задачи есть поле UserID, которое связывает задачу с пользователем
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}