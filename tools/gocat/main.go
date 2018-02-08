package main

import (
    "fmt"
	"os"

    "github.com/openshift/origin/tools/gocat/pkg/cmd"
)

func main() {
	command := cmd.NewCmdGocat(os.Args[0], os.Stdout, os.Stderr)
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

