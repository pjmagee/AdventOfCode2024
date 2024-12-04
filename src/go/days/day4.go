package days

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	directions = [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top-left, Top, Top-right
		{0, -1}, {0, 1}, // Left, Right
		{1, -1}, {1, 0}, {1, 1}, // Bottom-left, Bottom, Bottom-right
	}
	word = "XMAS"
)

type Day4 struct {
}

func (d *Day4) TextToGrid(text string) [][]rune {

	lines := strings.Split(text, "\n")

	data := make([]string, 0)

	for _, line := range lines {
		if line != "" {
			data = append(data, line)
		}
	}

	grid := make([][]rune, len(lines))

	for i, line := range data {
		grid[i] = []rune(strings.ReplaceAll(line, " ", ""))
	}

	return grid
}

func searchWord(grid [][]rune, word string, row, col, dir int) (bool, string) {
	var cells []string

	for i := 0; i < len(word); i++ {

		newRow := row + i*directions[dir][0]
		newCol := col + i*directions[dir][1]

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[newRow]) {
			return false, ""
		}

		if grid[newRow][newCol] != rune(word[i]) {
			return false, ""
		}

		cells = append(cells, fmt.Sprintf("%d,%d", newRow, newCol))
	}

	return true, strings.Join(cells, "|")
}

func countWordMatches(grid [][]rune, word string) int {
	uniqueMatches := make(map[string]bool) // To store unique matches
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			for dir := 0; dir < len(directions); dir++ {
				if found, key := searchWord(grid, word, row, col, dir); found {
					if !uniqueMatches[key] {
						uniqueMatches[key] = true
						count++
					}
				}
			}
		}
	}
	return count
}

func findXMas(grid [][]rune) int {
	// Define directions for the X-shape arms
	dx := []int{1, 1, -1, -1}
	dy := []int{1, -1, 1, -1}

	count := 0

	// Iterate through the grid, avoiding edges
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			// The center of the X must be 'A'
			if grid[y][x] != 'A' {
				continue
			}

			// Check the four diagonal directions
			nxts := []rune{}
			valid := true
			for i := 0; i < 4; i++ {
				ny := y + dy[i]
				nx := x + dx[i]

				// Bounds check
				if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
					valid = false
					break
				}

				nxts = append(nxts, grid[ny][nx])
			}

			// Skip invalid configurations
			if !valid {
				continue
			}

			// Ensure all surrounding characters are 'M' or 'S' and satisfy the X conditions
			if (nxts[0] == 'M' || nxts[0] == 'S') &&
				(nxts[1] == 'M' || nxts[1] == 'S') &&
				(nxts[2] == 'M' || nxts[2] == 'S') &&
				(nxts[3] == 'M' || nxts[3] == 'S') &&
				nxts[0] != nxts[3] && // Top-left and bottom-right must differ
				nxts[1] != nxts[2] { // Top-right and bottom-left must differ
				count++
			}
		}
	}

	return count
}

type Response struct {
	Part1 int `json:"part_1"`
	Part2 int `json:"part_2"`
}

func (d *Day4) Execute(ctx context.Context, input string) (string, error) {

	grid := d.TextToGrid(input)

	response := Response{
		Part1: countWordMatches(grid, word),
		Part2: findXMas(grid),
	}

	result, _ := json.Marshal(response)

	return string(result), nil
}
