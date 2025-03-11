// internal/handlers/user_handler.go
package handlers

import (
	"NewProjectGo/internal/web/users"
	"context"
	"NewProjectGo/internal/userService"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// Получить всех пользователей
func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, u := range allUsers {
		user := users.User{
			Id:       &u.ID,
			Email:    &u.Email,
			Password: &u.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// Создать нового пользователя
func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	newUser := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

// Обновить пользователя
func (h *UserHandler) PatchUsersUserId(ctx context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	id := request.UserId
	userRequest := request.Body

	updatedUser, err := h.Service.UpdateUserByID(id, userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	})
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersUserId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

// Удалить пользователя
func (h *UserHandler) DeleteUsersUserId(ctx context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	id := request.UserId
	err := h.Service.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersUserId204Response{}, nil
}
