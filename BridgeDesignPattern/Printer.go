package main

import "fmt"

type Printer interface {
	Print(f string)
}

type EpsonPrinter struct{}

func (e *EpsonPrinter) Print(f string) {
	fmt.Println("Printing by Epson Printer:", f)
}

type HpPrinter struct{}

func (h *HpPrinter) Print(f string) {
	fmt.Println("Printing by HP Printer:", f)
}
