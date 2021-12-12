package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func get_numbers(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	numbers := make([]string, 0, 1000)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	sort.Strings(numbers)
	return numbers
}

func get_bit_counts(start int, end int, bit_pos int, ones [][]int, zeros [][]int) (zero_counts int, one_counts int) {
	if start == 0 {
		return zeros[end][bit_pos], ones[end][bit_pos]
	} else {
		return zeros[end][bit_pos]-zeros[start-1][bit_pos], ones[end][bit_pos]-ones[start-1][bit_pos]
	}
}

func main() {
	f, _ := os.Open("input.txt")
	numbers := get_numbers(f)
	const bits_n int = 12
	var n int = len(numbers)
	ones := make([][]int, n)
	zeros := make([][]int, n)
	var (
		i int
		j int
		start int
		end int
	)

	for i=0;i<n;i++ {
		ones[i] = make([]int, bits_n)
		zeros[i] = make([]int, bits_n)

		if i==0 {
			for j=0;j<bits_n;j++ {
				if string(numbers[i][j]) == "0" {
					ones[i][j] = 0
					zeros[i][j] = 1
				} else {
					ones[i][j] = 1
					zeros[i][j] = 0
				}
			}
		} else {
			for j=0;j<bits_n;j++ {
				if string(numbers[i][j]) == "0" {
					ones[i][j] = ones[i-1][j]
					zeros[i][j] = zeros[i-1][j] + 1
				} else {
					ones[i][j] = ones[i-1][j] + 1
					zeros[i][j] = zeros[i-1][j]
				}
			}
		}
	}

	start = 0
	end = n-1
	var o2_number string
	for i=0;i<bits_n;i++ {
		zero_counts, one_counts := get_bit_counts(start, end, i, ones, zeros)
		for j=start;j<=end;j++ {
			if string(numbers[j][i]) == "1" {
				break
			}
		}
		if zero_counts > one_counts {
			end = j-1
		} else if zero_counts <= one_counts {
			start = j
		}
		if(end==start) {
			o2_number = numbers[start]
			break
		}
	}

	start = 0
	end = n-1
	var co2_number string
	for i=0;i<bits_n;i++ {
		zero_counts, one_counts := get_bit_counts(start, end, i, ones, zeros)
		for j=start;j<=end;j++ {
			if string(numbers[j][i]) == "1" {
				break
			}
		}
		if zero_counts > one_counts {
			start = j
		} else if zero_counts <= one_counts {
			end = j-1
		}
		if(end==start) {
			co2_number = numbers[start]
			break
		}
	}


	oxygen_rating, _ := strconv.ParseInt(o2_number, 2, 64)
	co2_rating, _ := strconv.ParseInt(co2_number, 2, 64)
	fmt.Println(oxygen_rating*co2_rating)
}