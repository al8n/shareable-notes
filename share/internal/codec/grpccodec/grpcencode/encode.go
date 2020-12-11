package grpcencode

import (
	"context"
	"github.com/ALiuGuanyan/margin/share/internal/codec/grpccodec"
	"github.com/ALiuGuanyan/margin/share/internal/utils"
	"github.com/ALiuGuanyan/margin/share/model/requests"
	"github.com/ALiuGuanyan/margin/share/model/responses"
	"github.com/ALiuGuanyan/margin/share/pb"
)

func ShareNoteRequest(_ context.Context, request interface{}) ( interface{}, error)  {
	req, ok := request.(requests.ShareNoteRequest)
	if !ok {
		return nil, utils.ErrorCodecCasting("ShareNote", utils.Request,utils.GRPC)
	}

	return grpccodec.ShareNoteReq2pbReq(req), nil
}

func ShareNoteResponse(_ context.Context, resp interface{}) (interface{}, error) {
	pbReply := &pb.ShareNoteResponse{}
	res, ok := resp.(responses.ShareNoteResponse)
	if !ok {
		return nil, utils.ErrorCodecCasting("ShareNote", utils.Response, utils.GRPC)
	}

	if res.Error != "" {
		return nil, utils.Str2Err(res.Error)
	}

	pbReply.Error = res.Error
	pbReply.NoteId = res.NoteID
	pbReply.Url = res.URL

	return pbReply, nil
}

func GetNoteRequest(_ context.Context, request interface{}) ( interface{}, error)  {
	req, ok := request.(requests.GetNoteRequest)
	if !ok {
		return nil, utils.ErrorCodecCasting("GetNote", utils.Request,utils.GRPC)
	}
	return grpccodec.GetNoteReq2pbReq(req), nil
}

func GetNoteResponse(_ context.Context, resp interface{}) (interface{}, error) {
	pbReply := &pb.GetNoteResponse{}
	res, ok := resp.(responses.GetNoteResponse)
	if !ok {
		return nil, utils.ErrorCodecCasting("GetNote", utils.Response, utils.GRPC)
	}

	if res.Error != "" {
		return nil, utils.Str2Err(res.Error)
	}

	pbReply.Content = res.Content
	pbReply.Name = res.Name

	return pbReply, nil
}

func PrivateNoteRequest(_ context.Context, request interface{}) ( interface{}, error)  {
	req, ok := request.(requests.PrivateNoteRequest)
	if !ok {
		return nil, utils.ErrorCodecCasting("PrivateNote", utils.Request,utils.GRPC)
	}
	return grpccodec.PrivateNoteReq2pbReq(req), nil
}

func PrivateNoteResponse(_ context.Context, resp interface{}) (interface{}, error)  {
	pbReply := &pb.PrivateNoteResponse{}
	res, ok := resp.(responses.PrivateNoteResponse)
	if !ok {
		return nil, utils.ErrorCodecCasting("PrivateNote", utils.Response, utils.GRPC)
	}
	if res.Error != "" {
		return nil, utils.Str2Err(res.Error)
	}

	pbReply.Error = res.Error
	return pbReply, nil
}