package libpostal

import (
	"testing"

	"github.com/joaoaab/libpostal-grpc/src/protos"
	parser "github.com/openvenues/gopostal/parser"
	"github.com/stretchr/testify/assert"
)

const WORLD_REGION = "test_world_region"
const COUNTRY = "test_country"
const COUNTRY_REGION = "test_country_region"
const STATE = "test_state"
const STATE_DISTRICT = "test_state_district"
const ISLAND = "test_island"
const CITY = "test_city"
const CITY_DISTRICT = "test_city_district"
const SUBURB = "test_suburb"
const POST_CODE = "test_post_code"
const PO_BOX = "test_po_box"
const ENTRANCE = "test_entrance"
const STAIRCASE = "test_staircase"
const LEVEL = "test_level"
const UNIT = "test_unit"
const ROAD = "test_road"
const HOUSE_NUMBER = "test_house_nmb"
const NEAR = "test_near"
const CATEGORY = "test_category"
const HOUSE = "test_house"

func TestParseAddress(t *testing.T) {
	var input = AddressParserInput{AddressLine: "17 Rue du Médecin-Colonel Calbairac 31000 Toulouse France", Country: "FR", Language: "fr"}
	var response = *ParseAddress(input)
	var expected = protos.ParsedAddressResponse{
		Country:     "france",
		City:        "toulouse",
		Postcode:    "31000",
		Road:        "rue du médecin-colonel calbairac",
		HouseNumber: "17",
	}
	assert.Equal(t, response, expected)
}

func TestParseAddressGivenAbsentOptions(t *testing.T) {
	var input = AddressParserInput{AddressLine: "17 Rue du Médecin-Colonel Calbairac 31000 Toulouse France"}
	var response = *ParseAddress(input)
	var expected = protos.ParsedAddressResponse{
		Country:     "france",
		City:        "toulouse",
		Postcode:    "31000",
		Road:        "rue du médecin-colonel calbairac",
		HouseNumber: "17",
	}
	assert.Equal(t, response, expected)
}

func TestToResponse(t *testing.T) {
	var parsedComponents = []parser.ParsedComponent{
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
	var response = *ToResponse(parsedComponents)
	var expected = protos.ParsedAddressResponse{
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
	assert.Equal(t, response, expected)
}

func TestToResponseGivenEmptyMap(t *testing.T) {
	var parsedComponents []parser.ParsedComponent
	var response = *ToResponse(parsedComponents)
	var expected = protos.ParsedAddressResponse{}
	assert.Equal(t, response, expected)
}
