package main

import (
	"testing"
)

func Test_phone(t *testing.T) {
	tests := []struct {
		name        string
		phoneNo     string
		countryCode string
		wantErr     bool
	}{
		{
			name:        "Success",
			phoneNo:     "919800000000",
			countryCode: "IN",
			wantErr:     false,
		},
		{
			name:        "Failure",
			phoneNo:     "919800000000",
			countryCode: "INDIA",
			wantErr:     true,
		},
		{
			name:        "Failure",
			phoneNo:     "91",
			countryCode: "INDIA",
			wantErr:     true,
		},
		{
			name:        "Success",
			phoneNo:     "+919800000000",
			countryCode: "IN",
			wantErr:     false,
		},
		{
			name:        "Success",
			phoneNo:     "+91-9800000000",
			countryCode: "IN",
			wantErr:     false,
		},
		{
			name:        "Success",
			phoneNo:     "+91-980-000-000-0",
			countryCode: "IN",
			wantErr:     false,
		},
		{
			name:        "Success",
			phoneNo:     "16502530000",
			countryCode: "US",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := phone(tt.phoneNo, tt.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
