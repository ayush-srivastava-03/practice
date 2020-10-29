package main

import (
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var membersList []MemberDetails

var networksMap = make(map[int64][]int64)

func main() {

	membersList = []MemberDetails{

		{
			User: &User{
				Key:  1,
				Name: "Alice",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Bob", "Chris", "Dean"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Study", "Reading", "Dancing"},
			},
		},

		{
			User: &User{
				Key:  2,
				Name: "Bob",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Alice", "Chris", "Dean"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Movies", "Study", "Gardening"},
			},
		},

		{
			User: &User{
				Key:  3,
				Name: "Chris",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Alice", "Bob", "Dean"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Study", "Movies", "Sports"},
			},
		},

		{
			User: &User{
				Key:  4,
				Name: "Dean",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Alice", "Bob", "Chris"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Astrology", "Gardening", "Study"},
			},
		},

		{
			User: &User{
				Key:  5,
				Name: "Evan",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Fred", "Gippy", "Harry"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Gardening", "Movies", "Astrology", "Music"},
			},
		},

		{
			User: &User{
				Key:  6,
				Name: "Fred",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Evan", "Gippy", "Harry"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Astrology", "Reading", "Study"},
			},
		},

		{
			User: &User{
				Key:  7,
				Name: "Gippy",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Evan", "Fred", "Harry"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Music", "Movies"},
			},
		},

		{
			User: &User{
				Key:  8,
				Name: "Harry",
			},
			CommonContacts: &Contacts{
				Contacts: []string{"Evan", "Fred", "Gippy"},
			},
			CommonInterests: &Interests{
				Interests: []string{"Astrology", "Gardening", "Movies", "Music"},
			},
		},
	}

	networksMap[1000] = []int64{1, 2, 3, 4}
	networksMap[2000] = []int64{1, 5, 6, 7}
	networksMap[3000] = []int64{2, 5, 7, 8}
	networksMap[4000] = []int64{3, 8}
	networksMap[5000] = []int64{4, 7}

	/*
		//The service that provide the data runs on a different port, as part of a different application
		grpcConn, err := grpc.Dial(":3031", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		networkClient := NewNetworkServiceClient(grpcConn)
		userClient := NewUserServiceClient(grpcConn)
		contactClient := NewContactServiceClient(grpcConn)
		interestsClient := NewInterestsServiceClient(grpcConn)
	*/
	lis, err := net.Listen("tcp", ":3030")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	/*
		Register your service implimentation here, like so:

		RegisterViewNetworkServiceServer(grpcServer, --your server implimentation here--)
	*/

	RegisterViewNetworkServiceServer(grpcServer, &server{})
	RegisterNetworkServiceServer(grpcServer, &server{})
	RegisterUserServiceServer(grpcServer, &server{})
	RegisterInterestsServiceServer(grpcServer, &server{})
	RegisterContactServiceServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}

}

func (s *server) ViewNetworkMembers(ctx context.Context, req *UserViewingNetwork) (*NetworkMembersView, error) {

	// Get the user key from the user viewing network
	var uKey *UserKey = req.GetUser()
	if uKey == nil {
		return nil, errors.New("User Key is null")
	}

	// Get the network key from the user viewing network
	var nKey *NetworkKey = req.GetNetwork()
	if nKey == nil {
		return nil, errors.New("Network Key is null")
	}

	uKeysInNetwork, err := s.GetNetworkMembers(ctx, nKey)
	if err != nil {
		return nil, err
	}

	var memberDetailsAdded []*MemberDetails

	for _, v := range uKeysInNetwork.Users {

		if v.Key == uKey.Key {
			continue
		}

		userFromUserKey, err := s.GetUser(ctx, v)
		if err != nil {
			return nil, err
		}

		twoKeys := &TwoUserKeys{
			User1: &UserKey{
				Key: uKey.Key,
			},
			User2: &UserKey{
				Key: v.Key,
			},
		}

		commonContacts, err := s.GetCommonContacts(ctx, twoKeys)
		if err != nil {
			return nil, err
		}

		sharedInterests, err := s.GetSharedInterests(ctx, twoKeys)
		if err != nil {
			return nil, err
		}

		memberDetailsAdded = append(memberDetailsAdded, &MemberDetails{
			User:            userFromUserKey,
			CommonContacts:  commonContacts,
			CommonInterests: sharedInterests,
		})

	}

	return &NetworkMembersView{Members: memberDetailsAdded}, nil
}

func (s *server) GetNetworkMembers(ctx context.Context, nKey *NetworkKey) (*UserKeys, error) {
	if _, ok := networksMap[nKey.Key]; !ok {
		return nil, errors.New("No such network key found")
	}

	var keys_list []int64 = networksMap[nKey.Key]

	var userKeysAdded []*UserKey

	for _, v := range keys_list {
		userKeysAdded = append(userKeysAdded, &UserKey{Key: v})
	}

	return &UserKeys{Users: userKeysAdded}, nil
}

func (s *server) GetUser(ctx context.Context, uKey *UserKey) (*User, error) {

	// Search for a user from the member details who has the same user key
	// as the user key obtained from the request
	found := false
	index := -1
	for k, v := range membersList {
		if v.User.Key == uKey.Key {
			found = true
			index = k
			break
		}
	}

	// If no such member details found
	// 		 return the error "No such member found"
	if !found {
		return nil, errors.New("No such member found")
	}

	return membersList[index].User, nil
}

func (s *server) GetSharedInterests(ctx context.Context, twoKeys *TwoUserKeys) (*Interests, error) {

	// Search for the two user keys in the member details
	foundCounter := 0
	index1 := -1
	index2 := -1
	for k, v := range membersList {
		if v.User.Key == twoKeys.User1.Key {
			foundCounter++
			index1 = k
		} else if v.User.Key == twoKeys.User2.Key {
			foundCounter++
			index2 = k
		}

		if foundCounter == 2 {
			break
		}
	}

	// If any of those 2 member details not found, then return error
	if foundCounter != 2 {
		return nil, errors.New("User details not found for one of the users")
	}

	var interestsAdded []string

	for _, v1 := range membersList[index1].CommonInterests.Interests {
		for _, v2 := range membersList[index2].CommonInterests.Interests {
			if v1 == v2 {
				interestsAdded = append(interestsAdded, v1)
				break
			}
		}
	}

	return &Interests{Interests: interestsAdded}, nil
}

func (s *server) GetCommonContacts(ctx context.Context, twoKeys *TwoUserKeys) (*Contacts, error) {

	// Search for the two user keys in the member details
	foundCounter := 0
	index1 := -1
	index2 := -1
	for k, v := range membersList {
		if v.User.Key == twoKeys.User1.Key {
			foundCounter++
			index1 = k
		} else if v.User.Key == twoKeys.User2.Key {
			foundCounter++
			index2 = k
		}

		if foundCounter == 2 {
			break
		}
	}

	// If any of those 2 member details not found, then return error
	if foundCounter != 2 {
		return nil, errors.New("User details not found for one of the users")
	}

	var contactsAdded []string

	for _, v1 := range membersList[index1].CommonContacts.Contacts {
		for _, v2 := range membersList[index2].CommonContacts.Contacts {
			if v1 == v2 {
				contactsAdded = append(contactsAdded, v1)
				break
			}
		}
	}

	return &Contacts{Contacts: contactsAdded}, nil
}
