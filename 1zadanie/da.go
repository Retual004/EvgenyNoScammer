package zadanie
import (
	"fmt"
	"net/http"
	"strconv"
)


var counter int 

func GetHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method == http.MethodGet{
	fmt.Fprintln(w, "counter равен", strconv.Itoa(counter))
 } else{
	fmt.Fprintln(w, "поддерживается только метод GET")
 }
}

func PostHandler(w http.ResponseWriter, r *http.Request){ 
if r.Method == http.MethodPost{
	counter++
	fmt.Fprintln(w, "Counter увеличен на один")
} else{
	fmt.Fprintf(w, "Поддерживается только метод POST")
}

}

func main(){
	http.HandleFunc("/get" , GetHandler) 
	http.HandleFunc("/post", PostHandler)	
	
	http.ListenAndServe("localhost:8080", nil)
	
	// fmt.Println("hello world")
}