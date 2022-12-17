use std::fs;

use aoc2022::day_06;

fn main() {
    let filepath = "inputs/day_06.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_06::part_1(&input);
    day_06::part_2(&input);
}
