package main

import (
	"context"
	"testing"

	"interview/proto"
	"interview/services/interests"
)

func TestCommonInterestService(t *testing.T) {
	s := &interests.InterestsServer{}
    // Request object
    req := &proto.TwoUserKeys{
		User1:	&proto.UserKey{
			Key:	2,
		},
		User2:	&proto.UserKey{
			Key:	3,
		},
	}
	//
	if _, err := s.GetSharedInterests(context.Background(), req); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}