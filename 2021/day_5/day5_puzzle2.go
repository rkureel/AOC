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

func abs(x int) int {
	if x>=0 {
		return x
	}
	return -x
}

func get_input(f *os.File) ([][]int, [][]int) {
	scanner := bufio.NewScanner(f)
	var start [][]int
	var end [][]int
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		point1_text := strings.Split(points[0], ",")
		point2_text := strings.Split(points[1], ",")
		
		point1 := make([]int, 2)
		point2 := make([]int, 2)
		point1[0], _ = strconv.Atoi(point1_text[0])
		point1[1], _ = strconv.Atoi(point1_text[1])
		point2[0], _ = strconv.Atoi(point2_text[0])
		point2[1], _ = strconv.Atoi(point2_text[1])
		if point1[0]<point2[0]||point1[0]==point2[0]&&point1[1]<point2[1] {
			start = append(start, point1)
			end = append(end, point2)
		} else {
			start = append(start, point2)
			end = append(end, point1)
		}
	}
	return start, end
}

func main() {
	f, _ := os.Open("input.txt")
	start, end := get_input(f)
	var (
		i int
		j int
		n int
	)
	n = len(start)
	var row_size int = 0
	var col_size int = 0
	for i=0;i<n;i++ {
		row_size = max(row_size, max(start[i][0], end[i][0]))
		col_size = max(col_size, max(start[i][1], end[i][1]))
	}
	rows := make([][]int, col_size+2)
	cols := make([][]int, row_size+2)
	diag1 := make([][]int, row_size+col_size+2)
	diag2 := make([][]int, row_size+col_size+2)
	for i=0;i<col_size+2;i++ {
		rows[i] = make([]int, row_size+2)
	}
	for i=0;i<row_size+2;i++ {
		cols[i] = make([]int, col_size+2)
	}
	for i=0;i<row_size+col_size+2;i++ {
		diag1[i] = make([]int, row_size+col_size+2)
		diag2[i] = make([]int, row_size+col_size+2)
	}
	for i=0;i<n;i++ {
		if start[i][0] == end[i][0] {
			cols[start[i][0]][start[i][1]] += 1
			cols[end[i][0]][end[i][1]+1] -= 1	
		} else if start[i][1] == end[i][1] {
			rows[start[i][1]][start[i][0]] += 1
			rows[end[i][1]][end[i][0]+1] -= 1
		} else if (end[i][1] - start[i][1])/(end[i][0] - start[i][0]) == -1 {
			start_index_y := start[i][0] + col_size+1-start[i][1]
			start_index_x := start[i][0] + start[i][1]
			end_index_y := end[i][0] + col_size+1-end[i][1]
			end_index_x := end[i][0] + end[i][1]
			diag1[start_index_x][start_index_y] += 1
			diag1[end_index_x][end_index_y+1] -= 1
		} else if (end[i][1] - start[i][1])/(end[i][0] - start[i][0]) == 1 {
			start_index_x := start[i][0] + col_size+1-start[i][1]
			start_index_y := start[i][0] + start[i][1]
			end_index_x := end[i][0] + col_size+1-end[i][1]
			end_index_y := end[i][0] + end[i][1]
			diag2[start_index_x][start_index_y] += 1
			diag2[end_index_x][end_index_y+1] -= 1
		}	
	}

	for i=0;i<=row_size;i++ {
		for j=1;j<=col_size;j++ {
			cols[i][j] += cols[i][j-1]
		}
	}
	for i=0;i<=col_size;i++ {
		for j=1;j<=row_size;j++ {
			rows[i][j] += rows[i][j-1]
		}
	}
	for i=0;i<row_size+col_size+2;i++ {
		for j=1;j<row_size+col_size;j++ {
			diag1[i][j] += diag1[i][j-1]
			diag2[i][j] += diag2[i][j-1]
		}
	}
	
	sum := 0
	for i=0;i<=row_size;i++ {
		for j=0;j<=col_size;j++ {
			a := i + col_size+1-j
			b := i+j
			diag1_count := diag1[b][a]
			diag2_count := diag2[a][b]
			if rows[j][i] + cols[i][j] + diag1_count + diag2_count> 1 {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}