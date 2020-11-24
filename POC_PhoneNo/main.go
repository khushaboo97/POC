package main

import (
	"fmt"

	"gitlab.connectwisedev.com/itsm/POC_PhoneNo/phonenumbers"
)

func phone(no, countryCode string) error {
	num, err := phonenumbers.Parse(no, countryCode)
	if err != nil {
		fmt.Println("Error=", err)
		return err
	}
	fmt.Println("Parsed number = ", num)

	formattedNum := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
	fmt.Println("Formatted number = ", formattedNum)
	return nil
}

func main() {
	_ = phone("6502530000", "US")
}
