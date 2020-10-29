package main

import (
	"context"
	"testing"

	"interview/proto"
	"interview/services/view"
)

func TestViewNetworkService(t *testing.T) {
	s := &view.ViewNetworkServer{}
	// Request object
	req := &proto.UserViewingNetwork{
		User: &proto.UserKey{
			Key: 1,
		},
		Network: &proto.NetworkKey{
			Key: 1000,
		},
	}
	//
	if _, err := s.ViewNetworkMembers(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
