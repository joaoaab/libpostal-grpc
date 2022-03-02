package main

import (
	"fmt"
	"log"
	"net"

	protos "github.com/joaoaab/libpostal-grpc/api/protos"
	"github.com/joaoaab/libpostal-grpc/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = ":50051"

func main() {
	s := grpc.NewServer()

	address_parser := server.NewAddressServer()

	protos.RegisterAddressServer(s, address_parser)
	reflection.Register(s)

	fmt.Println("Starting Server")

	tcpConnection, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(fmt.Println("error starting tcp listener on port 50051", err))
	}

	s.Serve(tcpConnection)
}
