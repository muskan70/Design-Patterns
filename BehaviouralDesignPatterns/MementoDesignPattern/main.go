package main

import "fmt"

func main() {
	originator := &Originator{}
	caretaken := &Caretaker{}

	originator.SetState("A")
	caretaken.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	caretaken.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	caretaken.AddMemento(originator.CreateMemento())

	fmt.Println(originator.GetState())

	originator.RestoreMemento(caretaken.Undo())
	fmt.Println(originator.GetState())

	originator.RestoreMemento(caretaken.GetMemento(0))
	fmt.Println(originator.GetState())

}
