package contacts

import (
	"context"
	"errors"

	"interview/data"
	"interview/proto"
)

// ContactsServer structure
type ContactsServer struct{}

// GetCommonContacts function
func (s *ContactsServer) GetCommonContacts(ctx context.Context, twoKeys *proto.TwoUserKeys) (*proto.Contacts, error) {

	// Search for the two user keys in the member details
	foundCounter := 0
	index1 := -1
	index2 := -1
	for k, v := range data.MembersList {
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

	for _, v1 := range data.MembersList[index1].CommonContacts.Contacts {
		for _, v2 := range data.MembersList[index2].CommonContacts.Contacts {
			if v1 == v2 {
				contactsAdded = append(contactsAdded, v1)
				break
			}
		}
	}

	return &proto.Contacts{Contacts: contactsAdded}, nil
}
