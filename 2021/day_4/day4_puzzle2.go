package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(f *os.File) ([]int, [][5][5]int) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	var numbers []int
	for i:=0;i<len(line);i++ {
		number, _ := strconv.Atoi(line[i])
		numbers = append(numbers, number)
	}
	
	var boards [][5][5]int
	var board [5][5]int
	var i int
	var j int
	var k int
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		line = strings.Split(scanner.Text(), " ")
		k = 0
		for j=0;j<5;j++ {
			for line[k] == "" && k<len(line) {
				k++
			}
			board[i][j], _ = strconv.Atoi(line[k])
			k++
		}
		i++
		if(i==5) {
			i = 0
			boards = append(boards, board)
		}
	}
	return numbers, boards
}

func main() {
	f, _ := os.Open("input.txt")
	numbers, boards := get_input(f)
	var (
		i int
		j int
		k int
		l int
		n int
	)
	n = len(boards)
	marked := make([][5][5]bool, n)
	win_marker := make([]bool, n)
	
	wins := 0
	for i=0;i<len(numbers);i++ {
		for j=0;j<n;j++ {
			
			if win_marker[j] == true {
				continue
			}
			for k=0;k<5;k++ {
				for l=0;l<5;l++ {
					if boards[j][k][l] == numbers[i] {
						marked[j][k][l] = true
					}
				}
			}
			for k=0;k<5;k++ {
				row := true
				col := true
				for l=0;l<5;l++ { 
					if marked[j][k][l] != true {
						row = false
					}
					if marked[j][l][k] != true {
						col = false
					}
				}
				if row==true || col==true{
					wins += 1
					win_marker[j] = true
					break
				}
			}
			if wins == len(boards) {
				break
			}
		}
		if wins == len(boards) {
			break
		}
	}

	unmarked_sum := 0
	for k=0;k<5;k++ {
		for l=0;l<5;l++ {
			if marked[j][k][l] != true {
				unmarked_sum += boards[j][k][l]
			}
		}
	}
	fmt.Println(unmarked_sum*numbers[i])
}