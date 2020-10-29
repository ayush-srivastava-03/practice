package main


import (
	"context"
	"testing"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"interview/proto"
	"interview/services/contacts"
	"interview/services/interests"
	"interview/services/network"
	"interview/services/users"
	"interview/services/view"
)
//
const bufSize = 1024 * 1024
//
var lis *bufconn.Listener
//
//
func init() {
    lis = bufconn.Listen(bufSize)
    grpcServer := grpc.NewServer()
	proto.RegisterViewNetworkServiceServer(grpcServer, &view.ViewNetworkServer{})
	proto.RegisterNetworkServiceServer(grpcServer, &network.NetworkServer{})
	proto.RegisterUserServiceServer(grpcServer, &users.UsersServer{})
	proto.RegisterInterestsServiceServer(grpcServer, &interests.InterestsServer{})
	proto.RegisterContactServiceServer(grpcServer, &contacts.ContactsServer{})
    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("Server exited with error: %v", err)
        }
    }()
}
//
//
func BufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

//
func TestContactServiceUsingBufcon(t *testing.T) {
    //
    ctx := context.Background()
    //
    conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(BufDialer), grpc.WithInsecure())
    if err != nil {
        t.Errorf("Failed to dail %v ",err)
    }
    defer conn.Close()
	//
	client := proto.NewContactServiceClient(conn)
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
	if _, err := client.GetCommonContacts(context.Background(), req); err != nil {
		t.Errorf("Failed to find common contacts %v",err)
	}
}