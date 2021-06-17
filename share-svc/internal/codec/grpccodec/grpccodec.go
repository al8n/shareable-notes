package grpccodec

import (
	"github.com/al8n/shareable-notes/share-svc/model/requests"
	"github.com/al8n/shareable-notes/share-svc/model/responses"
	"github.com/al8n/shareable-notes/share-svc/pb"
)



func PrivateNoteReq2pbReq(req requests.PrivateNoteRequest) (pbReq *pb.PrivateNoteRequest)  {
	pbReq = &pb.PrivateNoteRequest{}
	pbReq.NoteId = req.NoteID
	return
}

func PrivateNoteResp2pbResp(resp responses.PrivateNoteResponse) (pbResp *pb.PrivateNoteResponse) {
	pbResp = &pb.PrivateNoteResponse{}
	pbResp.Error  = resp.Error
	return
}

func PrivateNotepbResp2Resp(pbResp pb.PrivateNoteResponse) (resp *responses.PrivateNoteResponse) {
	resp = &responses.PrivateNoteResponse{Error: pbResp.Error}
	return
}

func ShareNoteReq2pbReq(req requests.ShareNoteRequest) (pbReq *pb.ShareNoteRequest) {
	pbReq = &pb.ShareNoteRequest{}
	pbReq.Name = req.Name
	pbReq.Content = req.Content
	return
}

func ShareNoteResp2pbResp(resp responses.ShareNoteResponse) (pbResp *pb.ShareNoteResponse) {
	pbResp = &pb.ShareNoteResponse{}
	pbResp.NoteId = resp.NoteID
	pbResp.Url = resp.URL
	pbResp.Error = resp.Error
	return
}

func ShareNotepbResp2Resp(pbResp pb.ShareNoteResponse) (resp *responses.ShareNoteResponse) {
	resp = &responses.ShareNoteResponse{}
	resp.NoteID = pbResp.NoteId
	resp.URL = pbResp.Url
	resp.Error = pbResp.Error
	return
}


func GetNoteReq2pbReq(req requests.GetNoteRequest) (pbReq *pb.GetNoteRequest)  {
	pbReq = &pb.GetNoteRequest{
		Id: req.NoteID,
	}
	return
}

func GetNoteResp2pbResp(resp responses.GetNoteResponse) (pbResp *pb.GetNoteResponse)  {
	pbResp = &pb.GetNoteResponse{
		Name:                 resp.Name,
		Content:              resp.Content,
	}
	return
}

func GetNotepbResp2Resp(pbResp pb.GetNoteResponse) (resp *responses.GetNoteResponse)  {
	resp = &responses.GetNoteResponse{
		Name:    pbResp.Name,
		Content: pbResp.Content,
		Error:   pbResp.Error,
	}
	return resp
}

