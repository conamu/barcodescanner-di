package main

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type flatDriver struct {
	barcode *barcode
	data []string
}

func (fd *flatDriver) getData() (bool, error, []string) {
	if !fd.barcode.valid {
		return true, notValid, nil
	}

	var dataRow []string
	file, err := os.OpenFile(config.dbPath, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	check(err)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)
	notCount := 1

	for _, record := range records {
		if strings.Contains(strings.Join(record, " "), fd.barcode.code) {
			dataRow = record
		} else if fd.barcode.code == "end" {
			log.Println("Scanned end code, exiting!")
			return false, exitCode, nil
		} else if !strings.Contains(strings.Join(record, " "), fd.barcode.code) {
			notCount++
		}
	}

	if notCount > len(records) {
		fmt.Println("This code is not stored in the system.")

		return true, notFound, nil
	}

	return true, nil, dataRow


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

	records = append(records, fd.data)

	writer := csv.NewWriter(file)
	writer.WriteAll(records)

	return true, nil
}

func (fd *flatDriver) editData() (bool, error) {
	if !fd.barcode.valid {
		return true, notValid
	}

	

}

func (fd *flatDriver) deleteData() (bool, error) {
	// TODO
}

func newFlatDriver() *flatDriver {
	return &flatDriver {
		validateBarcode(getBarcode()),
		nil,
	}
}