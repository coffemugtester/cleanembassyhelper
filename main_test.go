package main

import (
	"flag"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	test := []struct {
		args                []string
		expectedHomeCountry string
		expectedHostCountry string
		expectedError       string
	}{
		{
			args:                []string{"main", "-home", "USA", "-host", "Canada"},
			expectedHomeCountry: "USA",
			expectedHostCountry: "Canada",
		},
		{
			args:          []string{"main", "-home", "USA", "-host", "USA"},
			expectedError: "home and host country cannot be the same",
		},
	}

	for _, tt := range test {
		flag.CommandLine = flag.NewFlagSet(tt.args[0], flag.ExitOnError)

		os.Args = tt.args

		homeCountry, hostCountry, err := parseFlags()

		if tt.expectedError != "" {
			assert.Equal(t, err.Error(), tt.expectedError)
			return
		}

		assert.Equal(t, *homeCountry, tt.expectedHomeCountry)
		assert.Equal(t, *hostCountry, tt.expectedHostCountry)

	}
}
