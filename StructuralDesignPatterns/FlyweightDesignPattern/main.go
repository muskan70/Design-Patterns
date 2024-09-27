package main

import "fmt"

func main() {
	game := NewGame()
	for i := 0; i < len(game.TeamAPlayers); i++ {
		fmt.Println(game.TeamAPlayers[i].PlayerId, game.TeamAPlayers[i].TeamName, game.TeamAPlayers[i].Dress.getColor())
	}
	for i := 0; i < len(game.TeamBPlayers); i++ {
		fmt.Println(game.TeamBPlayers[i].PlayerId, game.TeamBPlayers[i].TeamName, game.TeamBPlayers[i].Dress.getColor())
	}
}
