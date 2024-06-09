package main

import (
	"flag"
	"fmt"
	"os"

	newcli "github.com/dynonguyen/new-cli"
)

func main() {
	var verbose bool

	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No paths provided\n")
		os.Exit(1)
	}

	newcli.NewCli(args, verbose)
}
