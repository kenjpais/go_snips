package main

import (
	"example.com/go_snips/builder"
	"example.com/go_snips/liskov_substitution_principle"
	"example.com/go_snips/open_closed_principle"
	"example.com/go_snips/factory"
)

func main() {
	builder.TestBuilder()
	open_closed_principle.TestOpenClosed()
	liskov_substitution_principle.TestShapes()
	factory.TestFactory()
}
