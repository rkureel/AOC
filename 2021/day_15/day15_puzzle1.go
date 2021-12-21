package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func min(a int, b int) int {
	if a<=b {
		return a
	}
	return b
}

func swap(a *node, b *node) {
	var temp node = *a
	*a = *b
	*b = temp
}

type node struct {
	i int
	j int
	value int
}

type heap struct {
	arr []node
	size int
}

func newheap(cap int) heap {
	h := heap{
		arr: make([]node, cap),
		size: 0,
	}
	return h
}

func (h *heap) heapify(i int) {
	n := h.size
	smallest := i
	left := 2*i+1
	right := 2*i+2
	if left <= n-1 && h.arr[left].value < h.arr[smallest].value {
		smallest = left
	}
	if right <= n-1 && h.arr[right].value < h.arr[smallest].value {
		smallest = right
	}
	if smallest != i {
		swap(&h.arr[smallest], &h.arr[i])
		h.heapify(smallest)
	}
}

func (h *heap) insert(x node) {
	h.arr = append(h.arr, x)
	h.arr[h.size] = x
	h.size+=1
	i := h.size - 1
	for {
		if i==0 {
			break
		}
		parent := (i-1)/2
		if h.arr[parent].value>h.arr[i].value {
			swap(&h.arr[i], &h.arr[parent])
			i = parent
		} else {
			break
		}
	}
}

func (h *heap) extract() node {
	element := h.arr[0]
	swap(&h.arr[0], &h.arr[h.size-1])
	h.size -= 1
	h.heapify(0)
	return element
}

func get_input(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		line_numbers := make([]int, 0)
		for _, char := range line {
			number, _ := strconv.Atoi(char)
			line_numbers = append(line_numbers, number)
		}
		grid = append(grid, line_numbers)
	}
	return grid
}

func main() {
	f, _ := os.Open("input.txt")
	grid := get_input(f)
	m := len(grid)
	n := len(grid[0])
	h := newheap(n*m)
	visited := make([][]bool, m)
	values := make([][]int, m)
	for i:=0;i<m;i++ {
		values[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j:=0;j<n;j++ {
			var value int
			if i==0 && j==0 {
				value = 0
			} else {
				value = math.MaxInt32
			}
			values[i][j] = value
			visited[i][j] = false
			node := node{
				i: i,
				j: j,
				value: value,
			}
			h.insert(node)
		}
	}

	var curr node
	for {
		curr = h.extract()
		var i int = curr.i
		var j int = curr.j
		if i == m-1 && j== n-1 {
			break
		}
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		
		var x int
		var y int
		if i>0 {
			x = i-1
			y = j
			if !visited[x][y] {
				values[x][y] = min(values[x][y], values[i][j] + grid[x][y])
				
				new_node := node {
					i: x,
					j: y,
					value: values[x][y],
				}
				h.insert(new_node)
			}
		}
		if i<m-1 {
			x = i+1
			y = j
			if !visited[x][y] {
				values[x][y] = min(values[x][y], values[i][j] + grid[x][y])
				
				new_node := node {
					i: x,
					j: y,
					value: values[x][y],
				}
				h.insert(new_node)
			}
		}
		if j<m-1 {
			x = i
			y = j+1
			if !visited[x][y] {
				values[x][y] = min(values[x][y], values[i][j] + grid[x][y])
				
				new_node := node {
					i: x,
					j: y,
					value: values[x][y],
				}
				h.insert(new_node)
			}
		}
		if j>0 {
			x = i
			y = j-1
			if !visited[x][y] {
				values[x][y] = min(values[x][y], values[i][j] + grid[x][y])
				
				new_node := node {
					i: x,
					j: y,
					value: values[x][y],
				}
				h.insert(new_node)
			}
		}
	}
	
	fmt.Println(values[m-1][n-1])
}