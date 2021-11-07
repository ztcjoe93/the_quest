package area

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Position struct {
	X, Y int
}

type Grid struct {
	areaCode                string
	matrix                  [][]Tile
	startingPos, currentPos Position
}

func CreateGrid(tiles []*Tile, maxX int, maxY int, pos Position) Grid {
	if len(tiles) == 0 {
		log.Fatal("No tiles provided in slice.")
	}

	matrix := make([][]Tile, maxY)
	for row := range matrix {
		matrix[row] = make([]Tile, maxX)
	}

	for _, tile := range tiles {
		matrix[tile.coordinates.Y][tile.coordinates.X] = *tile
	}

	return Grid{
		areaCode:    tiles[0].areaCode,
		matrix:      matrix,
		startingPos: Position{pos.X, pos.Y},
		currentPos:  Position{pos.X, pos.Y},
	}
}

func (grid *Grid) GetGridData() {
	for ind, row := range grid.matrix {
		fmt.Println("Currently at row ", ind)
		for colind, col := range row {
			if col.areaCode == "" {
				fmt.Printf("%v,%v\n", ind, colind)
			} else {
				fmt.Printf("%v\n", col.coordinates)
			}
		}
	}
}

func (grid *Grid) GetCurrentPos() {
	fmt.Printf("Current position: (%v,%v)\n", grid.currentPos.X, grid.currentPos.Y)
}

func (grid *Grid) MoveToTile(tile Tile) {
	grid.currentPos = tile.coordinates
}

/*
	Get accessible tiles based on the current position that the player is on the
	grid.

	Algorithm works where (on a grid of hexagon tiles):
		- tiles on even x-axis can access adjacent tiles with same/-1 y-axis
		- tiles on odd x-axis can access adjacent tiles with same/+1 y-axis

	Why hexagons? Because hexagons are the *_bestagons_*
*/
func (grid *Grid) GetAccessibleTiles() []Tile {
	fmt.Println(grid.currentPos)
	curPosEven := grid.currentPos.X%2 == 0

	accessible := make([]Tile, 0)
	x := grid.currentPos.X
	y := grid.currentPos.Y

	if curPosEven {
		// accounting for left-side edge case
		if x > 0 {
			if grid.matrix[y][x-1].areaCode != "" {
				accessible = append(accessible, grid.matrix[y][x-1])
			}

			// accounting for top-edge case
			if y > 0 {
				if grid.matrix[y-1][x-1].areaCode != "" {
					accessible = append(accessible, grid.matrix[y-1][x-1])
				}
			}
		}

		// accounting for right-side edge case
		if x+1 < len(grid.matrix[y]) {
			if grid.matrix[y][x+1].areaCode != "" {
				accessible = append(accessible, grid.matrix[y][x+1])
			}

			// accounting for top-edge case
			if y > 0 {
				if grid.matrix[y-1][x+1].areaCode != "" {
					accessible = append(accessible, grid.matrix[y-1][x+1])
				}
			}
		}
	} else {
		if grid.matrix[y][x-1].areaCode != "" {
			accessible = append(accessible, grid.matrix[y][x-1])
		}

		// accounting for bottom-side edge case
		if y+1 != len(grid.matrix) {
			if grid.matrix[y+1][x-1].areaCode != "" {
				accessible = append(accessible, grid.matrix[y+1][x-1])
			}
		}

		// accounting for right-side edge case
		if x+1 < len(grid.matrix[y]) {
			if grid.matrix[y][x+1].areaCode != "" {
				accessible = append(accessible, grid.matrix[y][x+1])
			}

			if y+1 != len(grid.matrix) {
				if grid.matrix[y+1][x+1].areaCode != "" {
					accessible = append(accessible, grid.matrix[y+1][x+1])
				}
			}
		}
	}

	// rest of the common checks for top-y and bottom-y cases
	if y > 0 {
		if grid.matrix[y-1][x].areaCode != "" {
			accessible = append(accessible, grid.matrix[y-1][x])
		}
	}
	if y+1 < len(grid.matrix) {
		if grid.matrix[y+1][x].areaCode != "" {
			accessible = append(accessible, grid.matrix[y+1][x])
		}
	}

	fmt.Printf("%d tile(s) to move to: \n", len(accessible))
	/*
		for _, tile := range accessible {
			fmt.Printf("%v,", tile.coordinates)
		}
	*/

	return accessible
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
	coordinates      Position
}

/*
	Takes a row from the database and create a Tile using the provided data

*/
func CreateTile(content string, areaCode string, monsterEncounter bool,
	x int, y int) (*Tile, error) {

	if x < 0 || y < 0 {
		log.Errorf("(%v,%v) coordinates provided contains negative value\n", x, y)
		return nil, errors.New("negative XY Coordinates")
	}

	return &Tile{
		content:          content,
		areaCode:         areaCode,
		monsterEncounter: monsterEncounter,
		coordinates:      Position{x, y},
	}, nil
}

func (tile *Tile) GetCoordinates() Position {
	return tile.coordinates
}
