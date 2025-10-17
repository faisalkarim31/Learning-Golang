package main

import (
	"log"
	"net/http"
)

type Mahasiswa struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	NIM   string `json:"nim"`
	Email string `json:"email"`
}

func main() {
	err := InitDatabase()
	if err != nil {
		log.Fatal("Error database:", err)
	}
	defer CloseDatabase()

	SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
