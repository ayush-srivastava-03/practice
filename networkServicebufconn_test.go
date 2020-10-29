package main


import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"interview/proto"
)
//
//
func TestNetworkServiceUsingBufcon(t *testing.T) {
    //
    ctx := context.Background()
    //
    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(BufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
	//
	client := proto.NewNetworkServiceClient(conn)
	//
	req := &proto.NetworkKey{
		Key:	1000,
	}
	//
	//
	if _, err := client.GetNetworkMembers(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}