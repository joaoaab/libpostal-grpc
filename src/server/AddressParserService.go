package server

import (
	"context"

	"github.com/joaoaab/libpostal-grpc/src/libpostal"
	protos "github.com/joaoaab/libpostal-grpc/src/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const EMPTY_STRING = ""

type ParserServer struct {
	protos.UnimplementedParserServer
}

func NewParserServer() *ParserServer {
	return &ParserServer{}
}

func (p ParserServer) ParseAddress(ctx context.Context, request *protos.ParseAddressRequest) (*protos.ParsedAddressResponse, error) {
	var language = request.Language
	var country = request.Country

	if (language != "" && country == "") || (language == "" && country != "") {
		err := status.Error(codes.InvalidArgument, "Must Pass All Options!")
		return nil, err
	}

	var parserAddressInput = libpostal.AddressParserInput{
		AddressLine: request.Address,
		Language:    language,
		Country:     country,
	}
	var parsedAddressResponse = libpostal.ParseAddress(parserAddressInput, requestContainsOptions(request))

	return parsedAddressResponse, nil
}

func requestContainsOptions(request *protos.ParseAddressRequest) bool {
	if request.Language != EMPTY_STRING && request.Country != EMPTY_STRING {
		return true
	}
	return false
}
