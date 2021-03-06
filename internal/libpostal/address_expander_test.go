package libpostal

import (
	"testing"

	expand "github.com/openvenues/gopostal/expand"

	"github.com/joaoaab/libpostal-grpc/api/protos"
	"github.com/stretchr/testify/assert"
)

var DEFAULT_OPTIONS = &protos.ExpandOptions{
	Languages: []string{},
	AddressComponents: (expand.AddressName |
		expand.AddressHouseNumber |
		expand.AddressStreet |
		expand.AddressPoBox |
		expand.AddressUnit |
		expand.AddressLevel |
		expand.AddressEntrance |
		expand.AddressStaircase |
		expand.AddressPostalCode),
	LatinAscii:             true,
	Transliterate:          true,
	StripAccents:           true,
	Decompose:              true,
	LowerCase:              true,
	TrimString:             true,
	ReplaceWordHyphens:     true,
	DeleteNumericHyphens:   false,
	SplitAlphaFromNumeric:  true,
	ReplaceNumericHyphens:  false,
	DeleteWordHyphens:      true,
	DeleteFinalPeriods:     true,
	DeleteAcronymPeriods:   true,
	DropEnglishPossessives: true,
	DeleteApostrophes:      true,
	ExpandNumex:            true,
	RomanNumerals:          true,
}

func TestExpandAddressGivenAbsentOptions(t *testing.T) {
	input := &protos.ExpandAddressRequest{Address: "Quatre-vingt-douze Ave des Ave des Champs-Élysées"}
	response := *ExpandAddress(input)
	expected := protos.ExpandedAddressResponse{
		ExpandedAddress: []string{
			"92 avenue des avenue des champs-elysees",
			"92 avenue des avenue des champs elysees",
			"92 avenue des avenue des champselysees"},
	}

	assert.Equal(t, expected, response)
}

func TestExpandAddressGivenPresentOptions(t *testing.T) {
	input := &protos.ExpandAddressRequest{
		Address: "Quatre-vingt-douze Ave des Ave des Champs-Élysées",
		Options: DEFAULT_OPTIONS}
	response := *ExpandAddress(input)
	expected := protos.ExpandedAddressResponse{
		ExpandedAddress: []string{
			"92 avenue des avenue des champs-elysees",
			"92 avenue des avenue des champs elysees",
			"92 avenue des avenue des champselysees"},
	}

	assert.Equal(t, expected, response)
}

func TestFromProtoGivenAbsentOptions(t *testing.T) {
	input := &protos.ExpandAddressRequest{Address: "Quatre-vingt-douze Ave des Ave des Champs-Élysées"}
	response := fromProto(input)
	expected := AddressExpanderInput{AddressLine: input.Address, Options: expand.GetDefaultExpansionOptions()}

	assert.Equal(t, expected, response)
}

func TestFromProtoGivenPresentOptions(t *testing.T) {
	input := &protos.ExpandAddressRequest{Address: "Quatre-vingt-douze Ave des Ave des Champs-Élysées", Options: DEFAULT_OPTIONS}
	response := fromProto(input)
	expectedOptions := expand.ExpandOptions{
		Languages:              []string{},
		AddressComponents:      uint16(DEFAULT_OPTIONS.AddressComponents),
		LatinAscii:             DEFAULT_OPTIONS.LatinAscii,
		Transliterate:          DEFAULT_OPTIONS.Transliterate,
		StripAccents:           DEFAULT_OPTIONS.StripAccents,
		Decompose:              DEFAULT_OPTIONS.Decompose,
		Lowercase:              DEFAULT_OPTIONS.LowerCase,
		TrimString:             DEFAULT_OPTIONS.TrimString,
		ReplaceWordHyphens:     DEFAULT_OPTIONS.ReplaceWordHyphens,
		DeleteNumericHyphens:   DEFAULT_OPTIONS.DeleteNumericHyphens,
		SplitAlphaFromNumeric:  DEFAULT_OPTIONS.SplitAlphaFromNumeric,
		ReplaceNumericHyphens:  DEFAULT_OPTIONS.ReplaceNumericHyphens,
		DeleteWordHyphens:      DEFAULT_OPTIONS.DeleteWordHyphens,
		DeleteFinalPeriods:     DEFAULT_OPTIONS.DeleteFinalPeriods,
		DeleteAcronymPeriods:   DEFAULT_OPTIONS.DeleteAcronymPeriods,
		DropEnglishPossessives: DEFAULT_OPTIONS.DropEnglishPossessives,
		DeleteApostrophes:      DEFAULT_OPTIONS.DeleteApostrophes,
		ExpandNumex:            DEFAULT_OPTIONS.ExpandNumex,
		RomanNumerals:          DEFAULT_OPTIONS.RomanNumerals,
	}
	expected := AddressExpanderInput{AddressLine: input.Address, Options: expectedOptions}

	assert.Equal(t, expected, response)
}

func testToProto(t *testing.T) {
	addresses := []string{"first string", "second string", "third string"}
	response := toProto(addresses)
	expected := &protos.ExpandedAddressResponse{ExpandedAddress: addresses}

	assert.Equal(t, expected, response)
}
