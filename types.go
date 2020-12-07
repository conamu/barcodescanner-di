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
var exitCode = errors.New("EXIT")

type dataInterface interface {
	getData() (bool, error , []string)
	writeData() (bool, error)
	editData() (bool, error)
	deleteData() (bool, error)
}

type codeSession struct {
	driver dataInterface
	barcode *barcode
	data []string
}

type barcode struct {
	code string
	valid bool
}

func newCodeSession(dataInterface dataInterface) *codeSession  {
	return &codeSession {
		dataInterface,
		&barcode{"", false},
		nil,
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var config = initConfig()

