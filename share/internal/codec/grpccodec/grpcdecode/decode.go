package grpcdecode

import (
	"context"
	"github.com/al8n/shareable-notes/share/internal/codec/grpccodec"
	"github.com/al8n/shareable-notes/share/model/requests"
	"github.com/al8n/shareable-notes/share/pb"
)

func ShareNoteRequest(_ context.Context, grpcReq interface{}) (interface{}, error)  {
	req := grpcReq.(*pb.ShareNoteRequest)
	return requests.ShareNoteRequest{
		Name: req.Name,
		Content: req.Content,
	}, nil
}

func ShareNoteResponse(_ context.Context, grpcReq interface{}) (interface{}, error)  {
	req := grpcReq.(*pb.ShareNoteResponse)
	return grpccodec.ShareNotepbResp2Resp(*req), nil
}


func PrivateNoteRequest(_ context.Context, grpcReq interface{}) (interface{}, error)  {
	req := grpcReq.(*pb.PrivateNoteRequest)
	return requests.PrivateNoteRequest{
		NoteID: req.NoteId,
	}, nil
}

func PrivateNoteResponse(_ context.Context, grpcReq interface{}) (interface{}, error)  {
	req := grpcReq.(*pb.PrivateNoteResponse)
	return grpccodec.PrivateNotepbResp2Resp(*req), nil
}

func GetNoteRequest(_ context.Context, grpcReq interface{}) (interface{}, error)  {
	req := grpcReq.(*pb.GetNoteRequest)

	return requests.GetNoteRequest{
		NoteID: req.Id,
	}, nil
}

func GetNoteResponse(_ context.Context, grpcReq interface{}) (interface{}, error)  {
	req := grpcReq.(*pb.GetNoteResponse)
	return grpccodec.GetNotepbResp2Resp(*req), nil
}