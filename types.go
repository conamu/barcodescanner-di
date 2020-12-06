package main

import (
	"bufio"
	"log"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

type dataInterface interface {
	getData() (bool,error , []string)
	writeData() (bool, error)
	editData() (bool, error)
	deleteData() (bool, error)
}

type codeSession struct {
	dataInterface *dataInterface
	code string
	codeValid bool
	data []string
}

func newCodeSession(dataInterface *dataInterface) codeSession  {
	return codeSession {
		dataInterface,
		"",
		false,
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
