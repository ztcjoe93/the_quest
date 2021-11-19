package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"the_quest/internal/area"
	"the_quest/internal/character"
	"the_quest/internal/database"

	log "github.com/sirupsen/logrus"
)

func init() {
	f, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)

	log.SetOutput(wrt)
	log.SetFormatter(&log.JSONFormatter{})

}

func main() {
	ch := character.CreateCharacter("test")
	fmt.Println(ch)

	db := database.Init()
	test01Tiles, maxX, maxY := db.GetTiles("test01")
	startPosX, startPosY := db.GetStartingPosition("test01")

	myGrid := area.CreateGrid(test01Tiles, maxX, maxY,
		area.Position{X: startPosX, Y: startPosY})

	quit := false
	for !quit {

		paths := myGrid.GetAccessibleTiles()
		for index, tile := range paths {
			fmt.Printf("%v %v\n", index+1, tile.GetCoordinates())
		}

		fmt.Print("\nChoose a path: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		switch input.Text() {
		case "c":
			myGrid.GetCurrentPos()
		case "m":
			myGrid.PrintGrid()
		case "q":
			quit = true
		default:
			value, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Fatal("Error converting text to int: %v", err)
			}
			if value > 0 && value <= len(paths) {
				myGrid.MoveToTile(paths[value-1])
			}
		}
	}
}
