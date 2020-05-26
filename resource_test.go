package resource_test

import (
	"log"
	"testing"
	"time"

	resource "github.com/digitalocean/github-resource"
	"github.com/stretchr/testify/assert"
)

func TestCheckRequestRead(t *testing.T) {
	tests := []struct {
		description string
		input       []byte
		expected    resource.CheckRequest
	}{
		{
			description: "invalid source args",
			input:       []byte(`{"source":{"config1":"","config2":"","config3":""},"version":{"oid":"7b262fa","pushed":"2019-11-01T19:57:08Z"}}`),
			expected: resource.CheckRequest{
				Source:  resource.Source{},
				Version: resource.Version{OID: "7b262fa", PushedDate: time.Date(2019, time.November, 1, 19, 57, 8, 0, time.UTC)},
			},
		},
		{
			description: "null version",
			input:       []byte(`{"source":{"repository":"digitalocean/concourse"},"version":null}`),
			expected: resource.CheckRequest{
				Source:  resource.Source{Repository: "digitalocean/concourse"},
				Version: resource.Version{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			var request resource.CheckRequest
			err := request.Read(tc.input)
			if err != nil {
				log.Fatalf("failed to read request input: %s", err)
			}

			assert.Equal(t, tc.expected.Source, request.Source)

			assert.Equal(t, tc.expected.Version.OID, request.Version.OID)
			//assert.Equal(t, tc.expected.Version.PushedDate, request.Version.PushedDate)
		})
	}
}

func TestSourceValidate(t *testing.T) {
	tests := []struct {
		description string
		input       resource.Source
		expected    error
	}{
		{
			description: "no source config",
			input:       resource.Source{},
			expected:    nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			err := tc.input.Validate()

			assert.Equal(t, tc.expected, err)
		})
	}
}
