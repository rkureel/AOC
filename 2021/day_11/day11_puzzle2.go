package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	var numbers [][]int
	for scanner.Scan() {
		var row []int
		line := strings.Split(scanner.Text(), "")
		for i:=0;i<len(line);i++ {
			number, _ := strconv.Atoi(line[i])
			row = append(row, number)
		}
		numbers = append(numbers, row)
	}
	return numbers
}

func main() {
	f, _ := os.Open("input.txt")
	numbers := get_input(f)
	for step:=0;;step++ {
		var x []int
		var y []int
		var visited [10][10]bool
		for i:=0;i<10;i++ {
			for j:=0;j<10;j++ {
				numbers[i][j] += 1
				visited[i][j] = false
			}
		}
		for i:=0;i<10;i++ {
			for j:=0;j<10;j++ {
				if numbers[i][j] > 9 && !visited[i][j] {
					visited[i][j] = true
					x = append(x, i)
					y = append(y, j)
				}
			}
		}
		
		for len(x)>0 {
			i := x[0]
			j := y[0]
			x = x[1:]
			y = y[1:]
			
			if i>0 {
				numbers[i-1][j] += 1
				if numbers[i-1][j] > 9 && !visited[i-1][j]{
					x = append(x, i-1)
					y = append(y, j)
					visited[i-1][j] = true
				}
				
				if j>0 {
					numbers[i-1][j-1] += 1
					if numbers[i-1][j-1] > 9 && !visited[i-1][j-1]{
						x = append(x, i-1)
						y = append(y, j-1)
						visited[i-1][j-1] = true
					}
				}
				if j<9 {
					numbers[i-1][j+1] += 1
					if numbers[i-1][j+1] > 9 && !visited[i-1][j+1]{
						x = append(x, i-1)
						y = append(y, j+1)
						visited[i-1][j+1] = true
					}
				}
			}
			if i<9 {
				numbers[i+1][j] += 1
				if numbers[i+1][j] > 9 && !visited[i+1][j]{
					x = append(x, i+1)
					y = append(y, j)
					visited[i+1][j] = true
				}
				
				if j>0 {
					numbers[i+1][j-1] += 1
					if numbers[i+1][j-1] > 9 && !visited[i+1][j-1]{
						x = append(x, i+1)
						y = append(y, j-1)
						visited[i+1][j-1] = true
					}
				}
				if j<9 {
					numbers[i+1][j+1] += 1
					if numbers[i+1][j+1] > 9 && !visited[i+1][j+1]{
						x = append(x, i+1)
						y = append(y, j+1)
						visited[i+1][j+1] = true
					}
				}
			}
			if j>0 {
				numbers[i][j-1] += 1
				if numbers[i][j-1] > 9 && !visited[i][j-1]{
					x = append(x, i)
					y = append(y, j-1)
					visited[i][j-1] = true
				}
			}
			if j<9 {
				numbers[i][j+1] += 1
				if numbers[i][j+1] > 9 && !visited[i][j+1]{
					x = append(x, i)
					y = append(y, j+1)
					visited[i][j+1] = true
				}
			}
		}
		count := 0
		for i:=0;i<10;i++ {
			for j:=0;j<10;j++ {
				if numbers[i][j] > 9 {
					numbers[i][j] = 0
					count += 1
				}
			}
		}
		if count == 100 {
			fmt.Println(step+1)
			break
		}
	}
}