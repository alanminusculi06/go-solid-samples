package main

// Interface Segregation Principle

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {
	// ...
}

func (m MultiFunctionPrinter) Fax(d Document) {
	// ...
}

func (m MultiFunctionPrinter) Scan(d Document) {
	// ...
}

type DefaultPrinter struct{}

func (o DefaultPrinter) Print(d Document) {
	// ...
}

func (o DefaultPrinter) Fax(d Document) {
	panic("operation not supported")
}

func (o DefaultPrinter) Scan(d Document) {
	panic("operation not supported")
}

// better approach
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type BasicPrinter struct{}

func (m BasicPrinter) Print(d Document) {
	// ...
}

type MultiFunctionPrinter2 struct{}

func (p MultiFunctionPrinter2) Scan(d Document) {
	// ...
}

func (p MultiFunctionPrinter2) Print(d Document) {
	// ...
}

func main() {

}
