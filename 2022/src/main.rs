use std::fs;

use aoc2022::day_01;

fn main() {
    let filepath = "inputs/day_01.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_01::part_1(&input);
    day_01::part_2(&input);
}
