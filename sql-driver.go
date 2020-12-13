package main

type sqlDriver struct {
	barcode *barcode
	data []string
}

<<<<<<< HEAD
func (sd *sqlDriver) getData() (bool, error, []string) {
=======
func (sd sqlDriver) readData() (bool, error) {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
	// TODO
}

func (sd *sqlDriver) writeData() (bool, error) {
	// TODO
}

func (sd *sqlDriver) editData() (bool, error) {
	// TODO
}

func (sd *sqlDriver) deleteData() (bool, error) {
	// TODO
}

<<<<<<< HEAD
func newSqlDriver() *sqlDriver {
	return &sqlDriver {
=======
func (sd sqlDriver) getData() []string {
	return sd.data
}

func newSqlDriver() sqlDriver {
	return sqlDriver {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
		validateBarcode(getBarcode()),
		nil,
	}
}