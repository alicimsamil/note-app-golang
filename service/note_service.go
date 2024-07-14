package service

import (
	"errors"
	"noteapp/data/repository"
)

type INoteService interface {
	GetAllNotes() ([]Note, error)
	AddNote(note Note) error
	GetNoteById(id int64) (Note, error)
	UpdateNote(note Note) error
}

type NoteService struct {
	repository repository.INoteRepository
}

func (service *NoteService) GetAllNotes() ([]Note, error) {
	return service.repository.GetAllNotes()
}

func (service *NoteService) AddNote(note Note) error {
	if len(note.Title) > 30 {
		return errors.New("title is long")
	} else if len(note.Body) > 255 {
		return errors.New("body is long")
	} else if len(note.ImageUrl) > 100 {
		return errors.New("image url is long")
	}

	return service.repository.InsertNote(note)
}

func (service *NoteService) GetNoteById(id int64) (Note, error) {
	return service.repository.GetNoteById(id)
}

func (service *NoteService) UpdateNote(note Note) error {
	return service.repository.UpdateNote(note)
}

func NewNoteService(repo repository.INoteRepository) INoteService {
	return &NoteService{repository: repo}
}
