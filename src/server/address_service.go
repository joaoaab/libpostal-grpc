package server

import (
	"context"

	"github.com/joaoaab/libpostal-grpc/src/libpostal"
	protos "github.com/joaoaab/libpostal-grpc/src/protos"
)

type AddressServer struct {
	protos.UnimplementedAddressServer
}

func NewAddressServer() *AddressServer {
	return &AddressServer{}
}

func (p AddressServer) ParseAddress(ctx context.Context, request *protos.ParseAddressRequest) (*protos.ParsedAddressResponse, error) {
	var parserAddressInput = libpostal.AddressParserInput{
		AddressLine: request.Address,
		Language:    request.Language,
		Country:     request.Country,
	}
	var parsedAddressResponse = libpostal.ParseAddress(parserAddressInput)

	return parsedAddressResponse, nil
}

func (p AddressServer) ExpandAddress(ctx context.Context, request *protos.ExpandAddressRequest) (*protos.ExpandedAddressResponse, error) {
	var expandedAddress = libpostal.ExpandAddress(request)
	return expandedAddress, nil
}
