package main

type Prototype interface {
	Clone() Prototype
	Print(indent string)
}
