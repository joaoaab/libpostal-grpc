package libpostal

import (
	"testing"

	"github.com/joaoaab/libpostal-grpc/api/protos"
	parser "github.com/openvenues/gopostal/parser"
	"github.com/stretchr/testify/assert"
)

const (
	WORLD_REGION   = "test_world_region"
	COUNTRY        = "test_country"
	COUNTRY_REGION = "test_country_region"
	STATE          = "test_state"
	STATE_DISTRICT = "test_state_district"
	ISLAND         = "test_island"
	CITY           = "test_city"
	CITY_DISTRICT  = "test_city_district"
	SUBURB         = "test_suburb"
	POST_CODE      = "test_post_code"
	PO_BOX         = "test_po_box"
	ENTRANCE       = "test_entrance"
	STAIRCASE      = "test_staircase"
	LEVEL          = "test_level"
	UNIT           = "test_unit"
	ROAD           = "test_road"
	HOUSE_NUMBER   = "test_house_nmb"
	NEAR           = "test_near"
	CATEGORY       = "test_category"
	HOUSE          = "test_house"
)

func TestParseAddress(t *testing.T) {
	input := protos.ParseAddressRequest{Address: "17 Rue du Médecin-Colonel Calbairac 31000 Toulouse France", Country: "FR", Language: "fr"}
	response := *ParseAddress(&input)
	expected := protos.ParsedAddressResponse{
		Country:     "france",
		City:        "toulouse",
		Postcode:    "31000",
		Road:        "rue du médecin-colonel calbairac",
		HouseNumber: "17",
	}
	assert.Equal(t, expected, response)
}

func TestParseAddressGivenAbsentOptions(t *testing.T) {
	var input = protos.ParseAddressRequest{Address: "17 Rue du Médecin-Colonel Calbairac 31000 Toulouse France"}
	var response = *ParseAddress(&input)
	var expected = protos.ParsedAddressResponse{
		Country:     "france",
		City:        "toulouse",
		Postcode:    "31000",
		Road:        "rue du médecin-colonel calbairac",
		HouseNumber: "17",
	}
	assert.Equal(t, expected, response)
}

func TestToResponse(t *testing.T) {
	parsedComponents := []parser.ParsedComponent{
		{Label: "world_region", Value: WORLD_REGION},
		{Label: "country", Value: COUNTRY},
		{Label: "country_region", Value: COUNTRY_REGION},
		{Label: "state", Value: STATE},
		{Label: "state_district", Value: STATE_DISTRICT},
		{Label: "island", Value: ISLAND},
		{Label: "city", Value: CITY},
		{Label: "city_district", Value: CITY_DISTRICT},
		{Label: "suburb", Value: SUBURB},
		{Label: "postcode", Value: POST_CODE},
		{Label: "po_box", Value: PO_BOX},
		{Label: "entrance", Value: ENTRANCE},
		{Label: "staircase", Value: STAIRCASE},
		{Label: "level", Value: LEVEL},
		{Label: "unit", Value: UNIT},
		{Label: "road", Value: ROAD},
		{Label: "house_number", Value: HOUSE_NUMBER},
		{Label: "near", Value: NEAR},
		{Label: "category", Value: CATEGORY},
		{Label: "house", Value: HOUSE},
	}
	response := *ToResponse(parsedComponents)
	expected := protos.ParsedAddressResponse{
		WorldRegion:   WORLD_REGION,
		Country:       COUNTRY,
		CountryRegion: COUNTRY_REGION,
		State:         STATE,
		StateDistrict: STATE_DISTRICT,
		Island:        ISLAND,
		City:          CITY,
		CityDistrict:  CITY_DISTRICT,
		Suburb:        SUBURB,
		Postcode:      POST_CODE,
		PoBox:         PO_BOX,
		Entrance:      ENTRANCE,
		Staircase:     STAIRCASE,
		Level:         LEVEL,
		Unit:          UNIT,
		Road:          ROAD,
		HouseNumber:   HOUSE_NUMBER,
		Near:          NEAR,
		Category:      CATEGORY,
		House:         HOUSE,
	}
	assert.Equal(t, expected, response)
}

func TestToResponseGivenEmptyMap(t *testing.T) {
	var parsedComponents []parser.ParsedComponent
	response := *ToResponse(parsedComponents)
	expected := protos.ParsedAddressResponse{}
	assert.Equal(t, expected, response)
}
