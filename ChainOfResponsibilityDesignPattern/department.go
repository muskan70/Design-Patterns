package main

type Department interface {
	Execute(p *Patient)
	SetNext(next Department)
}
