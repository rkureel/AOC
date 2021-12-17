package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func get_input(f *os.File) ([][]int, [][]bool) {
	scanner := bufio.NewScanner(f)
	var numbers [][]int
	var visited [][]bool
	for scanner.Scan() {
		var row []int
		line := strings.Split(scanner.Text(), "")
		for i:=0;i<len(line);i++ {
			number, _ := strconv.Atoi(line[i])
			row = append(row, number)
		}
		numbers = append(numbers, row)
		visited = append(visited, make([]bool, len(row)))
	}
	return numbers, visited
}

var numbers [][]int
var visited [][]bool

func dfs_sum(i int, j int, pi int, pj int) int {
	
	if numbers[i][j] == 9 || visited[i][j] == true {
		return 0
	}
	if pi!=-1 && pj!=-1 {
		low_x := pi
		low_y := pj
		if i>0 && numbers[i-1][j]<numbers[low_x][low_y] {
			low_x = i-1
			low_y = j
		}
		if i<len(numbers)-1&&numbers[i+1][j]<numbers[low_x][low_y] {
			low_x = i+1
			low_y = j
		}
		if j>0 && numbers[i][j-1]<numbers[low_x][low_y] {
			low_x = i
			low_y = j-1
		}
		if j<len(numbers[0])-1 && numbers[i][j+1]<numbers[low_x][low_y] {
			low_x = i
			low_y = j+1
		}
		if low_x != pi || low_y != pj {
			return 0
		}
	}
	sum := 1
	visited[i][j] = true
	if i>0 && numbers[i-1][j]>numbers[i][j] {
		if i-1!=pi || j!=pj {
			sum += dfs_sum(i-1, j, i, j)
		}
	}
	if i<len(numbers)-1&&numbers[i+1][j]>numbers[i][j] {
		if i+1!=pi || j!=pj {
			sum += dfs_sum(i+1, j, i, j)
		}
	}
	if j>0 && numbers[i][j-1]>numbers[i][j] {
		if i!=pi || j-1!=pj {	
			sum += dfs_sum(i, j-1, i, j)
		}
	}
	if j<len(numbers[0])-1 && numbers[i][j+1]>numbers[i][j] {
		if i!=pi || j+1!=pj {	
			sum += dfs_sum(i, j+1, i, j)
		}
	}
	return sum
}

func main() {
	f, _ := os.Open("input.txt")
	numbers, visited = get_input(f)
	for i:=0;i<len(numbers);i++ {
		for j:=0;j<len(numbers[0]);j++ {
			visited[i][j] = false
		}
	}

	var sizes []int
	for i, row := range numbers {
		for j, val := range row {
			ok := true
			if i>0 && val>=numbers[i-1][j] {
				ok = false
			}
			if i<len(numbers)-1 && val>=numbers[i+1][j] {
				ok = false
			}
			if j>0 && val>=numbers[i][j-1] {
				ok = false
			}	
			if j<len(row)-1 && val>=numbers[i][j+1] {
				ok = false
			}
			if ok {
				sizes = append(sizes, dfs_sum(i, j, -1, -1))
			}
		}
	}

	sort.Ints(sizes)
	n := len(sizes)
	fmt.Println(sizes[n-1]*sizes[n-2]*sizes[n-3])
}