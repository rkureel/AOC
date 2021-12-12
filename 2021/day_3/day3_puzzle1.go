package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	bits_n := 12
	ones := make([]int, bits_n)
	zeros := make([]int, bits_n)
	for scanner.Scan() {
		number := scanner.Text()
		for i:=0;i<len(number);i++ {
			if(string(number[i])=="0") {
				zeros[i] += 1
			} else {
				ones[i] +=1
			}
		}
	}
	var gamma strings.Builder
	var epsilon strings.Builder
	for i:=0;i<bits_n;i++ {
		if ones[i] >= zeros[i] {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}
	gamma_rate, _ := strconv.ParseInt(gamma.String(), 2, 64)
	epsilon_rate, _ := strconv.ParseInt(epsilon.String(), 2, 64)
	fmt.Print(gamma_rate*epsilon_rate)
}