package main

import (
	"conductor/services/maestro"
)

func main() {
	var m maestro.Maestro
	m = maestro.NewService()

	m.Deploy("p02", "blueprint")
}
