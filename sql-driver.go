package main

type sqlDriver struct {
	barcode *barcode
	data []string

<<<<<<< HEAD
<<<<<<< HEAD
func (sd *sqlDriver) getData() (bool, error, []string) {
=======
func (sd sqlDriver) readData() (bool, error) {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
=======
func (sd sqlDriver) readData() (bool, error) {
>>>>>>> 8443cf764a3a0fae86ce56b70563045e10ed045a
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
=======
func newSqlDriver() *sqlDriver {
	return &sqlDriver {
>>>>>>> 8443cf764a3a0fae86ce56b70563045e10ed045a
		validateBarcode(getBarcode()),
		nil,
	}
}