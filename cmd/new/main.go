package main

import (
	"fmt"
	"os"

	newcli "github.com/dynonguyen/go-cli-utils/internal/new"
)

func main() {
	args, verbose := newcli.GetArgs()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No paths provided\n")
		os.Exit(1)
	}

	newcli.NewCli(args, verbose)
}
