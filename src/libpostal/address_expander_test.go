package libpostal

import (
	"testing"

	"github.com/joaoaab/libpostal-grpc/src/protos"
	"github.com/stretchr/testify/assert"
)

func TestExpandAddressGivenAbsentOptions(t *testing.T) {
	var input = &protos.ExpandAddressRequest{Address: "Quatre-vingt-douze Ave des Ave des Champs-Élysées"}
	var response = *ExpandAddress(input)
	var expected = protos.ExpandedAddressResponse{
		ExpandedAddress: []string{
			"92 avenue des avenue des champs-elysees",
			"92 avenue des avenue des champs elysees",
			"92 avenue des avenue des champselysees"},
	}

	assert.Equal(t, response, expected)
}
