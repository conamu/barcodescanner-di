package main

import (
	"fmt"
	"strings"
)

func getBarcode() string {
	fmt.Println("Please scan or enter a barcode:")
	scanner.Scan()
	return scanner.Text()
}

func validateBarcode(barcode string) (string, bool) {

	var code string
	var valid bool

	if strings.HasPrefix(barcode, config.validityPrefix) || barcode == "end" {
		code = barcode
		valid = true
	} else if !config.api.enable{
		fmt.Println("Please enter a valid barcode.")
		return "", false
	} else {
		return "", false
	}
	return code, valid
}

func getParams() (name, category, description string) {

	fmt.Print("Whats the Product Name? (max. 150 characters): ")
	scanner.Scan()
	name = scanner.Text()
	name = charLimiter(name, 150)

	fmt.Print("Whats the Product Category? (max. 20 characters): ")
	scanner.Scan()
	category = scanner.Text()
	category = charLimiter(category, 20)

	fmt.Print("Whats the Product Description? (max. 500 characters): ")
	scanner.Scan()
	description = scanner.Text()
	description = charLimiter(description, 500)

	return name, category, description
}