package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func getBarcode() *barcode {
	fmt.Println("Please scan or enter a barcode:")
	scanner.Scan()
	return &barcode{
		scanner.Text(),
		false,
	}
}

func validateBarcode(barcode *barcode) *barcode {

	if strings.HasPrefix(barcode.code, config.validityPrefix) || barcode.code == "end" {
		barcode.valid = true
	} else if !config.api.enable{
		fmt.Println("Please enter a valid barcode.")
	} else {
	}
	return barcode
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

func charLimiter(code string, limit int) string {
	//create a new reader, that is gonna read through s string
	reader := strings.NewReader(code)
	//create a buffer (slice of bytes), who's size gonna be limited
	buff := make([]byte, limit)
	// Fill the buff slice initially with spaces so we can TrimSpace the Strings to display them easier.
	for i := len(code); i < len(buff); i++ {
		buff[i] = 32
	}
	//using ReadAtLeast we gonna read (s) into buff until it has read at least minimum byte (limit)
	//it will read also further, but buff is limited by (limit) and it will not take more characters than that
	n, _ := io.ReadAtLeast(reader, buff, limit)
	if n != 0 {
		if len(code) > limit {
			fmt.Printf("You wrote %d characters. Only %d of them will be accepted.\n", len(code), limit)
		}
		return string(buff)
	}
	return code

}