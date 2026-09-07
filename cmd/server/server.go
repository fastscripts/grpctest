package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fastscripts/grpctest"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("schicke Antwort: " + in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// mtls setup
	/*
		caCert, err := os.ReadFile("ca.crt")
		if err != nil {
			log.Fatal(err)
		}

		caPool := x509.NewCertPool()

		if !caPool.AppendCertsFromPEM(caCert) {
			log.Fatal("CA konnte nicht geladen werden")
		}

		serverCert, err := tls.LoadX509KeyPair(
			"server.crt",
			"server.key",
		)
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig := &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  caPool,
			Certificates: []tls.Certificate{
				serverCert,
			},
		}

		creds := credentials.NewTLS(tlsConfig)

		s := grpc.NewServer(
			grpc.Creds(creds),
		)
	*/
	// TLS setup
	/*
		creds, err := credentials.NewServerTLSFromFile(
			"server.crt",
			"server.key",
		)
		if err != nil {
			log.Fatalf("TLS setup failed: %v", err)
		}
		s := grpc.NewServer(
			grpc.Creds(creds),
		)
	*/
	// Without TLS

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
