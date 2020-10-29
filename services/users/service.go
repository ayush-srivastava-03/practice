package users

import (
	"context"
	"errors"

	"interview/data"
	"interview/proto"
)

// Server structure
type Server struct{}

// GetUser function
func (s *Server) GetUser(ctx context.Context, uKey *proto.UserKey) (*proto.User, error) {

	// Search for a user from the member details who has the same user key
	// as the user key obtained from the request
	found := false
	index := -1
	for k, v := range data.MembersList {
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

	return data.MembersList[index].User, nil
}
