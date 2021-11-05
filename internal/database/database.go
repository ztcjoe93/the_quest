package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"the_quest/internal/area"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	database sql.DB
}

func Init() *Database {
	// initial loading of .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("db_username"), os.Getenv("db_password"), os.Getenv("db_name"),
	))

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to mysql.")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Database{database: *db}
}

// fetch all tiles by specified areaCode
func (database *Database) GetTiles(areaCode string) []*area.Tile {
	db := database.database
	var tiles []*area.Tile

	var (
		dbContent      string
		dbAreaCode     string
		dbMonEncounter bool
		dbX            int
		dbY            int
	)
	rows, err := db.Query(`SELECT content, area_code, mon_encounter, x, y 
		FROM tile WHERE area_code = ?`, areaCode)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&dbContent, &dbAreaCode, &dbMonEncounter, &dbX, &dbY)
		if err != nil {
			log.Fatal(err)
		}
		tile := area.CreateTile(dbContent, dbAreaCode, dbMonEncounter,
			dbX, dbY)
		tiles = append(tiles, tile)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return tiles
}
