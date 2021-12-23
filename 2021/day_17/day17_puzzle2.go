package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(f *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	coords := strings.Split(line, ": ")[1]
	x_range := strings.Split(strings.Split(coords, ", ")[0], "=")[1]
	y_range := strings.Split(strings.Split(coords, ", ")[1], "=")[1]
	x_range_l, _ := strconv.Atoi(strings.Split(x_range, "..")[0])
	x_range_h, _ := strconv.Atoi(strings.Split(x_range, "..")[1])
	y_range_l, _ := strconv.Atoi(strings.Split(y_range, "..")[0])
	y_range_h, _ := strconv.Atoi(strings.Split(y_range, "..")[1])
	return []int{x_range_l, x_range_h}, []int{y_range_l, y_range_h}
}

func max(a int, b int) int {
	if a>=b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a<=b {
		return a
	}
	return b
}

func main() {
	f, _ := os.Open("input.txt")
	x_range, y_range := get_input(f)
	count := 0
	max_x := x_range[1]
	min_x := 0
	for {
		sum := min_x*(min_x+1)/2
		if sum >= x_range[0] {
			break
		}
		min_x++
	}
	max_y_steps := 1 + 2*(-1*y_range[0]-1)
	max_y := -1*y_range[0]-1
	for x:=min_x;x<=max_x;x++ {
		sum := 0
		max_step := 0
		min_steps := -1
		var num int
		for num=x;num>0;num-- {
			sum += num
			if sum>x_range[1] {
				break
			}
			if min_steps == -1 && sum>=x_range[0] {
				min_steps = max_step
			}
			max_step += 1
		}
		min_steps = max(min_steps, 0)
		
		if num == 0 {
			max_step = max_y_steps
		} 
		
		for y:=max_y;y>=y_range[0];y-- {
			var ok bool = false
			var y_coord int
			var x_coord int = 0
			var x_adder int = x
			for step := 0;step<=max_step;step++ {
				if step<=y {
					y_coord = y*(y+1)/2 - (y-step)*(y-step+1)/2
				} else {
					y_coord = y*(y+1)/2 - (step-y)*(step-y+1)/2
				}
				x_coord += x_adder
				if x_adder > 0 {
					x_adder -= 1
				}
				if x_coord>=x_range[0]&&x_coord<=x_range[1]&&y_coord>=y_range[0]&&y_coord<=y_range[1] {
					ok = true
				}
			}
			if ok {
				count += 1
			}
		}
	}
	fmt.Println(count)
}