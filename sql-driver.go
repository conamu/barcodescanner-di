package main

type sqlDriver struct {
	barcode *barcode
	data []string
}

func (sd sqlDriver) readData() (bool, error) {
	// TODO
}

func (sd sqlDriver) writeData() (bool, error) {
	// TODO
}

func (sd sqlDriver) editData() (bool, error) {
	// TODO
}

func (sd sqlDriver) deleteData() (bool, error) {
	// TODO
}

func (sd sqlDriver) getData() []string {
	return sd.data
}

func newSqlDriver() sqlDriver {
	return sqlDriver {
		validateBarcode(getBarcode()),
		nil,
	}
}