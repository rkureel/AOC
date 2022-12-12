use std::fs;

use aoc2022::day_05;

fn main() {
    let filepath = "inputs/day_05.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_05::part_1(&input);
    day_05::part_2(&input);
}
