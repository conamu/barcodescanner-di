package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type dataInterface interface {
	getData() (bool,error , []string)
	writeData() (bool, error)
	editData() (bool, error)
	deleteData() (bool, error)
}

type codeSession struct {
	dataInterface *dataInterface
	code string
	codeValid bool
	data []string
}

func newCodeSession(dataInterface *dataInterface) codeSession  {
	return codeSession {
		dataInterface,
		"",
		false,
		nil,
	}
}

func getBarcode() {
	
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var config = initConfig()

func main() {
	// Check if we have the data/ folder, if not, check if we need it, if yes, create it.
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if config.dbType != "sql" {
			os.Mkdir("data", 0755)
		}
	}

	// Create a new CodeSession based on configured db Type
	var session codeSession
	switch config.dbType {
	case "flat":
		session = newCodeSession(flatConnector)
	case "kv":
		session = newCodeSession(kvConnector)
	case "sql":
		session = newCodeSession(sqlConnector)
	}


}