package libpostal

import (
	"github.com/joaoaab/libpostal-grpc/api/protos"
	parser "github.com/openvenues/gopostal/parser"
	"github.com/stoewer/go-strcase"
)

type AddressParserInput struct {
	AddressLine string
	Country     string
	Language    string
}

func ParseAddress(request *protos.ParseAddressRequest) *protos.ParsedAddressResponse {
	parserOptions := parser.ParserOptions{Language: request.Language, Country: request.Country}

	parsedAddress := parser.ParseAddressOptions(request.Address, parserOptions)

	return ToResponse(parsedAddress)
}

func ToResponse(result []parser.ParsedComponent) *protos.ParsedAddressResponse {
	var labelTokenMap map[string]string = make(map[string]string)
	for _, v := range result {
		labelTokenMap[strcase.UpperCamelCase(v.Label)] = v.Value
	}

	return &protos.ParsedAddressResponse{
		WorldRegion:   labelTokenMap["WorldRegion"],
		Country:       labelTokenMap["Country"],
		CountryRegion: labelTokenMap["CountryRegion"],
		State:         labelTokenMap["State"],
		StateDistrict: labelTokenMap["StateDistrict"],
		Island:        labelTokenMap["Island"],
		City:          labelTokenMap["City"],
		CityDistrict:  labelTokenMap["CityDistrict"],
		Suburb:        labelTokenMap["Suburb"],
		Postcode:      labelTokenMap["Postcode"],
		PoBox:         labelTokenMap["PoBox"],
		Entrance:      labelTokenMap["Entrance"],
		Staircase:     labelTokenMap["Staircase"],
		Level:         labelTokenMap["Level"],
		Unit:          labelTokenMap["Unit"],
		Road:          labelTokenMap["Road"],
		HouseNumber:   labelTokenMap["HouseNumber"],
		Near:          labelTokenMap["Near"],
		Category:      labelTokenMap["Category"],
		House:         labelTokenMap["House"],
	}
}
