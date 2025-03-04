package handlers

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"

	"github.com/Yhrone/rest-api/db"
	"go.mongodb.org/mongo-driver/bson"
)

func MongoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var filter bson.M
	if err := json.Unmarshal(body, &filter); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	results, err := db.FetchFromMongoDB(filter)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
