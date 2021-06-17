package requests

type ShareNoteRequest struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
}

type PrivateNoteRequest struct {
	NoteID string `json:"note_id"`
}

type GetNoteRequest struct {
	NoteID string `json:"note_id"`
}
