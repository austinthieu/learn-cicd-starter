package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		wantErr bool
		wantKey string
	}{
		{
			name: "Has API Key",
			header: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			wantErr: false,
			wantKey: "valid_api_key",
		},
		{
			name:    "Missing Authorization Header",
			header:  http.Header{},
			wantErr: true,
			wantKey: "",
		},
		{
			name: "Malformed Authorization Header",
			header: http.Header{
				"Authorization": []string{"InvalidAPIKey token"},
			},
			wantErr: true,
			wantKey: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if apiKey != tt.wantKey {
				t.Errorf("GetAPIKey(%q) = %v, want %v", tt.header, apiKey, tt.wantKey)
			}
		})
	}
}
