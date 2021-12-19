package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a int, b int) int {
	if a>=b {
		return a
	}
	return b
}

func get_input(f *os.File) ([][]int, [][]int) {
	scanner := bufio.NewScanner(f)
	coords := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		coord := make([]int, 2)
		coord[0], _ = strconv.Atoi(strings.Split(line, ",")[0])
		coord[1], _ = strconv.Atoi(strings.Split(line, ",")[1])
		coords = append(coords, coord)
	}
	folds := make([][]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		fold := make([]int, 2)
		fold_direction := line[2][0]
		if fold_direction == 'x' {
			fold[0] = 0
		} else {
			fold[0] = 1
		}
		fold[1], _ = strconv.Atoi(line[2][2:])
		folds = append(folds, fold)
	}
	return coords, folds
}

func main() {
	f, _ := os.Open("input.txt")
	coords, folds := get_input(f)
	max_x := 0
	max_y := 0
	for _, coord := range coords {
		max_x = max(max_x, coord[0])
		max_y = max(max_y, coord[1])
	}

	for _, fold := range folds {
		new_coords := make([][]int, 0)
		if fold[0] == 0 {
			left := fold[1]
			right := max_x - fold[1]
			for _, coord := range coords {
				if left >= right {
					if coord[0] < fold[1] {
						new_coords = append(new_coords, coord)
					} else if coord[0] > fold[1] {
						coord[0] = 2*fold[1] - coord[0]
						new_coords = append(new_coords, coord)
					}
				} else {
					if coord[0] < fold[1] {
						coord[0] = 2*fold[1] - coord[0]
						coord[0] = max_x - coord[0]
						new_coords = append(new_coords, coord)
					} else if coord[0] > fold[1] {
						coord[0] = max_x - coord[0]
						new_coords = append(new_coords, coord)
					}
				}
			}
			if left >= right {
				max_x = left - 1
			} else {
				max_x = right - 1
			}
		} else {
			top := fold[1]
			bottom := max_y - fold[1]
			
			for _, coord := range coords {
				if top >= bottom {
					if coord[1] < fold[1] {
						new_coords = append(new_coords, coord)
					} else if coord[1] > fold[1] {
						coord[1] = 2*fold[1] - coord[1]
						new_coords = append(new_coords, coord)
					}
				} else {
					if coord[1] < fold[1] {
						coord[1] = 2*fold[1] - coord[1]
						coord[1] = max_y - coord[1]
						new_coords = append(new_coords, coord)
					} else if coord[1] > fold[1] {
						coord[1] = max_x - coord[1]
						new_coords = append(new_coords, coord)
					}
				}
			}
			if top >= bottom {
				max_y = top-1
			} else {
				max_y = bottom-1
			}
		}
		coords = new_coords
		break
	}

	grid := make([][]int, max_y+1)
	for i:=0;i<=max_y;i++ {
		grid[i] = make([]int, max_x+1)
	}

	for _, coord := range coords {
		grid[coord[1]][coord[0]] = 1
	}
	count := 0
	for i:=0;i<=max_y;i++ {
		for j:=0;j<=max_x;j++ {
			if grid[i][j] == 1 {
				count += 1
			}
		}
	}
	fmt.Println(count)
	
}