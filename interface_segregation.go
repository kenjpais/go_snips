package main

type Document struct{}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxer interface {
	Fax(d Document)
}

// because photocopier can print as well as scan
type Photocopier struct {
	Scanner
	Printer
}

func (p *Photocopier) Print(d Document) {
	p.Printer.Print(d)
}

func (p *Photocopier) Scan(d Document) {
	p.Scanner.Scan(d)
}

func Fax(d Document) {
}

func test() {
	var d Document
	ph := Photocopier{}
	ph.Print(d)
}
