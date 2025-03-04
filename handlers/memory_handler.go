package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Yhrone/rest-api/db"
)

func MemoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var user db.User
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		db.SaveToMemoryDB(user)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "User added successfully"}`))

	case http.MethodGet:
		users := db.FetchFromMemoryDB()
		response, _ := json.Marshal(users)

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
