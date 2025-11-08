package main

import (
	"final_project/internal/db"
	"final_project/internal/handlers"
	"log"
	"net/http"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer database.Close()

	handlers.SetDB(database)

	if err := db.Migrate(database); err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}

	http.HandleFunc("/candidates", handlers.GetCandidates)          //GET
	http.HandleFunc("/candidates/create", handlers.CreateCandidate) // POST
	http.HandleFunc("/candidates/update", handlers.UpdateCandidate) // PUT
	http.HandleFunc("/candidates/delete", handlers.DeleteCandidate) // DELETE

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
