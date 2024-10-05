package main

type Originator struct {
	State string
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) SetState(s string) {
	o.State = s
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{State: o.State}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.SetState(m.GetSavedState())
}
