package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type requestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	// DB.Where("is_done = ?", true).Find(&tasks) для фильтрации, найти только выполненные задачи
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody requestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody) // мы декодируем(переносим) данные из r
	// в адресс ячейки памяти requestBody, передали адресс, изменили в нем значение
	if err != nil {
		http.Error(w, "Неверный json", http.StatusBadRequest)
		return
	}
// создаем новую задачу
	task := Task{
		Task:   requestBody.Task,
		IsDone: requestBody.IsDone,
	}
// сохраняем задачу в бд
	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, "Ошибка при сохранении в БД", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "json успешно записан")
}

func PatchHandler(w http.ResponseWriter, r *http.Request){
//Получаем ID задачи из URL
vars := mux.Vars(r)
id := vars["id"]
// декодируем данные для обновления
var requestBody requestBody
err := json.NewDecoder(r.Body).Decode(&requestBody)
if err != nil{
	http.Error(w, "Неверный json", http.StatusBadRequest)
	return
}
//находим задачу по id
var task Task
result := DB.First(&task, id)
if result.Error != nil{
	http.Error(w, "Задача не найдена" , http.StatusNotFound)
	return
}
// обновляем только те понял, которые были переданы
updatedData := Task{}
if requestBody.Task != ""{
	updatedData.Task = requestBody.Task
}
if requestBody.IsDone{
	updatedData.IsDone = requestBody.IsDone
}
//обновнялем задачу в базе данных 
result = DB.Model(&task).Updates(updatedData)
if result.Error != nil{
	http.Error(w, "Ошибка при обновлении задачи", http.StatusInternalServerError)
}
fmt.Fprintf(w, "Задача обновлена")
}


func DeleteHandler(w http.ResponseWriter, r *http.Request){
	//Получаем ID задачи из URL и находим задачу по id
	vars := mux.Vars(r)
	id := vars["id"]
// Находим задачу по ID
	var task Task
result := DB.First(&task, id)
if result.Error != nil{
	http.Error(w, "Задача не найдена" , http.StatusNotFound)
	return
}
// удаляем задачу 
result = DB.Delete(&task)
if result.Error != nil {
	http.Error(w, "Ошибка при удалении задачи", http.StatusInternalServerError)
}
fmt.Fprintf(w, "Задача удалена")

}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})// миграция базы данных
	router := mux.NewRouter()
	router.HandleFunc("/post", PostHandler).Methods("POST")
	router.HandleFunc("/get", GetHandler).Methods("GET")
	router.HandleFunc("/task/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/task/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
