package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fastscripts/grpctest/gen/go/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUserByID(ctx, &pb.UserIDRequest{UserId: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", r)
}

/*
	// Call the GetUserByID method
	userID := int32(1) // Replace with the desired user ID
	req := &pb.UserIDRequest{UserId: userID}
	res, err := c.GetUserByID(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
*/
