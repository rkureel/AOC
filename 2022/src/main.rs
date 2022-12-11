use std::fs;

use aoc2022::day_03;

fn main() {
    let filepath = "inputs/day_03.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_03::part_1(&input);
    day_03::part_2(&input);
}
