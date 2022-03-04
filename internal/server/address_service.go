package server

import (
	"context"

	protos "github.com/joaoaab/libpostal-grpc/api/protos"
	"github.com/joaoaab/libpostal-grpc/internal/libpostal"
)

type Server struct {
	protos.UnimplementedAddressServer
}

func NewAddressServer() protos.AddressServer {
	return &Server{}
}

func (p Server) ParseAddress(ctx context.Context, request *protos.ParseAddressRequest) (*protos.ParsedAddressResponse, error) {
	parsedAddressResponse := libpostal.ParseAddress(request)

	return parsedAddressResponse, nil
}

func (p Server) ExpandAddress(ctx context.Context, request *protos.ExpandAddressRequest) (*protos.ExpandedAddressResponse, error) {
	expandedAddress := libpostal.ExpandAddress(request)
	return expandedAddress, nil
}
