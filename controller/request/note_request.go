package request

type AddNoteRequest struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"imageUrl"`
}

type UpdateNoteRequest struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"imageUrl"`
}
