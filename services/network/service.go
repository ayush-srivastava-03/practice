package network

import (
	"context"
	"errors"

	"interview/data"
	"interview/proto"
)

// NetworkServer structure
type NetworkServer struct{}

// GetNetworkMembers function
func (s *NetworkServer) GetNetworkMembers(ctx context.Context, nKey *proto.NetworkKey) (*proto.UserKeys, error) {
	if _, ok := data.NetworksMap[nKey.Key]; !ok {
		return nil, errors.New("No such network key found")
	}

	var keysList []int64 = data.NetworksMap[nKey.Key]

	var userKeysAdded []*proto.UserKey

	for _, v := range keysList {
		userKeysAdded = append(userKeysAdded, &proto.UserKey{Key: v})
	}

	return &proto.UserKeys{Users: userKeysAdded}, nil
}
