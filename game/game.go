package game

import(
	"fmt"
	// "os"
)

// PrintText foobarcom
func PrintText(s string){
	fmt.Println(s)
}

func createGameMap(size int,) [][]int {
	var gameMap = make([][]int, size)
	for i := range gameMap {
    	gameMap[i] = make([]int, size)
	}
	return gameMap
}

func printMap(gameMap *[][]int){
	for i := range *gameMap {
    	fmt.Println(i, *gameMap[i])
	}
}

// Game foobarcom
func Game() error {
	gameMap := createGameMap(5)
	fmt.Println(gameMap)


	// var name string
	// fmt.Fscan(os.Stdin, &name)
	// fmt.Println(name)

	return nil
}
