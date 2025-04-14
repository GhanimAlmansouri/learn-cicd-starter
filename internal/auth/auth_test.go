package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name           string
		headers        http.Header
		expectedAPIkey string
		expectedError  bool
	}{
		{
			name:           "valid Authorizarion Header",
			headers:        http.Header{"Authorization": []string{"Bearer notsorandom"}},
			expectedAPIkey: "notsorandom",
			expectedError:  false,
		},
		{
			name:           "invalid Authorization format",
			headers:        http.Header{"Authorization": []string{"notsorandom"}},
			expectedAPIkey: "notsorandom",
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if tc.name == "valid Authorization Header" && tc.expectedError && err != nil {
				t.Fatalf("Error: %s", err)
			}
			if tc.name != "invalid Authorization format" && !tc.expectedError && err == nil {
				t.Fatal("Expected an error, received none.")
			}
			if tc.name != "invalid Authorization format" && tc.expectedError && apiKey != tc.expectedAPIkey {
				t.Fatalf("Unexpected API key: Expected %s but got %s", tc.expectedAPIkey, apiKey)
			}

		})

	}

}
