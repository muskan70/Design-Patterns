package main

type Memento struct {
	State string
}

func (m *Memento) GetSavedState() string {
	return m.State
}
