package main

import (
	"conductor/services/maestro"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Subcommands
	maestroCmd := flag.NewFlagSet("maestro", flag.ExitOnError)
	conductorCmd := flag.NewFlagSet("conductor", flag.ExitOnError)

	// Flags for maestro
	const (
		fileDefault   = ""
		fileUsage     = "File to submit.\nUsage: -file /path/to/file"
		regionDefault = ""
		regionUsage   = "Region to submit\nUsage: -region us-west-1"
	)
	fileInput := maestroCmd.String("file", fileDefault, fileUsage)
	regionInput := maestroCmd.String("region", regionDefault, regionUsage)

	// Make sure that subcommand was provided
	if len(os.Args) < 2 {
		fmt.Println("conductor or maestro subcommand is required.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "conductor":
		conductorCmd.Parse(os.Args[2:])
	case "maestro":
		maestroCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if maestroCmd.Parsed() {
		if *fileInput == "" {
			maestroCmd.PrintDefaults()
			os.Exit(1)
		}
		if *regionInput == "" {
			maestroCmd.PrintDefaults()
			os.Exit(1)
		}

		var m maestro.Maestro
		m = maestro.NewService()

		m.Deploy(*regionInput, *fileInput)
	}
}
