package handler

import (
	"encoding/json"
	"fmt"
	"gorilla-tutorial/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	user := repository.Users[id]
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(json)
}
