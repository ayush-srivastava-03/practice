package main

import (
	"context"
	"testing"

	"interview/proto"
	"interview/services/contacts"
)

func TestContactService(t *testing.T) {
	//
	s := &contacts.Server{}
	// Request object
	req := &proto.TwoUserKeys{
		User1: &proto.UserKey{
			Key: 2,
		},
		User2: &proto.UserKey{
			Key: 3,
		},
	}
	//
	if _, err := s.GetCommonContacts(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
