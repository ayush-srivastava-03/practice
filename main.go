package main

import (
	"net"

	"google.golang.org/grpc"
	"interview/data"
	"interview/proto"
	"interview/services/contacts"
	"interview/services/interests"
	"interview/services/network"
	"interview/services/users"
	"interview/services/view"
)

func init() {
	data.MembersList = []proto.MemberDetails{

							proto.MemberDetails {
								User:   			&proto.User{
														Key: 1,
														Name: "Alice",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Bob","Chris","Dean"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Study","Reading","Dancing"},
													},
							},

							proto.MemberDetails {
								User:   			&proto.User{
														Key: 2,
														Name: "Bob",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Alice","Chris","Dean"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Movies","Study", "Gardening"},
													},
							},

							proto.MemberDetails {
								User:   			&proto.User{
														Key: 3,
														Name: "Chris",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Alice","Bob","Dean"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Study","Movies","Sports"},
													},
							},
							
							proto.MemberDetails {
								User:   			&proto.User{
														Key: 4,
														Name: "Dean",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Alice","Bob","Chris"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Astrology","Gardening","Study"},
													},
							},
							
							proto.MemberDetails {
								User:   			&proto.User{
														Key: 5,
														Name: "Evan",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Fred","Gippy","Harry"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Gardening","Movies", "Astrology", "Music"},
													},
							},
							
							proto.MemberDetails {
								User:   			&proto.User{
														Key: 6,
														Name: "Fred",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Evan","Gippy","Harry"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Astrology","Reading","Study"},
													},
							},
							
							proto.MemberDetails {
								User:   			&proto.User{
														Key: 7,
														Name: "Gippy",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Evan","Fred","Harry"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Music","Movies"},
													},
							},
							
							proto.MemberDetails {
								User:   			&proto.User{
														Key: 8,
														Name: "Harry",
													},
								CommonContacts: 	&proto.Contacts{
														Contacts: []string{"Evan","Fred","Gippy"},
													},
								CommonInterests:    &proto.Interests{
														Interests: []string{"Astrology","Gardening","Movies","Music"},
													},
							},
					}

	data.NetworksMap[1000] = []int64{1,2,3,4}
	data.NetworksMap[2000] = []int64{1,5,6,7}
	data.NetworksMap[3000] = []int64{2,5,7,8}
	data.NetworksMap[4000] = []int64{3,8}
	data.NetworksMap[5000] = []int64{4,7}
}

func main() {

	lis, err := net.Listen("tcp", ":3030")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	
	// Registering of all the services
	proto.RegisterViewNetworkServiceServer(grpcServer, &view.ViewNetworkServer{})
	proto.RegisterNetworkServiceServer(grpcServer, &network.NetworkServer{})
	proto.RegisterUserServiceServer(grpcServer, &users.UsersServer{})
	proto.RegisterInterestsServiceServer(grpcServer, &interests.InterestsServer{})
	proto.RegisterContactServiceServer(grpcServer, &contacts.ContactsServer{})
	
	
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}