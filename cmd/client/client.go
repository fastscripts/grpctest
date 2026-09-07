package main

import (
	"context"
	"log"
	"time"

	pb "github.com/fastscripts/grpctest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// mtls setup

	/*
		clientCert, err := tls.LoadX509KeyPair(
			"client.crt",
			"client.key",
		)
		caCert, err := os.ReadFile("ca.crt")

		caPool := x509.NewCertPool()
		caPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			ServerName: "grpc-server",
			RootCAs:    caPool,
			Certificates: []tls.Certificate{
				clientCert,
			},
		}
		creds := credentials.NewTLS(tlsConfig)

		conn, err := grpc.NewClient(
			"localhost:50051",
			grpc.WithTransportCredentials(creds),
		)
	*/

	// TLS setup
	/*
		creds, err := credentials.NewClientTLSFromFile(
			"server.crt",
			"localhost",
		)
		if err != nil {
			log.Fatal(err)
		}

		conn , err := grpc.NewClient("localhost:50051",grpc.WithTransportCredentials(creds))
	*/
	// Without TLS
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Micha"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
