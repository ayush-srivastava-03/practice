package main


import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"interview/proto"
)
//
func TestUserServiceUsingBufcon(t *testing.T) {
    //
    ctx := context.Background()
    //
    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(BufDialer), grpc.WithInsecure())
    if err != nil {
        t.Errorf("Failed to dail %v",err)
    }
    defer conn.Close()
	//
	client := proto.NewUserServiceClient(conn)
	//
	req := &proto.UserKey{
		Key:	2,
	}
	//
	//
	if _, err := client.GetUser(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}