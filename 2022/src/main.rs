use std::fs;

use aoc2022::day_07;

fn main() {
    let filepath = "inputs/day_07.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_07::part_1(&input);
    day_07::part_2(&input);
}
