package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_input(f *os.File) map[string][]string {
	scanner := bufio.NewScanner(f)
	input := make(map[string][]string)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		input[line[0]] = append(input[line[0]], line[1])
		input[line[1]] = append(input[line[1]], line[0])
	}
	return input
}

var adj map[string][]string
var path_counts map[string]int

func dfs(node string, parent string, visited bool) int {
	if node == "end" {
		return 1
	}
	path_counts[node]+=1
	sum := 0
	
	for _, next := range adj[node] {
		if next == "start" {
			continue
		} else if next[0]>='A'&&next[0]<='Z' { 
			sum += dfs(next, node, visited)
		} else if next[0]>='a'&&next[0]<='z'&&path_counts[next]<=1 {
			if visited {
				if path_counts[next]==0 {
					sum += dfs(next, node, true)
				}
			} else {
				if path_counts[next] == 0 {
					sum += dfs(next, node, false)
				} else if path_counts[next] == 1 {
					sum += dfs(next, node, true)
				}
			}
		}	
	}
	path_counts[node] -= 1
	return sum
}

func main() {
	f, _ := os.Open("input.txt")
	adj = get_input(f)
	path_counts = make(map[string]int)
	for key := range adj {
		path_counts[key] = 0
	}
	ans := dfs("start", "nil", false)
	fmt.Println(ans)
}