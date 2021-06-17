package httpencode

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/al8n/shareable-notes/share-svc/internal/codec/httpcodec"
	"github.com/al8n/shareable-notes/share-svc/internal/utils"
	"github.com/al8n/shareable-notes/share-svc/model/requests"
	"github.com/al8n/shareable-notes/share-svc/model/responses"
	"io/ioutil"
	"net/http"
	"net/url"
)



func GetNoteRequest(ctx context.Context, req *http.Request, request interface{}) error  {
	r := request.(requests.GetNoteRequest)
	noteID := url.QueryEscape(r.NoteID)

	req.URL.Path = "/note/" + noteID
	return GenericRequest(ctx, req, request)
}

// GenericRequest is a transport/http.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func GenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func ShareNoteResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) (err error)  {

	response, ok := resp.(responses.ShareNoteResponse)
	if !ok {

		httpcodec.ErrorEncoder(
			ctx,
			utils.ErrorCodecCasting(
				"ShareNote",
				utils.Response,
				utils.HTTP),
			w)
		return nil
	}

	if response.Error != "" {
		httpcodec.ErrorEncoder(
			ctx,
			utils.Str2Err(response.Error),
			w)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}

func PrivateNoteResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error  {

	response, ok := resp.(responses.PrivateNoteResponse)
	if !ok {
		httpcodec.ErrorEncoder(
			ctx,
			utils.ErrorCodecCasting(
				"PrivateNote",
				utils.Response,
				utils.HTTP),
			w)
		return nil
	}
	if response.Error != "" {
		httpcodec.ErrorEncoder(
			ctx,
			utils.Str2Err(response.Error),
			w)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}

func GetNoteResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error  {

	response, ok := resp.(responses.GetNoteResponse)
	if !ok {
		httpcodec.ErrorEncoder(
			ctx,
			utils.ErrorCodecCasting(
				"GetNote",
				utils.Response,
				utils.HTTP),
			w)
		return nil
	}

	if response.Error != "" {
		httpcodec.ErrorEncoder(
			ctx,
			utils.Str2Err(response.Error),
			w)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}