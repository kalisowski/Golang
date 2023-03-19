package main

import (
	"fmt"
	"time"
)

type GameMap [][]bool

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	m := NewMap(20, 20)
	m[0][1] = true
	m[0][0] = true
	m[0][2] = true
	m[2][2] = true
	m[3][3] = true
	m[4][4] = true
	m[5][7] = true
	m[6][5] = true
	m[6][6] = true
	m[6][7] = true
	m[7][5] = true
	m[7][6] = true
	m[7][7] = true

	for {
		fmt.Print("\033[H\033[2J")
		clearScreen()
		PrintMap(m)
		m = m.NextState()
		time.Sleep(time.Second)
	}

}

func NewMap(width, height int) GameMap {
	m := make(GameMap, height)
	for i := range m {
		m[i] = make([]bool, width)
	}
	return m
}

func PrintMap(m GameMap) {
	for _, row := range m {
		for _, cell := range row {
			if cell {
				fmt.Print(" X ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}

func (m GameMap) LiveNeighbors(x, y int) int {
	neighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			row := y + i
			col := x + j
			if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
				continue
			}
			if m[row][col] {
				neighbors++
			}
		}
	}
	return neighbors
}

func (m GameMap) NextState() GameMap {
	next := NewMap(len(m[0]), len(m))
	for y, row := range m {
		for x, cell := range row {
			neighbors := m.LiveNeighbors(x, y)
			if cell {
				if neighbors == 2 || neighbors == 3 {
					next[y][x] = true
				}
			} else {
				if neighbors == 3 {
					next[y][x] = true
				} else {
					next[y][x] = false
				}
			}
		}
	}
	return next
}
