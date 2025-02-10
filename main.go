package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

var task string

type requestBody struct {
Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello", task)
}

func PostHandler(w http.ResponseWriter, r *http.Request){
	var requestBody requestBody
 	err := json.NewDecoder(r.Body).Decode(&requestBody)// мы декодируем(переносим) данные из r 
	// в адресс ячейки памяти requestBody, передали адресс, изменили в нем значение
	if err != nil {
		http.Error(w, "Неверный json", http.StatusBadRequest)
		return 
	}

	task = requestBody.Message //записываем то что передали в глобальную переменную 
	fmt.Fprintf(w, "json успешно записан")
}

func main(){
router := mux.NewRouter() 

router.HandleFunc("/post", PostHandler).Methods("POST")
router.HandleFunc("/get", GetHandler).Methods("GET")

http.ListenAndServe(":8080", router)

}