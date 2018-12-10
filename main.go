package main

import(
	"fmt"
	"errors"
)

type Player struct {
	name string
}

type Position struct {
	X int
	Y int
}

type GameMap struct {
	gameMap [][]int
	positions map[*Player]Position
}

func createGameMap(size int, positions map[*Player]Position) GameMap {
	var gameMapArr = make([][]int, size)
	for i := range gameMapArr {
    	gameMapArr[i] = make([]int, size)
	}
	gameMap := GameMap{gameMapArr, positions}
	return gameMap
}

func (g *GameMap) movePlayer(p *Player, direction string) error {
	switch direction {
	case "left":
		fmt.Println("move left")
		g.positions[p].X -= 1
	case "right":
		fmt.Println("move right")
	case "down":
		fmt.Println("move down")
	case "up":
		fmt.Println("move up")
	default:
		return errors.New("Непрделенное движение, " + direction)
	}
	return nil
}

func (g *GameMap) print() {
	fmt.Println(*g)
}

func main(){
	fmt.Println("fmt mock")
	player1 := Player{}
	player2 := new(Player)
	positions := make(map[*Player]Position)
	positions[&player1] = Position{2,2}
	positions[player2] = Position{1,1}
	gameMap := createGameMap(5, positions)
	gameMap.print()
	gameMap.movePlayer(&player1, "up")
	gameMap.movePlayer(player2, "down")
}
