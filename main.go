package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	name   string
	email  string
	alamat string
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "Get" {
		users := []User{
			User{"wahyu", "wahyu@gmail.com", "Medan"},
			User{"fadil", "padil@gmail.com", "Jakarta"},
		}
		resultJSON, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Gagal Convert", http.StatusInternalServerError)
			return
		}
		w.Write((resultJSON))
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/v1/users", GetUserController)
	fmt.Println("Server Rest Api telah berjalan di port 8000")
	http.ListenAndServe(":8000", nil)
}
