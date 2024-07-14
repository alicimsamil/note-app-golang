package repository

import (
	"database/sql"
	"log"
)

type INoteRepository interface {
	GetAllNotes() ([]Note, error)
	InsertNote(note Note) error
	GetNoteById(id int64) (Note, error)
	UpdateNote(note Note) error
}

type NoteRepository struct {
	dbConn *sql.DB
}

func (noteRepository *NoteRepository) GetAllNotes() ([]Note, error) {
	rows, err := noteRepository.dbConn.Query("SELECT * FROM note")
	if err != nil {
		return []Note{}, err
	}

	return extractNotesFromRows(rows), nil
}

func (noteRepository *NoteRepository) InsertNote(note Note) error {
	_, err := noteRepository.dbConn.Exec("INSERT INTO note (title, body, imageUrl) VALUES ($1, $2, $3)", note.Title, note.Body, note.ImageUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (noteRepository *NoteRepository) GetNoteById(id int64) (Note, error) {
	row := noteRepository.dbConn.QueryRow("SELECT * FROM note WHERE id = $1", id)

	note, err := extractNoteFromRow(row)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (noteRepository *NoteRepository) UpdateNote(note Note) error {
	_, err := noteRepository.dbConn.Exec("UPDATE note SET title = $1, body = $2, imageurl = $3 WHERE id = $4", note.Title, note.Body, note.ImageUrl, note.Id)
	if err != nil {
		return err
	}
	return nil
}

func extractNotesFromRows(rows *sql.Rows) []Note {
	var notes []Note
	var id int64
	var title string
	var body string
	var imageUrl string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &body, &imageUrl); err != nil {
			log.Println(err)
		} else {
			notes = append(notes, Note{
				Id:       id,
				Title:    title,
				Body:     body,
				ImageUrl: imageUrl,
			})
		}
	}

	return notes
}

func extractNoteFromRow(row *sql.Row) (Note, error) {
	var note Note
	var id int64
	var title string
	var body string
	var imageUrl string
	if err := row.Scan(&id, &title, &body, &imageUrl); err != nil {
		log.Println(err)
		return Note{}, err
	} else {
		note = Note{
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
