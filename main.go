package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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