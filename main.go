package main

import (
	"log"
	"net/http"
	"noteapp/controller"
	"noteapp/data/database"
	"noteapp/data/repository"
	"noteapp/service"
)

func main() {
	dbConn, err := database.CreateDBConn()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewNoteRepository(dbConn)

	noteService := service.NewNoteService(repo)

	noteController := controller.NewNoteController(noteService)

	log.Fatal(http.ListenAndServe(":80", noteController.InitRouter()))
}
