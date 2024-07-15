package service

import (
	"errors"
	responseModel "noteapp/controller/model"
	"noteapp/data/repository"
	"noteapp/service/model"
)

type INoteService interface {
	GetAllNotes() ([]model.Note, error)
	AddNote(note responseModel.AddNoteRequest) error
	GetNoteById(id string) (model.Note, error)
	UpdateNote(note responseModel.UpdateNoteRequest) error
}

type NoteService struct {
	repository repository.INoteRepository
}

func (service *NoteService) GetAllNotes() ([]model.Note, error) {
	return service.repository.GetAllNotes()
}

func (service *NoteService) AddNote(note responseModel.AddNoteRequest) error {
	if len(note.Title) > 30 {
		return errors.New("title is long")
	} else if len(note.Body) > 255 {
		return errors.New("body is long")
	} else if len(note.ImageUrl) > 100 {
		return errors.New("image url is long")
	}

	return service.repository.InsertNote(model.Note{Title: note.Title, Body: note.Body, ImageUrl: note.ImageUrl})
}

func (service *NoteService) GetNoteById(id string) (model.Note, error) {
	return service.repository.GetNoteById(id)
}

func (service *NoteService) UpdateNote(note responseModel.UpdateNoteRequest) error {
	return service.repository.UpdateNote(model.Note{Id: note.Id, Title: note.Title, Body: note.Body, ImageUrl: note.ImageUrl})
}

func NewNoteService(repo repository.INoteRepository) INoteService {
	return &NoteService{repository: repo}
}
