package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpMess)
	}
	log.SetFlags(0)

	// Parse the args to be read by flag
	flag.Parse()

	err := execCommand()
	if err != nil {
		log.Fatal(err)
	}
}
