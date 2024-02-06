use clap::Parser;
use std::{path::{Path, PathBuf}, fs};

pub mod solutions;

#[derive(Parser, Debug)]
struct Args {
    /// AOC problem day number [1-25]
    #[arg(short, long)]
    day: u8,

    #[arg(short, long)]
    example: bool,
}

const EXAMPLE_FILE_NAME: &str = "example.txt";
const INPUTS_DIR: &str = "inputs";
const INPUT_FILE_NAME: &str = "input.txt";

fn solve(day: u8, use_example: bool) {
    println!("Solving for day {day}. Using example: {use_example}.");
    
    let input_file_path: PathBuf = Path::new(INPUTS_DIR)
        .join(format!("day_{}", day))
        .join(if use_example {EXAMPLE_FILE_NAME} else {INPUT_FILE_NAME});

    let input_file_name_string: &str = input_file_path
        .to_str()
        .unwrap();

    let input: String = fs::read_to_string(input_file_name_string)
        .expect("Unable to read input file!");

    match day {
        1 => solutions::day_1::solve(&input),
        _ => print!("Invalid day!"),
    }
}


fn main() {
    println!("Hello, world!");
    let args =  <Args as Parser>::parse();

    let day_number: u8 = args.day;
    let use_example: bool = args.example;
    if day_number < 1 || day_number > 25 {
        println!("Enter a valid day number!");
        return;
    }

    solve(day_number, use_example);
}
