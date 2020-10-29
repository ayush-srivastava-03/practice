package main

import (
	"context"
	"testing"

	"interview/proto"
	"interview/services/network"
)

func TestNetworkService(t *testing.T) {
	s := &network.Server{}
	// Request object
	req := &proto.NetworkKey{
		Key: 1000,
	}
	//
	if _, err := s.GetNetworkMembers(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
