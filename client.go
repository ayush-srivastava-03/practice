package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewViewNetworkServiceClient(conn)

	req := &UserViewingNetwork{
		User: &UserKey{
			Key: 1,
		},
		Network: &NetworkKey{
			Key: 1000,
		},
	}

	if response, err := client.ViewNetworkMembers(context.Background(), req); err == nil {
		fmt.Printf("Members in the network of the user key %v are:\n", req.User.Key)
		for _, v := range response.Members {
			//fmt.Printf("User Name: %s\n", v.User.Name)
			fmt.Println(v)
		}
	} else {
		log.Printf("Error: %v", err.Error())
	}

	fmt.Println()

	networkClient := NewNetworkServiceClient(conn)

	req2 := &NetworkKey{
		Key: 1000,
	}

	if response, err := networkClient.GetNetworkMembers(context.Background(), req2); err == nil {
		fmt.Printf("Members in the network key %v are: ", req2.Key)
		for _, v := range response.Users {
			fmt.Printf("%v ", v.Key)
		}
		fmt.Println()
	} else {
		log.Printf("Error: %v", err.Error())
	}

	fmt.Println()

	userClient := NewUserServiceClient(conn)

	req3 := &UserKey{
		Key: 2,
	}

	if response, err := userClient.GetUser(context.Background(), req3); err == nil {
		fmt.Printf("User found with key %v---> Key: %v Name: %v\n", response.Key, response.Key, response.Name)
	} else {
		log.Printf("Error: %v", err.Error())
	}

	fmt.Println()

	interestsClient := NewInterestsServiceClient(conn)

	req4 := &TwoUserKeys{
		User1: &UserKey{
			Key: 2,
		},
		User2: &UserKey{
			Key: 3,
		},
	}

	if response, err := interestsClient.GetSharedInterests(context.Background(), req4); err == nil {
		fmt.Printf("Common Interests between the users with keys %v and %v are: %v\n", req4.User1.Key, req4.User2.Key, response.Interests)
	} else {
		log.Printf("Error: %v", err.Error())
	}

	fmt.Println()

	contactClient := NewContactServiceClient(conn)

	if response, err := contactClient.GetCommonContacts(context.Background(), req4); err == nil {
		fmt.Printf("Common Contacts between the users with keys %v and %v are: %v\n", req4.User1.Key, req4.User2.Key, response.Contacts)
	} else {
		log.Printf("Error: %v", err.Error())
	}

}
