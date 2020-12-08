package main

type kvDriver struct {
	barcode *barcode
	data []string
}

func (kd kvDriver) readData() (bool, error) {
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

func (kd kvDriver) getData() []string {
	return kd.data
}

func newKvDriver() kvDriver {
	return kvDriver {
		validateBarcode(getBarcode()),
		nil,
	}
}
