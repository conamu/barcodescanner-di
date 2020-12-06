package main

import (
	"log"
	"os"
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