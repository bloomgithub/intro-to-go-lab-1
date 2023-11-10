package main

// Update world based on game of life rules
func calculateNextState(p golParams, world [][]byte) [][]byte {

	rows := len(world)
	cols := len(world[0])

	// Create temporary world to store new state
	tmp := make([][]byte, rows)
	for i := range tmp {
		tmp[i] = make([]byte, cols)
	}

	// Iterate through cells
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// Get the count of live neighbours
			liveNeighbors := calculateLiveNeighbours(world, rows, cols, i, j)

			// Apply rules
			if world[i][j] == 255 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					tmp[i][j] = 0 // Die unless 2-3 neighbors
				} else {
					tmp[i][j] = 255 // Continue living
				}
			} else {
				if liveNeighbors == 3 {
					tmp[i][j] = 255 // Birth if exactly 3 neighbors
				}
			}

		}
	}

	// Update world state
	world = tmp

	return world
}

func calculateLiveNeighbours(world [][]byte, rows, cols, i, j int) int {

	// Count of live neighbours
	count := 0

	// Loop through neighbouring cells
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {

			// Exclude self
			if x == 0 && y == 0 {
				continue
			}

			// Wrap row and col indexes
			row := (x + i + rows) % rows
			col := (y + j + cols) % cols

			// Count live cells
			if world[row][col] == 255 {
				count++
			}
		}
	}

	return count
}

func calculateAliveCells(p golParams, world [][]byte) []cell {

	alive := []cell{}

	rows := len(world)
	cols := len(world[0])

	// Iterate over world
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// Check if cell is alive
			// N.B columns first in test code
			if world[j][i] == 255 {
				alive = append(alive, cell{i, j})
			}
		}
	}

	return alive
}
