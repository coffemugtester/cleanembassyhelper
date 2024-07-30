package scraper

import (
	"clean_embassy_helper/internal/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name           string
		hostCountry    string
		homeCountry    string
		expectedResult []models.Embassy
		expectedError  error
		mockHTML       string
	}{
		{
			name:        "successful data scraping",
			homeCountry: "spain",
			hostCountry: "germany",
			mockHTML:    spainEmbassiesInGermanyHTML,
			expectedResult: []models.Embassy{
				{
					HomeCountry: "spain",
					HostCountry: "germany",
					Name:        "Spain Embassy in Berlin",
					MapLink:     "https://www.google.com/maps?q=Spain+Embassy+in+Berlin+Lichtensteinallee%2C+1+10787+Berlin+Germany",
					City:        "Berlin",
				},
				{
					HomeCountry: "spain",
					HostCountry: "germany",
					Name:        "Spain Consulate in Düsseldorf",
					MapLink:     "https://www.google.com/maps?q=Spain+Consulate+in+Düsseldorf+Hombergerstr.%2C+16+40474+D%C3%BCsseldorf+Germany",
					City:        "Düsseldorf",
				},
				{
					HomeCountry: "spain",
					HostCountry: "germany",
					Name:        "Spain Consulate in Stuttgart",
					MapLink:     "https://www.google.com/maps?q=Spain+Consulate+in+Stuttgart+Lenzhalde%2C+61+70192+Stuttgart+Germany",
					City:        "Stuttgart",
				},
				{
					HomeCountry: "spain",
					HostCountry: "germany",
					Name:        "Spain Consulate in Hamburg",
					MapLink:     "https://www.google.com/maps?q=Spain+Consulate+in+Hamburg+Mittelweg%2C+37+20148+Hamburg+Germany",
					City:        "Hamburg",
				},
				{
					HomeCountry: "spain",
					HostCountry: "germany",
					Name:        "Spain Consulate in Frankfurt",
					MapLink:     "https://www.google.com/maps?q=Spain+Consulate+in+Frankfurt+Nibelungenplatz%2C+3+60318+Frankfurt+Germany",
					City:        "Frankfurt",
				},
				{
					HomeCountry: "spain",
					HostCountry: "germany",
					Name:        "Spain Consulate in München",
					MapLink:     "https://www.google.com/maps?q=Spain+Consulate+in+München+Oberf%C3%B6hringer+Str.%2C+45+81925+M%C3%BCnchen+Germany",
					City:        "München",
				},
			},
		},
		{
			name:          "country misspelled or not in the list",
			homeCountry:   "spainx",
			hostCountry:   "germany",
			mockHTML:      noEmbassiesFound,
			expectedError: fmt.Errorf("no embassies found for spainx in germany"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.mockHTML))
			})
			server := httptest.NewServer(handler)
			defer server.Close()

			client := NewColyClient(server.URL)

			result, err := GetEmbassies(client, tt.homeCountry, tt.hostCountry)
			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError, err)
				return
			}
			assert.Equal(t, tt.expectedResult, result)

		})
	}
}
