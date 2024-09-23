package main

import "fmt"

type Computer interface {
	PrintFile(f string)
	SetPrinter()
}

type Mac struct {
	printer Printer
}

func (m *Mac) PrintFile(f string) {
	fmt.Println("PRINT REQUEST BY MAC")
	m.printer.Print(f)
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) PrintFile(f string) {
	fmt.Println("PRINT REQUEST BY WINDOWS")
	w.printer.Print(f)
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
