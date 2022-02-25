package main

import (
	"fmt"
	"log"
	"net"

	protos "github.com/joaoaab/libpostal-grpc/src/protos"
	"github.com/joaoaab/libpostal-grpc/src/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	address_parser := server.NewAddressServer()

	protos.RegisterAddressServer(s, address_parser)
	reflection.Register(s)

	t1, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(fmt.Println("error starting tcp listener on port 50051", err))
	}

	s.Serve(t1)
}
