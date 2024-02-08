use clap::Parser;
use std::fs;

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
const INPUT_FILE_NAME: &str = "input.txt";

fn get_input(day: u8, use_example: bool) -> String {
    let file_name: &str = if use_example {EXAMPLE_FILE_NAME} else {INPUT_FILE_NAME};
    let input_file_path: String = format!("inputs/day_{}/{}", day, file_name);
    // TODO: Change this to use a buffered reader
    let input: String = fs::read_to_string(input_file_path)
        .expect("Unable to read input file!");
    return input;
}

fn solve(day: u8, use_example: bool) { 
    let input: String = get_input(day, use_example); 
    match day {
        1 => solutions::day_1::solve(&input),
        _ => print!("Invalid day!"),
    }
}


fn main() {
    let args =  <Args as Parser>::parse();

    let day_number: u8 = args.day;
    let use_example: bool = args.example;
    if day_number < 1 || day_number > 25 {
        println!("Enter a valid day number!");
        return;
    }

    solve(day_number, use_example);
}
