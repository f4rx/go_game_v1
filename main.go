package main

import (
	"errors"
	"fmt"
	"os"
)

type Player struct {
	name string
}

type Position struct {
	X int
	Y int
}

type Game struct{
	gameMap *GameMap
	players []*Player
	turn *Player
}

func (p *Position) addX(i int) {
	p.X += i
}

func (p *Position) addY(i int) {
	p.Y += i
}

type GameMap struct {
	gameMap   [][]int
	positions map[*Player]*Position
}

func createGameMap(size int, positions map[*Player]*Position) GameMap {
	var gameMapArr = make([][]int, size)
	for i := range gameMapArr {
		gameMapArr[i] = make([]int, size)
	}
	gameMap := GameMap{gameMapArr, positions}
	return gameMap
}

func (g *GameMap) movePlayer(p *Player, direction string) error {
	switch direction {
	case "3":
		fmt.Println("move left")
	case "4":
		fmt.Println("move right")
	case "2":
		fmt.Println("move down")
	case "1":
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
	fmt.Println(g.gameMap)

	for k, v := range g.positions {
    	fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}

func (g *Game) gameAction(str *string) error {
	switch *str {
	case "exit":
		fmt.Println("goodbuy")
		os.Exit(0)
	case "1", "2", "3", "4":
		g.gameMap.movePlayer(g.turn, *str)
		fmt.Println("move")
		//func createGameMap(size int, positions map[*Player]*Position) GameMap {
	case "5":
		g.gameMap.print()
		fmt.Println("print fucntion ")
		//func (g *GameMap) print() {
	default:
		return nil
	}
	return nil
}

func main() {
	fmt.Println("fmt mock")
	player1 := Player{"player1"}
	player2 := new(Player)
	player2.name = "player2"
	// fmt.Println(player1)
	// fmt.Println(player2)
	positions := make(map[*Player]*Position)
	positions[&player1] = &Position{2, 2}
	positions[player2] = &Position{1, 1}
	gameMap := createGameMap(5, positions)
	gameMap.print()
	gameMap.movePlayer(&player1, "up")
	gameMap.movePlayer(player2, "down")
	gameMap.print()

	game := new(Game)
	game.players = append(game.players, &player1)
	game.players = append(game.players, player2)
	game.gameMap = &gameMap
	game.turn = &player1

    var input string

	for ;; {
		fmt.Printf("Ход игрока: %s\n", player1.name)
		fmt.Printf("Выберите направление: 1. Вверх 2. Вниз 3. Влево 4. Вправо 5. Печать карты\n")
		fmt.Scanln(&input)
		game.gameAction(&input)
		// chooseGameAction(&input)
		// fmt.Print(input)
		// err := gameMap.movePlayer(&player1, input)
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }

	}
}
