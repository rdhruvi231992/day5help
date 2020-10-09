package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRecord struct {
	ID    string
	Email string
}

func main() {

	conn, err := sqlx.Connect(
		"postgres",
		"postgres://rdhruvi23:Mayoor20@localhost:5432/imagio?sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)
		return

	}

	err = conn.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	UserStore = []*User{}

	// Create a new Router
	r := chi.NewRouter()

	// Use the middleware longger on each request
	r.Use(middleware.Logger)

	// Declare your routes
	r.Get("/", helloHandler)

	// State that the server is running
	fmt.Println("Running on :8080")
	r.Get("/api/users", listUserHandler)
	r.Post("/api/users/create", createUserHandler)
	r.Get("/api/users", readUserHandler)
	r.Post("/api/users/update", updateUserHandler)
	r.Delete("/api/users/delete", deleteUserHandler)

	// Run the server
	log.Fatalln(http.ListenAndServe(":8080", r))

}

func listUserHandler(w http.ResponseWriter, req *http.Request) { // Handler
	err := json.NewEncoder(w).Encode(UserStore)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// CreateUserRequest ()
type CreateUserRequest struct {
	Email string `json:"email"`
}

// CreateUserResponse ()
type CreateUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// JSON Decode to struct
	req := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := uuid.Must(uuid.NewV4()).String()
	// Business Logic
	newUser := &User{
		ID:    userID,
		Email: req.Email,
	}
	createUser(newUser)

	// JSON Encode from struct
	// Prepare response

	resp := &CreateUserResponse{
		ID:    userID,
		Email: req.Email,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func readUserHandler(w http.ResponseWriter, r *http.Request) {

}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
func helloHandler(w http.ResponseWriter, r *http.Request) {

}