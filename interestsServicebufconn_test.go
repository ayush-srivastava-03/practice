package main


import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"interview/proto"
)
//
func TestInterestServiceUsingBufcon(t *testing.T) {
    //
    ctx := context.Background()
    //
    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(BufDialer), grpc.WithInsecure())
    if err != nil {
        t.Errorf("Failed to dail %v ",err)
    }
    defer conn.Close()
	//
	client := proto.NewInterestsServiceClient(conn)
	//
	//
	req := &proto.TwoUserKeys{
		User1:	&proto.UserKey{
			Key:	2,
		},
		User2:	&proto.UserKey{
			Key:	3,
		},
	}
	//
	//
	if _, err := client.GetSharedInterests(context.Background(), req); err != nil {
		t.Errorf("Failed to find common contacts %v",err)
	}
}