package game

import(
	"fmt"
	"strconv"
	"errors"

	// "os"
)

type Position struct {
	positionX int
	positionY int
}

type Player struct {
    level int
	position Position
}

func (p *Player) move(direction string) error {
	switch direction {
	case "left":
		fmt.Println("move left")
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

type GameMap struct {
	gameMap [][]int
	players []Player
}

// PrintText foobarcom
func PrintText(s string){
	fmt.Println(s)
}

func createGameMap(size int, players []Player) GameMap {
	var gameMapArr = make([][]int, size)
	for i := range gameMapArr {
    	gameMapArr[i] = make([]int, size)
	}
	gameMap := GameMap{gameMapArr, players}
	return gameMap
}

func (gameMap *GameMap) print() {
	for x,b := range gameMap.gameMap {
		for y := range b{
			char := "x"
			if player.position.positionX == x && player.position.positionY == y {
				char = "P"
			} else if b[y] == 1 {
				char = "R"
			} else if b[y] != 0 {
				char = strconv.Itoa(b[y])
			}
			fmt.Print(char, " ")
		}
		fmt.Println()
	}
}

// func printMap(gameMap [][]int, player Player){
// 	for x,b := range gameMap {
// 		for y := range b{
// 			char := "x"
// 			if player.position.positionX == x && player.position.positionY == y {
// 				char = "P"
// 			} else if b[y] == 1 {
// 				char = "R"
// 			} else if b[y] != 0 {
// 				char = strconv.Itoa(b[y])
// 			}
// 			fmt.Print(char, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// Game foobarcom
func Game() error {
	mapSize := 5

	var players []Player
	player := Player{level: 0, position: Position{mapSize / 2 , mapSize / 2}}
	players = append(players, player)

	gameMap := createGameMap(mapSize, players)
	printMap(gameMap, player)

	// var name string
	// fmt.Fscan(os.Stdin, &name)
	// fmt.Println(name)

	return nil
}
