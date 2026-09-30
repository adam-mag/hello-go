package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}


var users = []User{
	{ID: 1, Name: "Adam", Role: "Backend Developer"},
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HELLO WORLD ! FROM SERVER MUX")
	})

	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("POST /users", createUser)

	fmt.Println("Server berjalan di port 8080...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}


func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}
