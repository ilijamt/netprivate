package main

import (
	"github.com/ilijamt/netprivate/cmd/is-private/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
