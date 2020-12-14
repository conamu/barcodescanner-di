package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var session *codeSession
var notFound = errors.New("CNF")
var notValid = errors.New("NV")
var alreadyExists = errors.New("AE")
var exitCode = errors.New("EXIT")

type dataInterface interface {
	readData() (bool, error)
	writeData() (bool, error)
	editData() (bool, error)
	deleteData() (bool, error)
	getData() []string
}

type codeSession struct {
	driver dataInterface
	barcode *barcode
	endless bool
	edit bool
	delete bool
}

type barcode struct {
	code string
	valid bool
}

func newCodeSession(dataInterface dataInterface) *codeSession  {
	return &codeSession {
		dataInterface,
		&barcode{"", false},
		false,
		false,
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var config = initConfig()

