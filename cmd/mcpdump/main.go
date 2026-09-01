package main

import (
	"os"

	"github.com/baldaworks/mcpdump/cmd/mcpdump/cmd"
)

func main() {
	if err := command.Command().Execute(); err != nil {
		os.Exit(1)
	}
}
