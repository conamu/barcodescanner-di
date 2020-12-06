package main

type kvDriver struct {
	barcode *barcode
	data []string
}

func (kd kvDriver) getData() (bool, error, []string) {
	// TODO
}

func (kd kvDriver) writeData() (bool, error) {
	// TODO
}

func (kd kvDriver) editData() (bool, error) {
	// TODO
}

func (kd kvDriver) deleteData() (bool, error) {
	// TODO
}

func newKvDriver() kvDriver {
	return kvDriver {
		&barcode{
			"",
			false,
		},
		nil,
	}
}
