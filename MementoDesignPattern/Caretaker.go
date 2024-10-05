package main

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) Undo() *Memento {
	c.mementos = c.mementos[:len(c.mementos)-1]
	return c.mementos[len(c.mementos)-1]
}

func (c *Caretaker) GetMemento(idx int) *Memento {
	if idx < len(c.mementos) {
		return c.mementos[idx]
	}
	return nil
}
