package main

type Player struct {
	PlayerId int
	TeamName string
	Dress    IDress
	Xpos     int
	Ypos     int
}

var x = 0

func NewPlayer(teamName string) *Player {
	dress, _ := getDressFactory().getDressByTeamName(teamName)
	x++
	return &Player{
		PlayerId: x,
		TeamName: teamName,
		Dress:    dress,
	}
}

func (p *Player) displayLocation(x, y int) {
	p.Xpos = x
	p.Ypos = y
}
