package main

import (
	"testing"

	"gitlab.connectwisedev.com/itsm/POC_Address/address"
)

func Test_validate(t *testing.T) {
	tests := []struct {
		name    string
		addr    address.Address
		wantErr bool
	}{
		{
			name: "Success",
			addr: address.Address{
				Country:      "IN",
				Name:         "Test",
				Organization: "Test",
				StreetAddress: []string{
					"Some Street",
				},
				Locality:           "Test",
				AdministrativeArea: "MH",
				PostCode:           "412212",
			},
			wantErr: false,
		},
		{
			name: "Fail",
			addr: address.Address{
				Country:      "INDIA",
				Name:         "Test",
				Organization: "Test",
				StreetAddress: []string{
					"Some Street",
				},
				Locality:           "Test",
				AdministrativeArea: "MH",
				PostCode:           "412212",
			},
			wantErr: true,
		},
		{
			name: "Fail",
			addr: address.Address{
				Country:      "INDIA",
				Name:         "Test",
				Organization: "Test",
				StreetAddress: []string{
					"Some Street",
				},
				Locality:           "Test",
				AdministrativeArea: "MH",
				PostCode:           "412",
			},
			wantErr: true,
		},
		{
			name: "Fail",
			addr: address.Address{
				Country:      "IN",
				Name:         "Test",
				Organization: "Test",
				StreetAddress: []string{
					"Some Street",
				},
				Locality:           "Test",
				AdministrativeArea: "MP",
				PostCode:           "412212",
			},
			wantErr: true,
		},
		{
			name: "Fail",
			addr: address.Address{
				Country:      "US",
				Name:         "Test",
				Organization: "Test",
				StreetAddress: []string{
					"Some Street",
				},
				Locality:           "Test",
				AdministrativeArea: "MP",
				PostCode:           "412212",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
