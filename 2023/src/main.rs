use clap::Parser;

pub mod solutions;
pub mod utils;

#[derive(Parser, Debug)]
struct Args {
    /// AOC problem day number [1-25]
    #[arg(short, long)]
    day: u8,

    #[arg(short, long)]
    example: bool,
}

fn solve(day: u8, use_example: bool) { 
    let input: Vec<String> = utils::get_input(day, use_example); 
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
