package main

type kvDriver struct {
	barcode *barcode
	data []string
}

<<<<<<< HEAD
func (kd *kvDriver) getData() (bool, error, []string) {
=======
func (kd kvDriver) readData() (bool, error) {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
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
func newKvDriver() *kvDriver {
	return &kvDriver {
=======
func (kd kvDriver) getData() []string {
	return kd.data
}

func newKvDriver() kvDriver {
	return kvDriver {
>>>>>>> aad3ea0565caa4cd852ff693c54c350352e27ef3
		validateBarcode(getBarcode()),
		nil,
	}
}
