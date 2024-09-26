package main

import (
	"example.com/go_snips/builder"
	"example.com/go_snips/factory"
	"example.com/go_snips/liskov_substitution_principle"
	"example.com/go_snips/open_closed_principle"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [open_closed_principle | liskov_substitution_principle | builder | factory]")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case "open_closed_principle":
			open_closed_principle.TestOpenClosed()
		case "liskov_substitution_principle":
			liskov_substitution_principle.TestShapes()
		case "builder":
			builder.TestBuilder()
		case "factory":
			factory.TestFactory()
		default:
			fmt.Printf("Error: '%s' is not a valid option.\n", arg)
			fmt.Println("Usage: [open_closed_principle | liskov_substitution_principle | builder | factory]")
			os.Exit(1)
		}
	}
}
