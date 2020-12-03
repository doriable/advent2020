package main

import (
	"io/ioutil"
	"os"
)

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
