package area

import (
	"fmt"

	err "the_quest/internal/error"

	log "github.com/sirupsen/logrus"
)

type Position struct {
	X, Y int
}

type Grid struct {
	areaCode                string
	matrix                  [][]Tile
	startingPos, currentPos Position
	maxPos                  Position
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
		maxPos:      Position{maxX, maxY},
	}
}

/*
	Prints a map of all tiles onto the console,
	currently a debugging/map-show tool but can be used for character
	position indication if needed later on.

	Algorithm calculates the exact position where the hexagon is to be
	generated and populates the 2d matrix with the data if a hexagon tile
	is found within the Grid.matrix
*/
func (grid *Grid) PrintGrid() {
	picture := make([][]string, (grid.maxPos.Y*2)+2)

	for row := range picture {
		picture[row] = make([]string, (grid.maxPos.X*3)+1)
		for cell := range picture[row] {
			picture[row][cell] = " "
		}
	}

	// base ehxagon
	hexagon := make([][]string, 3)
	for row := range hexagon {
		hexagon[row] = make([]string, 4)
	}

	// hexagon shapes in the matrix
	hexagon[0][0] = " "
	hexagon[0][1] = "_"
	hexagon[0][2] = "_"
	hexagon[0][3] = " "
	hexagon[1][0] = "/"
	hexagon[1][1] = " "
	hexagon[1][2] = " "
	hexagon[1][3] = "\\"
	hexagon[2][0] = "\\"
	hexagon[2][1] = "_"
	hexagon[2][2] = "_"
	hexagon[2][3] = "/"

	for ind, row := range grid.matrix {
		for colInd, col := range row {
			if col.areaCode != "" {
				xStart := colInd * 3
				//xEnd := xStart + 3
				yStart := ind*2 + (colInd % 2)
				//yEnd := yStart + 2

				for i := 0; i < 3; i++ {
					for j := 0; j < 4; j++ {
						if picture[i+yStart][j+xStart] == " " {
							picture[i+yStart][j+xStart] = hexagon[i][j]
						}
					}
				}
			}
		}
	}

	for ind, row := range picture {
		for colInd, _ := range row {
			fmt.Printf("%v", picture[ind][colInd])
		}
		fmt.Printf("\n")
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
		return nil, err.ErrNegativeValue
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
