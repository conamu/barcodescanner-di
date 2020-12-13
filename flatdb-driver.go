package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type flatDriver struct {
	barcode *barcode
	data []string
}

func (fd *flatDriver) readData() (bool, error) {
	if !fd.barcode.valid {
		return true, notValid
	}

	file, err := os.OpenFile(config.dbPath, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	check(err)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)
	notCount := 1

	for _, record := range records {
		if strings.Contains(strings.Join(record, " "), fd.barcode.code) {
			fd.data = record
		} else if fd.barcode.code == "end" && session.endless {
			log.Println("Scanned end code, exiting!")
			return false, exitCode
		} else if !strings.Contains(strings.Join(record, " "), fd.barcode.code) {
			notCount++
		}
	}

	if notCount > len(records) {
		fmt.Println("This code is not stored in the system.")

		return true, notFound
	}

	return true, nil


}

func (fd *flatDriver) writeData() (bool, error) {
	if !fd.barcode.valid {
		return true, notValid
	}

	file, err := os.OpenFile(config.dbPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)

	newRecords := getProductData()

	for i := 0; i < len(records); i++ {
		record := strings.Join(records[i], "")
		if strings.Contains(record, fd.barcode.code) && !session.edit {
			return true, alreadyExists
		}
	}
	if !session.edit {
		records = append(records, fd.data)
	} else if session.edit {
		for i := 0; i < len(records); i++ {
			if records[i][0] == fd.barcode.code {
				switch itemEditMenu() {
				case "1":
					records[i][1] = newRecords[0]
				case "2":
					records[i][2] = newRecords[1]
				case "3":
					records[i][3] = newRecords[2]
				}
			}
		}
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	check(err)

	return true, nil
}

func (fd *flatDriver) editData() (bool, error) {
	if !fd.barcode.valid {
		return true, notValid
	}

	_, err := fd.readData()

	if errors.Is(err, notFound) {
		fmt.Println("This Code was not found.\nPlease Add it first.")
	} else if errors.Is(err, notValid) {
		fmt.Println("This Code is not valid.")
	} else {
		data := fd.getData()
		itemDisplay(data[0], data[1], data[2], data[3])
	}

	newData := getProductData()

	fd.data[1] = newData[1]
	fd.data[2] = newData[2]
	fd.data[3] = newData[3]

	_, err = fd.writeData()

	return true, nil

}

func (fd *flatDriver) deleteData() (bool, error) {
	// TODO
}

func (fd *flatDriver) getData() []string {
	return fd.data
}

func newFlatDriver() *flatDriver {
	return &flatDriver {
		validateBarcode(getBarcode()),
		nil,
	}
}