package main

import (
	"context"
	"testing"

	"interview/proto"
	"interview/services/users"
)

func TestUserService(t *testing.T) {
	s := &users.UsersServer{}
	// Request object
	req := &proto.UserKey{
		Key: 2,
	}
	//
	if _, err := s.GetUser(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
