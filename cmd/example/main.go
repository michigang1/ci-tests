package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	lab2 "github.com/michigang1/ci-tests"
)

var (
	cmdInput  = flag.String("e", "", "Expression to compute")
	fileInput = flag.String("f", "", "Reading from file")
	saveInput = flag.String("o", "", "Saving to file")
)

func main() {
	flag.Parse()

	handler := &lab2.ComputeHandler{}
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
		defer file.Close()
	}

	if *saveInput != "" {
		file, err := os.Create(*saveInput)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		handler.Output = file
	}

	if handler.Output == nil {
		handler.Output = os.Stdout
	}

	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
	}

	res, _ := lab2.PostfixToInfix("2 2 +")
	fmt.Println(res)
}
