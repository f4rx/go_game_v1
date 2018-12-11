package main

import (
	"errors"
	"fmt"
)

type Player struct {
	name string
}

type Position struct {
	X int
	Y int
}

func (p *Position) addX(i int) {
	p.X += i
}

func (p *Position) addY(i int) {
	p.Y += i
}

type GameMap struct {
	gameMap   [][]int
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
	case "right":
		fmt.Println("move right")
	case "down":
		fmt.Println("move down")
	case "up":
		fmt.Println("move up")
		playerPosition := g.positions[p]
		playerPosition.X--
		g.positions[p] = playerPosition
	default:
		return errors.New("Неопрделенное движение, " + direction)
	}
	return nil
}

func (g *GameMap) print() {
	fmt.Println(*g)
}

func main() {
	fmt.Println("fmt mock")
	player1 := Player{"player1"}
	player2 := new(Player)
	positions := make(map[*Player]Position)
	positions[&player1] = Position{2, 2}
	positions[player2] = Position{1, 1}
	gameMap := createGameMap(5, positions)
	gameMap.print()
	gameMap.movePlayer(&player1, "up")
	gameMap.movePlayer(player2, "down")
	gameMap.print()
}
