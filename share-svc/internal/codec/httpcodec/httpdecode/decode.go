package httpdecode

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/al8n/shareable-notes/share-svc/model/requests"
	"github.com/al8n/shareable-notes/share-svc/model/responses"
	"github.com/gorilla/mux"
	"net/http"
)

var ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")

func GetNoteRequest(ctx context.Context, r *http.Request) (interface{}, error) {


	var (
		req requests.GetNoteRequest
	)

	bid, ok := mux.Vars(r)["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	id, err := base64.URLEncoding.DecodeString(bid)
	if err != nil {
		return nil, err
	}

	req.NoteID = string(id)
	return req, nil
}

func PrivateNoteRequest(ctx context.Context, r *http.Request) (interface{}, error)  {

	var req requests.PrivateNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}


	return req, nil
}

func ShareNoteRequest(ctx context.Context, r *http.Request) (interface{}, error)  {
	var (
		req requests.ShareNoteRequest
	)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func PrivateNoteResponse(_ context.Context, r *http.Response) (interface{}, error)  {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp responses.PrivateNoteResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func ShareNoteResponse(_ context.Context, r *http.Response) (interface{}, error)  {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp responses.ShareNoteResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func GetNoteResponse(_ context.Context, r *http.Response) (interface{}, error)  {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp responses.GetNoteResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}