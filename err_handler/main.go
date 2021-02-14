package main

import (
	"errors"
	"fmt"
)

// error handler
func validatePhoneNumber(x string) error {
	if len(x) != 10 {
		return errors.New("Phone number invalid.")
	}
	return nil
}


func main() {

	//error handler
	phoneNumber := "082231111"
	if err := validatePhoneNumber(phoneNumber); err != nil {
		fmt.Println(err.Error())
	}

	err := validatePhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println(err.Error())
	}
}
