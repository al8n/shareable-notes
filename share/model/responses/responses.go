package responses

type ShareNoteResponse struct {
	URL string `json:"url"`
	NoteID string `json:"note_id"`
	Error  string `json:"error,omitempty"`
}

type GetNoteResponse struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Error     string `json:"error,omitempty"`
}

type PrivateNoteResponse struct {
	Error string `json:"error,omitempty"`
}