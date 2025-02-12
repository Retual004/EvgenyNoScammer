package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
var task string

type requestBody struct {
Task string `json:"task"`
IsDone bool   `json:"is_done"`
}

func GetHandler(w http.ResponseWriter, r *http.Request){
	var tasks []Task
	 DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request){
	var requestBody requestBody
 	err := json.NewDecoder(r.Body).Decode(&requestBody)// мы декодируем(переносим) данные из r 
	// в адресс ячейки памяти requestBody, передали адресс, изменили в нем значение
	if err != nil {
		http.Error(w, "Неверный json", http.StatusBadRequest)
		return 
	}
	task := Task{
		Task: requestBody.Task, 
		IsDone: requestBody.IsDone,
	}
	result:= DB.Create(&task)
	if result.Error != nil{
		http.Error(w, "Ошибка при сохранении в БД", http.StatusInternalServerError)
        return
	}
	fmt.Fprintf(w, "json успешно записан")
}

func main(){
InitDB()
DB.AutoMigrate(&Task{})
router := mux.NewRouter() 
router.HandleFunc("/post", PostHandler).Methods("POST")
router.HandleFunc("/get", GetHandler).Methods("GET")
http.ListenAndServe(":8080", router)

}