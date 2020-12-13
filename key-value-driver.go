package main

type kvDriver struct {
	barcode *barcode
	data []string
}

<<<<<<< HEAD
<<<<<<< HEAD
func (kd *kvDriver) getData() (bool, error, []string) {
=======
func (kd kvDriver) readData() (bool, error) {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
=======

func (kd kvDriver) readData() (bool, error) {
>>>>>>> 8443cf764a3a0fae86ce56b70563045e10ed045a
	// TODO
}

func (kd *kvDriver) writeData() (bool, error) {
	// TODO
}

func (kd *kvDriver) editData() (bool, error) {
	// TODO
}

func (kd *kvDriver) deleteData() (bool, error) {
	// TODO
}

<<<<<<< HEAD
<<<<<<< HEAD
func newKvDriver() *kvDriver {
	return &kvDriver {
=======
func (kd kvDriver) getData() []string {
	return kd.data
}

func newKvDriver() kvDriver {
	return kvDriver {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
=======
func newKvDriver() *kvDriver {
	return &kvDriver {
>>>>>>> 8443cf764a3a0fae86ce56b70563045e10ed045a
		validateBarcode(getBarcode()),
		nil,
	}
}
