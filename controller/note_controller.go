package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"noteapp/controller/request"
	"noteapp/service"
)

type NoteController struct {
	service service.INoteService
}

func (controller *NoteController) GetAllNotes(rw http.ResponseWriter, r *http.Request) {
	notes, err := controller.service.GetAllNotes()
	if err != nil {
		fmt.Fprintln(rw, err)
	} else {
		var response []request.UpdateNoteRequest
		for _, element := range notes {
			response = append(response,
				request.UpdateNoteRequest{
					Id:       element.Id,
					Title:    element.Title,
					Body:     element.Body,
					ImageUrl: element.ImageUrl,
				},
			)
		}

		json.NewEncoder(rw).Encode(response)
	}
}

func (controller *NoteController) AddNote(rw http.ResponseWriter, r *http.Request) {
	var note request.AddNoteRequest
	json.NewDecoder(r.Body).Decode(&note)
	err := controller.service.AddNote(note)
	if err != nil {
		fmt.Fprintln(rw, err)
	} else {
		fmt.Fprintln(rw, "Note added successfully.")
	}
}

func (controller *NoteController) GetNoteById(rw http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	noteId := args["id"]
	note, err := controller.service.GetNoteById(noteId)
	if err != nil {
		fmt.Fprintln(rw, err)
	} else {
		var noteResponse = request.UpdateNoteRequest{
			Id:       note.Id,
			Title:    note.Title,
			Body:     note.Body,
			ImageUrl: note.ImageUrl,
		}

		json.NewEncoder(rw).Encode(noteResponse)
	}
}

func (controller *NoteController) UpdateNote(rw http.ResponseWriter, r *http.Request) {
	var note request.UpdateNoteRequest
	json.NewDecoder(r.Body).Decode(&note)
	err := controller.service.UpdateNote(note)
	if err != nil {
		fmt.Fprintln(rw, err)
	} else {
		fmt.Fprintln(rw, "Successfully updated.")
	}
}

func (controller *NoteController) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Path("/notes").HandlerFunc(controller.GetAllNotes).Methods(http.MethodGet)
	router.Path("/note/{id}").HandlerFunc(controller.GetNoteById).Methods(http.MethodGet)
	router.Path("/note").HandlerFunc(controller.AddNote).Methods(http.MethodPost)
	router.Path("/note").HandlerFunc(controller.UpdateNote).Methods(http.MethodPut)
	return router
}

func NewNoteController(noteService service.INoteService) *NoteController {
	return &NoteController{service: noteService}
}
