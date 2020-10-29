package view

import (
	"context"
	"errors"

	//"interview/data"
	"interview/proto"
	"interview/services/contacts"
	"interview/services/interests"
	"interview/services/network"
	"interview/services/users"
)

// Server structure
type Server struct{}

// ViewNetworkMembers function
func (s *Server) ViewNetworkMembers(ctx context.Context, req *proto.UserViewingNetwork) (*proto.NetworkMembersView, error) {

	networkServer := &network.Server{}
	usersServer := &users.Server{}
	contactsServer := &contacts.Server{}
	interestsServer := &interests.Server{}

	// Get the user key from the user viewing network
	var uKey *proto.UserKey = req.GetUser()
	if uKey == nil {
		return nil, errors.New("User Key is null")
	}

	// Get the network key from the user viewing network
	var nKey *proto.NetworkKey = req.GetNetwork()
	if nKey == nil {
		return nil, errors.New("Network Key is null")
	}

	uKeysInNetwork, err := networkServer.GetNetworkMembers(ctx, nKey)
	if err != nil {
		return nil, err
	}

	var memberDetailsAdded []*proto.MemberDetails

	for _, v := range uKeysInNetwork.Users {

		if v.Key == uKey.Key {
			continue
		}

		userFromUserKey, err := usersServer.GetUser(ctx, v)
		if err != nil {
			return nil, err
		}

		twoKeys := &proto.TwoUserKeys{
			User1: &proto.UserKey{
				Key: uKey.Key,
			},
			User2: &proto.UserKey{
				Key: v.Key,
			},
		}

		commonContacts, err := contactsServer.GetCommonContacts(ctx, twoKeys)
		if err != nil {
			return nil, err
		}

		sharedInterests, err := interestsServer.GetSharedInterests(ctx, twoKeys)
		if err != nil {
			return nil, err
		}

		memberDetailsAdded = append(memberDetailsAdded, &proto.MemberDetails{
			User:            userFromUserKey,
			CommonContacts:  commonContacts,
			CommonInterests: sharedInterests,
		})

	}

	return &proto.NetworkMembersView{Members: memberDetailsAdded}, nil
}
