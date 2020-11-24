package main

import (
	"fmt"

	"gitlab.connectwisedev.com/itsm/POC_Address/address"
)

func main() {
	addr := address.New(
		address.WithCountry("AU"), // Must be an ISO 3166-1 country code
		address.WithName("John Citizen"),
		address.WithOrganization("Some Company Pty Ltd"),
		address.WithStreetAddress([]string{
			"525 Collins Street",
		}),
		address.WithLocality("Melbourne"),
		address.WithAdministrativeArea("VIC"), // If the country has a pre-defined list of admin areas (like here), you must use the key and not the name
		address.WithPostCode("3000"),
	)

	_ = validate(addr)
}

func validate(addr address.Address) error {
	err := address.Validate(addr)
	if err != nil {
		return err
	}

	//if err != nil {
	// If there was an error and you want to find out which validations failed,
	// type switch it as a *multierror.Error to access the list of errors
	// 	if merr, ok := err.(*multierror.Error); ok {
	// 		for _, subErr := range merr.Errors {
	// 			if subErr == address.ErrInvalidCountryCode {
	// 				fmt.Println(subErr)
	// 			}
	// 		}
	// 	}
	// }
	fmt.Println(addr)
	return nil
}
