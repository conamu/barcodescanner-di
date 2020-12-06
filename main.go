package main

import (
	"github.com/conamu/cliutilsmodule/menustyling"
	"os"
)

func main() {
	// Check if we have the data/ folder, if not, check if we need it, if yes, create it.
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if config.dbType != "sql" {
			os.Mkdir("data", 0755)
		}
	}

	initMainMenu()
	menu := menustyling.GetStoredMenu("main")
	for true {
		// Create a new CodeSession based on configured db Type
		var session codeSession
		switch config.dbType {
		case "flat":
			session = newCodeSession(newFlatDriver())
		case "kv":
			session = newCodeSession(newKvDriver())
		case "sql":
			session = newCodeSession(newSqlDriver())
		}

		menu.DisplayMenu()


	}


}