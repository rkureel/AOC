package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

func hex_to_binary(hex string) string {
	mapper := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	var bits strings.Builder
	for _, char := range hex {
		bits.WriteString(mapper[char])
	}
	return bits.String()
}

func sum(values []int64) int64 {
	var sum int64 = 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func prod(values []int64) int64 {
	var prod int64 = 1
	for _, val := range values {
		prod *= val
	}
	return prod
}

func min(values []int64) int64 {
	var sum int64 = values[0]
	for _, val := range values {
		if val < sum {
			sum = val
		}
	}
	return sum
}

func max(values []int64) int64 {
	var sum int64 = values[0]
	for _, val := range values {
		if val > sum {
			sum = val
		}
	}
	return sum
}

func great(values []int64) int64 {
	if values[0] > values[1] {
		return 1
	}
	return 0
}

func less(values []int64) int64 {
	if values[0] < values[1] {
		return 1
	}
	return 0
}

func equal(values []int64) int64 {
	if values[0] == values[1] {
		return 1
	}
	return 0
}

var packet string

func parse_packet(start int, end int) (int64, []int64, int) {
	version_number, _ := strconv.ParseInt(packet[start:start+3], 2, 64)
	type_id, _ := strconv.ParseInt(packet[start+3:start+6], 2, 64)
	if type_id == 4 {
		var literal_value strings.Builder
		curr := 6+start
		for {
			var last bool = false
			if packet[curr] == '0' {
				last = true
			}
			literal_value.WriteString(packet[curr+1:curr+5])
			curr += 5
			if last {
				break
			}			
		}
		literal_number, _ := strconv.ParseInt(literal_value.String(), 2, 64)
		return version_number, []int64{literal_number}, curr
	} else {
		literal_values := make([]int64, 0)
		var last_bit int64
		if packet[start+6] == '0' {
			len_packet, _ := strconv.ParseInt(packet[start+7:start+7+15], 2, 64)
			last_bit = len_packet+7+15+int64(start)
			packet_end := start+7+15
			for {
				packet_version, literal_value, next_bit := parse_packet(packet_end, int(last_bit))
				literal_values = append(literal_values, literal_value...)
				packet_end = next_bit
				version_number += packet_version
				if packet_end >= int(last_bit) {
					break
				}
			}
		} else {
			n_packets, _ := strconv.ParseInt(packet[start+7:start+7+11], 2, 64)
			packet_end := start+7+11
			for ;n_packets>0;n_packets--{
				packet_version, literal_value, next_bit := parse_packet(packet_end, end)
				literal_values = append(literal_values, literal_value...)
				packet_end = next_bit
				version_number += packet_version
			}
			last_bit = int64(packet_end)
		}

		var value int64
		
		switch type_id {
		case 0: 
			value = sum(literal_values)
		case 1:
			value = prod(literal_values)
		case 2:
			value = min(literal_values)
		case 3:
			value = max(literal_values)
		case 5:
			value = great(literal_values)
		case 6:
			value = less(literal_values)
		case 7:
			value = equal(literal_values)
		}

		return version_number, []int64{value}, int(last_bit)
	}
}

func main() {
	f, _ := os.Open("input.txt")
	hex_input := get_input(f)
	packet = hex_to_binary(hex_input)
	
	_, literal_values, _ := parse_packet(0, len(packet))
	fmt.Println(literal_values[0])
}