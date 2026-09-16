package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fastscripts/grpctest/gen/go/proto/user/v1"
	"github.com/fastscripts/grpctest/internal/config"
	"github.com/fastscripts/grpctest/internal/database"
	"github.com/fastscripts/grpctest/internal/logger"
	"github.com/fastscripts/grpctest/internal/service"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServer
	Database *database.DatabaseService
}

func (s *server) GetUserByID(ctx context.Context, in *pb.UserIDRequest) (*pb.UserResponse, error) {

	service := service.NewUserService(*s.Database)

	user, err := service.FindByID(ctx, uint(in.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil

}

func main() {

	logger := logger.NewZerologLogger("info", nil)
	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("failed to load config")
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.Fatal("failed to listen")
	}
	s := grpc.NewServer()

	databaseService, err := database.NewDatabaseService(config)
	if err != nil {
		logger.Fatal("failed to connect to database")
	}

	pb.RegisterUserServer(s, &server{Database: &databaseService})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatal("failed to serve")
	}

}
