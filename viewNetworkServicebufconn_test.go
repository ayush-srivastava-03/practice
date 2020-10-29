package main


import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"interview/proto"
)
//
//
func TestViewNetworkUsingBufcon(t *testing.T) {
    //
    ctx := context.Background()
    //
    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(BufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
	//
	client := proto.NewViewNetworkServiceClient(conn)
	//
	req := &proto.UserViewingNetwork{
		User: 		&proto.UserKey{
						Key:	1,
					},	
		Network: 	&proto.NetworkKey{
						Key:	1000,
					},
	}
	//
	//
	if _, err := client.ViewNetworkMembers(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
