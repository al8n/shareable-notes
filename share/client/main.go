package main

import (
	"context"
	"fmt"
	"github.com/ALiuGuanyan/margin/share/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := pb.NewShareClient(conn)
	resp, err := client.ShareNote(context.Background(), &pb.ShareNoteRequest{Name: "note 7", Content: "asdasjldkjalkdjaslkdjlaksdj"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Url)
	fmt.Println(resp.Error)
	fmt.Println(resp.NoteId)
}
