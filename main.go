package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/fatih/color"

)

type Player struct {
	name string
	color string
}

type Position struct {
	X int
	Y int
}

type Game struct {
	gameMap *GameMap
	players []*Player
	turn    *Player
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
		if g.positions[p].X == 0 {
			g.positions[p].X = len(g.gameMap) - 1
		} else {
			g.positions[p].X--
		}
		g.print()
	case "4":
		if g.positions[p].X == len(g.gameMap)-1 {
			g.positions[p].X = 0
		} else {
			g.positions[p].X++
		}
		g.print()
	case "2":
		if g.positions[p].Y == len(g.gameMap)-1 {
			g.positions[p].Y = 0
		} else {
			g.positions[p].Y++
		}
		g.print()
	case "1":
		if g.positions[p].Y == 0 {
			g.positions[p].Y = len(g.gameMap) - 1
		} else {
			g.positions[p].Y--
		}
		g.print()
	default:
		return errors.New("Неопрделенное движение, " + direction)
	}
	return nil
}

func getColor(colorText string) func(a ...interface{}) string {
	switch colorText {
	case "red":
		return color.New(color.FgRed).SprintFunc()
	case "yellow":
		return color.New(color.FgYellow).SprintFunc()
	default:
		return color.New(color.FgBlack).SprintFunc()
	}

}


func (g *GameMap) print() {
	var positions []*Position

	for _, v := range g.positions {
		positions = append(positions, v)
	}

	for y := 0; y < len(g.gameMap); y++ {
		for x := 0; x < len(g.gameMap[y]); x++ {
			colorText := getColor("black")
			char := "x"
			for player, position := range g.positions {
				if x == position.X && y == position.Y {

					// fmt.Printf("this is a %s and this is %s.\n", yellow("warning"), red("error"))
					char = "P"
					colorText = getColor(player.color)
				}
			}
			fmt.Printf("%s ", colorText(char))
		}
		fmt.Println()
	}

	// for k, v := range g.positions {
	// 	fmt.Printf("key[%s] value[%s]\n", k, v)
	// }
	// fmt.Printf("%#v", g.positions)

}

func (g *Game) gameAction(str string) error {
	switch str {
	case "exit":
		fmt.Println("goodbuy")
		os.Exit(0)
	case "1", "2", "3", "4":
		g.gameMap.movePlayer(g.turn, str)
		// fmt.Println("move")
	case "5":
		g.gameMap.print()
	default:
		return nil
	}
	return nil
}

func main() {
	fmt.Println("fmt mock")
	player1 := Player{"player1", "red"}
	player2 := new(Player)
	player2.name = "player2"
	player2.color = "yellow"
	// fmt.Println(player1)
	// fmt.Println(player2)
	positions := make(map[*Player]*Position)
	positions[&player1] = &Position{2, 2}
	positions[player2] = &Position{1, 1}
	gameMap := createGameMap(5, positions)
	// gameMap.print()
	// gameMap.movePlayer(&player1, "up")
	// gameMap.movePlayer(player2, "down")
	// gameMap.print()

	game := new(Game)
	game.players = append(game.players, &player1)
	game.players = append(game.players, player2)
	game.gameMap = &gameMap
	game.turn = &player1
	game.gameMap.print()


	var input string

	for {
		fmt.Printf("Ход игрока: %s\n", player1.name)
		fmt.Printf("Выберите направление: 1. Вверх 2. Вниз 3. Влево 4. Вправо 5. Печать карты\n")
		fmt.Scanln(&input)
		game.gameAction(input)
		// chooseGameAction(&input)
		// fmt.Print(input)
		// err := gameMap.movePlayer(&player1, input)
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }

	}
}
