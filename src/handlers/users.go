package handlers

import "net/http"

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
