package repository

import (
	"database/sql"
	"log"
	"noteapp/service/model"
)

type INoteRepository interface {
	GetAllNotes() ([]model.Note, error)
	InsertNote(note model.Note) error
	GetNoteById(id int64) (model.Note, error)
	UpdateNote(note model.Note) error
}

type NoteRepository struct {
	dbConn *sql.DB
}

func (noteRepository *NoteRepository) GetAllNotes() ([]model.Note, error) {
	rows, err := noteRepository.dbConn.Query("SELECT * FROM note")
	if err != nil {
		return []model.Note{}, err
	}

	return extractNotesFromRows(rows), nil
}

func (noteRepository *NoteRepository) InsertNote(note model.Note) error {
	_, err := noteRepository.dbConn.Exec("INSERT INTO note (title, body, imageUrl) VALUES ($1, $2, $3)", note.Title, note.Body, note.ImageUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (noteRepository *NoteRepository) GetNoteById(id int64) (model.Note, error) {
	row := noteRepository.dbConn.QueryRow("SELECT * FROM note WHERE id = $1", id)

	note, err := extractNoteFromRow(row)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (noteRepository *NoteRepository) UpdateNote(note model.Note) error {
	_, err := noteRepository.dbConn.Exec("UPDATE note SET title = $1, body = $2, imageurl = $3 WHERE id = $4", note.Title, note.Body, note.ImageUrl, note.Id)
	if err != nil {
		return err
	}
	return nil
}

func extractNotesFromRows(rows *sql.Rows) []model.Note {
	var notes []model.Note
	var id int64
	var title string
	var body string
	var imageUrl string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &body, &imageUrl); err != nil {
			log.Println(err)
		} else {
			notes = append(notes, model.Note{
				Id:       id,
				Title:    title,
				Body:     body,
				ImageUrl: imageUrl,
			})
		}
	}

	return notes
}

func extractNoteFromRow(row *sql.Row) (model.Note, error) {
	var note model.Note
	var id int64
	var title string
	var body string
	var imageUrl string
	if err := row.Scan(&id, &title, &body, &imageUrl); err != nil {
		log.Println(err)
		return model.Note{}, err
	} else {
		note = model.Note{
			Id:       id,
			Title:    title,
			Body:     body,
			ImageUrl: imageUrl,
		}
		return note, nil
	}
}

func NewNoteRepository(dbConn *sql.DB) INoteRepository {
	return &NoteRepository{dbConn: dbConn}
}
