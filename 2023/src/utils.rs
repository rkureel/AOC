use std::fs::read_to_string;

const EXAMPLE_FILE_NAME: &str = "example.txt";
const INPUT_FILE_NAME: &str = "input.txt";


pub fn get_input(day: u8, use_example: bool) -> Vec<String> {
    let file_name: &str = if use_example {EXAMPLE_FILE_NAME} else {INPUT_FILE_NAME};
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
