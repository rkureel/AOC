use std::fs;

use aoc2022::day_09;

fn main() {
    let filepath = "inputs/day_09.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_09::part_1(&input);
    day_09::part_2(&input);
}
