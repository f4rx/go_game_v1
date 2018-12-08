package main

import(
	"first_game/game"
)

func main(){
	error := game.Game()
	if error != nil {
		panic("Something broken")
	}
}
