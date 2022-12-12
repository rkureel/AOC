use std::fs;

use aoc2022::day_04;

fn main() {
    let filepath = "inputs/day_04.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_04::part_1(&input);
    day_04::part_2(&input);
}
