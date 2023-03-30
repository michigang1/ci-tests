package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/michigang1/ci-tests"
)

var (
	cmdInput   = flag.String("e", "", "Expression to compute")
	fileInput  = flag.String("f", "", "Reading from file")
	saveOutput = flag.String("o", "", "Saving to file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	if *cmdInput != "" && *fileInput != "" {
		log.Fatal("Can't use -e and -f simultaneously!")
	}

	if *cmdInput != "" {
		handler.Input = strings.NewReader(*cmdInput)
	}

	if *fileInput != "" {
		file, err := os.Open(*fileInput)
		if err != nil {
			log.Fatal(err)
		}
		handler.Input = file
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
	}

	if *saveOutput != "" {
		file, err := os.Create(*saveOutput)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
		handler.Output = file
	}

	if handler.Output == nil {
		handler.Output = os.Stdout
	}

	if handler.Input == nil {
		handler.Input = os.Stdin
		log.Fatal("No input provided!")
	}

	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
	}
}
