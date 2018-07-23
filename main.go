package main

import (
	"conductor/drivers"
	"conductor/services/maestro"
	"conductor/tests"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Subcommands
	maestroCmd := flag.NewFlagSet("maestro", flag.ExitOnError)
	workerCmd := flag.NewFlagSet("worker", flag.ExitOnError)

	// Flags for maestro
	const (
		fileDefault    = ""
		fileUsage      = "File to submit.\nUsage: -file /path/to/file"
		jobNameDefault = ""
		jobNameUsage   = "Job name to submit\nUsage: -job-name calculate1"
		sourceDefault  = ""
		sourceUsage    = "Source to load job,\nUsage: -source-type file"
	)
	fileInput := maestroCmd.String("file", fileDefault, fileUsage)
	jobNameInput := maestroCmd.String("job-name", jobNameDefault, jobNameUsage)
	sourceInput := maestroCmd.String("source-type", sourceDefault, sourceUsage)

	// Make sure that subcommand was provided
	if len(os.Args) < 2 {
		fmt.Println("worker or maestro subcommand is required.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "worker":
		workerCmd.Parse(os.Args[2:])
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
		if *jobNameInput == "" {
			maestroCmd.PrintDefaults()
			os.Exit(1)
		}
		if *sourceInput == "" {
			maestroCmd.PrintDefaults()
			os.Exit(1)
		}

		var m maestro.Maestro
		m = maestro.NewService()

		job, err := m.LoadJob(*fileInput, *sourceInput)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		jobContent, err := m.Deploy(*jobNameInput, job)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(jobContent)

	}

	if workerCmd.Parsed() {
		callTestModel := tests.CallTest

		shell := drivers.Shell{}
		res, err := shell.Deploy(callTestModel)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(res)

	}
}
