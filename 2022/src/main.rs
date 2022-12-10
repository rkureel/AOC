use std::fs;

use day_01::{day_01_part_1, day_02_part_2};

pub mod day_01;

fn main() {
    let filepath = "inputs/day_01.txt";
    let input = fs::read_to_string(filepath)
    .expect("Unable to read input file.");
    
    day_01_part_1(&input);
    day_02_part_2(&input);
}
