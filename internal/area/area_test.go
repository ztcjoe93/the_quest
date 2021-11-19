package area

import (
	"reflect"
	"testing"
)

/*
	Test to ensure that parameters passed into tile remains the same after
	initialization via constructor.
*/
func TestCreateTile(t *testing.T) {
	tile, _ := CreateTile("test", "areaCode01", false, 5, 4)

	if tile.content != "test" {
		t.Fatalf(`<content> argument provided to Tile does not match: 
		<value> = %v, <expected_value> = %v`, tile.content, "test")
	}

	if tile.areaCode != "areaCode01" {
		t.Fatalf(`<areaCode> argument provided to Tile does not match: 
		<value> = %v, <expected_value> = %v`, tile.areaCode, "areaCode01")
	}

	if tile.monsterEncounter != false {
		t.Fatalf(`<monsterEncounter> argument provided to Tile does not match: 
		<value> = %t, <expected_value> = %t`, tile.monsterEncounter, false)
	}

	// ensure that Position is constructed with passed x, y coordinates
	if !reflect.DeepEqual(Position{X: 5, Y: 4}, tile.coordinates) {
		t.Fatalf(`<x, y> argument provided to constructor does not match
		Position in Tile: <value> = %v, <expected_value> = %v`,
			tile.coordinates, "{5 4}")
	}
}

func TestCreateTileWithNegativeYPos(t *testing.T) {
	_, err := CreateTile("test", "areaCode01", false, 5, -1)
	if err == nil {
		t.Fatalf(`Negative coordinates provided but no error returned.
		<x_value> = %v, <y_value> = %v`, 5, -1)
	}
}

func TestCreateTileWithNegativeXPos(t *testing.T) {
	_, err := CreateTile("test", "areaCode01", false, -2, 3)
	if err == nil {
		t.Fatalf(`Negative coordinates provided but no error returned.
		<x_value> = %v, <y_value> = %v`, -2, 3)
	}
}

func TestTileGetCoordinatesFunction(t *testing.T) {
	tile, _ := CreateTile("test", "areaCode01", false, 4, 6)

	if !reflect.DeepEqual(Position{X: 4, Y: 6}, tile.GetCoordinates()) {
		t.Fatalf(`<x, y> argument provided to constructor does not match
		Position in Tile: <value> = %v, <expected_value> = %v`,
			tile.coordinates, "{4 6}")
	}
}
