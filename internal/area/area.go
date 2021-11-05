package area

type Grid struct {
	areaCode string
	matrix   [][]Tile
}

/*
	A single Tile represented on a Grid.
	Content of the tile represents text to display on
	page when player reaches said tile.
	Some Tiles are specifically set to not trigger random
	encounters -- to put encounter on a PRD spread?
	Tile's position on a Grid is represented by the x,y
	co-ordinates.
*/
type Tile struct {
	content          string
	areaCode         string
	monsterEncounter bool
	x                int
	y                int
}

/*
	Takes a row from the database and create a Tile using the provided data

*/
func CreateTile(content string, areaCode string, monsterEncounter bool,
	x int, y int) *Tile {
	return &Tile{
		content:          content,
		areaCode:         areaCode,
		monsterEncounter: monsterEncounter,
		x:                x,
		y:                y,
	}
}
