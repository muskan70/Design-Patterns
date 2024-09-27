package main

import "errors"

const (
	TeamA = "A"
	TeamB = "B"
)

type IDress interface {
	getColor() string
}

type TeamADress struct {
	color string
}

func NewTeamADress() *TeamADress {
	return &TeamADress{color: "red"}
}

func (a *TeamADress) getColor() string {
	return a.color
}

type TeamBDress struct {
	color string
}

func NewTeamBDress() *TeamBDress {
	return &TeamBDress{color: "blue"}
}

func (b *TeamBDress) getColor() string {
	return b.color
}

type DressFactory struct {
	dressMap map[string]IDress
}

var dressFractoryInstance = &DressFactory{
	dressMap: make(map[string]IDress),
}

func getDressFactory() *DressFactory {
	return dressFractoryInstance
}

func (d *DressFactory) getDressByTeamName(teamName string) (IDress, error) {
	dress, ok := d.dressMap[teamName]
	if ok {
		return dress, nil
	}
	if teamName == TeamA {
		d.dressMap[teamName] = NewTeamADress()
		return d.dressMap[teamName], nil
	}

	if teamName == TeamB {
		d.dressMap[teamName] = NewTeamBDress()
		return d.dressMap[teamName], nil
	}

	return nil, errors.New("invalid team name")
}
