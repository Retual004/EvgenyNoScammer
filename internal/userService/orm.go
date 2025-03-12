package userService

import (
    "gorm.io/gorm"
    "NewProjectGo/internal/taskService"
)

type User struct {
    gorm.Model
    Email    string `json:"email"`
    Password string `json:"password"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"`
    Tasks []taskService.Task `gorm:"ForeignKey:UserID"`
}