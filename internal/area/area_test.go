package area

import (
	"reflect"
	"testing"
)

func TestCreateTile(t *testing.T) {
	tile, _ := CreateTile("test", "areaCode01", false, 5, 4)

	// ensure that content
	if tile.content != "test" {
		t.Fatalf(`<content> argument provided to Tile does not match: 
		<value> = %v, <expected_value> = %v`, tile.content, "test")
	}

	if tile.areaCode != "areaCode01" {
		t.Fatalf(`<areaCode> argument provided to Tile does not match: 
		<value> = %v, <expected_value> = %v`, tile.areaCode, "areaCode01")
	}

	// ensure that Position is constructed with passed x, y coordinates
	if !reflect.DeepEqual(Position{X: 5, Y: 4}, tile.coordinates) {
		t.Fatalf(`<x, y> argument provided to constructor does not match
		Position in Tile: <value> = %v, <expected_value> = %v`,
			tile.areaCode, "areaCode01")
	}
}

func TestCreateTileWithNegativeCoordinates(t *testing.T) {
	_, err := CreateTile("test", "areaCode01", false, 5, -1)
	if err == nil {
		t.Fatalf(`Negative coordinates provided but no error returned.
		<x_value> = %v, <y_value> = %v`, 5, -1)
	}
}
