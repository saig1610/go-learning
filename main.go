package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users []User
var nextID = 1

func createUser(w http.ResponseWriter, r *http.Request) {
    var u User
    json.NewDecoder(r.Body).Decode(&u)
    u.ID = nextID
    nextID++
    users = append(users, u)
    json.NewEncoder(w).Encode(u)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i := range users {
        if users[i].ID == id {
            json.NewDecoder(r.Body).Decode(&users[i])
            users[i].ID = id
            json.NewEncoder(w).Encode(users[i])
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i := range users {
        if users[i].ID == id {
            users = append(users[:i], users[i+1:]...)
            w.Write([]byte("User deleted"))
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users", getUsers).Methods("GET")
    r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}
