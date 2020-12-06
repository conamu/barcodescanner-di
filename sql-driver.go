package main

type sqlDriver struct {
	barcode *barcode
	data []string
}

func (sd sqlDriver) getData() (bool, error, []string) {
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

func newSqlDriver() sqlDriver {
	return sqlDriver {
		&barcode{
			"",
			false,
		},
		nil,
	}
}