package main

type flatDriver struct {
	barcode *barcode
	data []string
}

func (fd flatDriver) getData() (bool, error, []string) {
	// TODO
}

func (fd flatDriver) writeData() (bool, error) {
	// TODO
}

func (fd flatDriver) editData() (bool, error) {
	// TODO
}

func (fd flatDriver) deleteData() (bool, error) {
	// TODO
}

func newFlatDriver() flatDriver {
	return flatDriver {
		&barcode{
			"",
			false,
		},
		nil,
	}
}