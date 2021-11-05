package main

import (
	"fmt"
	"io"
	"os"
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
	fmt.Println("Main application entry point")

	ch := character.CreateCharacter("test")
	fmt.Println(ch)

	db := database.Init()
	test01Tiles := db.GetTiles("test01")

	for _, tile := range test01Tiles {
		fmt.Printf("%+v\n", tile)
	}
}
