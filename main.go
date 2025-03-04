package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yhrone/rest-api/db"
	"github.com/Yhrone/rest-api/handlers"
)

func main() {
	// MongoDB Configuration
	mongoURI := "mongodb://localhost:27017"
	dbName := "testdb"
	collectionName := "users"

	db.ConnectMongoDB(mongoURI, dbName, collectionName)

	http.HandleFunc("/mongo", handlers.MongoHandler)
	http.HandleFunc("/memory", handlers.MemoryHandler)

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
