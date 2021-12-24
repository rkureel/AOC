package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reduce(line string) string {
	for {
		explode := false
		split := false
		var i int
		var j int
		var split_number int
		var split_right_index int
		lvl := 0
		for i=0;i<len(line);i++ {
			if line[i] == '[' {
				lvl ++
				if lvl == 5 {
					explode = true
					break
				}
			} else if line[i] == ']' {
				lvl --
			}
		}
		
		if explode {
			for j=i+1;line[j]!=']';j++ {}
			j+=1
			explosion_cell := line[i:j-1]
			line = line[:i] + "0" + line[j:]
			left_no, _ := strconv.Atoi(strings.Split(explosion_cell[1:], ",")[0])
			right_no, _ := strconv.Atoi(strings.Split(explosion_cell, ",")[1])
			j = i+1
			for ;j<len(line);j++ {
				if line[j]>='0'&&line[j]<='9' {
					var k int = j+1
					for {
						if k==len(line) || !(line[k]>='0'&&line[k]<='9') {
							break
						}
						k++
					}
					right_side_no, _ := strconv.Atoi(line[j:k])
					right_side_no += right_no
					right_string := strconv.Itoa(right_side_no)
					line = line[:j] + right_string + line[k:]
					break
				}
			}
			j = i-1
			for ;j>=0;j-- {
				if line[j]>='0'&&line[j]<='9' {
					var k int = j-1
					for {
						if k<0||!(line[k]>='0'&&line[k]<='9') {
							break
						}
						k--
					}
					j+=1
					k+=1
					left_side_no, _ := strconv.Atoi(line[k:j])
					left_side_no += left_no
					left_string := strconv.Itoa(left_side_no)
					line = line[:k] + left_string + line[j:]
					break
				}
			}
			continue
		}

		for i=0;i<len(line);i++ {
			if line[i]>='0'&&line[i]<='9' {
				j = i+1
				for {
					if !(line[j]>='0'&&line[j]<='9') {
						break
					}
					j++
				}
				split_right_index = j
				split_number, _ = strconv.Atoi(line[i:j])
				if split_number >= 10 {
					split = true
					break
				}
			}
		}
		
		if split {
			left_split := split_number/2
			right_split := left_split
			if split_number%2==1 {
				right_split += 1
			}
			left_split_string := strconv.Itoa(left_split)
			right_split_string := strconv.Itoa(right_split)
			line = line[:i] + "[" + left_split_string + "," + right_split_string + "]" + line[split_right_index:]
		}
		if !(explode||split) {
			break
		}
	}
	return line
}

func get_input(f *os.File) (string) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()
	input = reduce(input)
	for scanner.Scan() {
		line := scanner.Text()
		input = "[" + input + "," + line + "]"
		input = reduce(input)
	}
	return input
}

var input string

func get_magnitude(start int, end int) int {
	if input[start]>='0'&&input[start]<='9' {
		number, _ := strconv.Atoi(input[start:end+1])
		return number
	}
	level := 0
	location := -1
	for i:=start+1;i<end;i++ {
		if input[i] == '[' {
			level += 1
		} else if input[i] == ']' {
			level -= 1
		} else if input[i] == ',' && level==0{
			location = i
		}
	}
	return 3*get_magnitude(start+1, location-1) + 2*get_magnitude(location+1, end-1)
} 

func main() {
	f, _ := os.Open("input.txt")
	input = get_input(f)
	fmt.Println(get_magnitude(0, len(input)-1))
}