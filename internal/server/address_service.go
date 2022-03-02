package server

import (
	"context"

	protos "github.com/joaoaab/libpostal-grpc/api/protos"
	"github.com/joaoaab/libpostal-grpc/internal/libpostal"
)

type AddressServer struct {
	protos.UnimplementedAddressServer
}

func NewAddressServer() *AddressServer {
	return &AddressServer{}
}

func (p AddressServer) ParseAddress(ctx context.Context, request *protos.ParseAddressRequest) (*protos.ParsedAddressResponse, error) {
	parsedAddressResponse := libpostal.ParseAddress(request)

	return parsedAddressResponse, nil
}

func (p AddressServer) ExpandAddress(ctx context.Context, request *protos.ExpandAddressRequest) (*protos.ExpandedAddressResponse, error) {
	expandedAddress := libpostal.ExpandAddress(request)
	return expandedAddress, nil
}
