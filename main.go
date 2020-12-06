package main

import (
	"errors"
	"fmt"
	"github.com/conamu/cliutilsmodule/menustyling"
	"os"
	"time"
)

func sleep() {
	time.Sleep(time.Second * 3)
}

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
		switch config.dbType {
		case "flat":
			session = newCodeSession(newFlatDriver())
		case "kv":
			session = newCodeSession(newKvDriver())
		case "sql":
			session = newCodeSession(newSqlDriver())
		}

		// Display the main menu
		menu.DisplayMenu()

		switch menu.GetInputData() {
		case "1":
			_, err, data := session.driver.getData()
			if errors.Is(err, notFound) {
				fmt.Println("This Code was not found.\nPlease Add it first.")
			} else {
				itemDisplay(data[0], data[1], data[2], data[3])
			}
			sleep()
		case "2":
			_, err := session.driver.editData()
			if errors.Is(err, notFound) {
				fmt.Println("This Code was not found.\nPlease Add it first.")
			}
			sleep()
		case "3":
			fmt.Println("!!! - WARNING: SCANNED CODE WILL BE PERMANENTLY REMOVED FROM DATABASE - !!!")
			_, err := session.driver.deleteData()
			if errors.Is(err, notFound) {
				fmt.Println("This Code was not found.\nPlease Add it first.")
			}
			sleep()
		case "4":
			_, err := session.driver.writeData()
			if errors.Is(err, notFound) {
				fmt.Println("This Code already exists in the Database.")
			}
			sleep()
		case "5":
			for true {
				continues, err, data := session.driver.getData()
				if !continues {
					return
				} else {
					if errors.Is(err, notFound) {
						fmt.Println("This Code was not found.\nPlease Add it first.")
					} else {
						itemDisplay(data[0], data[1], data[2], data[3])
					}
				}
			}
		case "6":
			for true {
				continues, err := session.driver.writeData()
				if !continues {
					return
				} else {
					if errors.Is(err, notFound) {
						fmt.Println("This Code already exists in the Database.")
					}
				}
			}
		}


	}


}