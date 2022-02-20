package server

import (
	"context"

	"github.com/joaoaab/libpostal-grpc/src/libpostal"
	protos "github.com/joaoaab/libpostal-grpc/src/protos"
)

type ParserServer struct {
	protos.UnimplementedParserServer
}

func NewParserServer() *ParserServer {
	return &ParserServer{}
}

func (p ParserServer) ParseAddress(ctx context.Context, request *protos.ParseAddressRequest) (*protos.ParsedAddressResponse, error) {
	var parserAddressInput = libpostal.AddressParserInput{
		AddressLine: request.Address,
		Language:    request.Language,
		Country:     request.Country,
	}
	var parsedAddressResponse = libpostal.ParseAddress(parserAddressInput)

	return parsedAddressResponse, nil
}
