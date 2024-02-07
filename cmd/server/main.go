package main

import (
	"log"
	"net/http"
	"os"
	"rinha-backend-2024/db"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Connecting to database")
	dsn := os.Getenv("DATABASE_URL")
	dbConn, err := db.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer dbConn.DB.Close()

	// create the table
	_, err = dbConn.DB.Exec("CREATE TABLE IF NOT EXISTS clientes (id PRIMARY KEY, limite INT, saldo INT)")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on port 8000")
	router := mux.NewRouter()
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
