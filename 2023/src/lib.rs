use std::fs::read_to_string;
mod solutions;

const INPUT_FILE_NAME: &str = "input.txt";

pub fn solve(day: u8) { 
    let input: Vec<String> = get_input(day); 
    match day {
        1 => solutions::day_1::solve(&input),
        2 => solutions::day_2::solve(&input),
        3 => solutions::day_3::solve(&input),
        _ => print!("Invalid day!"),
    }
}

pub fn generate_test_input(input: &str) -> Vec<String> {
    String::from(input)
        .lines()
        .map(String::from)
        .collect()
}

fn get_input(day: u8) -> Vec<String> {
    let file_name: &str =  INPUT_FILE_NAME;
    let input_file_path: String = format!("inputs/day_{}/{}", day, file_name);
    // TODO: Change this to use a buffered reader
    read_lines(&input_file_path)
}

fn read_lines(file_name: &str) -> Vec<String> {
    read_to_string(file_name)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

