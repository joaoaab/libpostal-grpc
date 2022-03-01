package libpostal

import (
	"github.com/joaoaab/libpostal-grpc/api/protos"
	expand "github.com/openvenues/gopostal/expand"
)

type AddressExpanderInput struct {
	AddressLine string
	Options     expand.ExpandOptions
}

func ExpandAddress(proto *protos.ExpandAddressRequest) *protos.ExpandedAddressResponse {
	expandInput := fromProto(proto)

	expandedAddress := expand.ExpandAddressOptions(expandInput.AddressLine, expandInput.Options)
	return toProto(expandedAddress)
}

func fromProto(proto *protos.ExpandAddressRequest) AddressExpanderInput {
	if proto.Options == nil {
		return AddressExpanderInput{AddressLine: proto.Address, Options: expand.GetDefaultExpansionOptions()}
	}

	return AddressExpanderInput{
		AddressLine: proto.Address,
		Options: expand.ExpandOptions{
			Languages:              proto.Options.Languages,
			AddressComponents:      uint16(proto.Options.AddressComponents),
			LatinAscii:             proto.Options.LatinAscii,
			Transliterate:          proto.Options.Transliterate,
			StripAccents:           proto.Options.StripAccents,
			Decompose:              proto.Options.Decompose,
			Lowercase:              proto.Options.LowerCase,
			TrimString:             proto.Options.TrimString,
			ReplaceWordHyphens:     proto.Options.ReplaceWordHyphens,
			DeleteWordHyphens:      proto.Options.DeleteWordHyphens,
			ReplaceNumericHyphens:  proto.Options.ReplaceNumericHyphens,
			DeleteNumericHyphens:   proto.Options.DeleteNumericHyphens,
			SplitAlphaFromNumeric:  proto.Options.SplitAlphaFromNumeric,
			DeleteFinalPeriods:     proto.Options.DeleteFinalPeriods,
			DeleteAcronymPeriods:   proto.Options.DeleteAcronymPeriods,
			DropEnglishPossessives: proto.Options.DropEnglishPossessives,
			DeleteApostrophes:      proto.Options.DeleteApostrophes,
			ExpandNumex:            proto.Options.ExpandNumex,
			RomanNumerals:          proto.Options.RomanNumerals,
		}}
}

func toProto(addresses []string) *protos.ExpandedAddressResponse {
	return &protos.ExpandedAddressResponse{ExpandedAddress: addresses}
}
