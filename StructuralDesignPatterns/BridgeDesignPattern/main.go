package main

import "fmt"

func main() {
	hpPrinter := &HpPrinter{}
	epsonPrinter := &EpsonPrinter{}

	macComputer := &Mac{}

	f := "Hello, I am a new file"

	macComputer.SetPrinter(hpPrinter)
	macComputer.PrintFile(f)
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.PrintFile(f)
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.PrintFile(f)
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.PrintFile(f)
	fmt.Println()
}
