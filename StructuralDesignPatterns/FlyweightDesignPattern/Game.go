package main

type Game struct {
	TeamAPlayers []*Player
	TeamBPlayers []*Player
}

func NewGame() *Game {
	game := &Game{}
	for i := 0; i < 4; i++ {
		game.TeamAPlayers = append(game.TeamAPlayers, NewPlayer(TeamA))
		game.TeamBPlayers = append(game.TeamBPlayers, NewPlayer(TeamB))
	}
	return game
}
