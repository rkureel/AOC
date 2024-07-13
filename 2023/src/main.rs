use clap::Parser;

#[derive(Parser, Debug)]
struct Args {
    /// AOC problem day number [1-25]
    #[arg(short, long)]
    day: u8,
}

fn main() {
    let args =  <Args as Parser>::parse();

    let day_number: u8 = args.day;
    if day_number < 1 || day_number > 25 {
        println!("Enter a valid day number!");
        return;
    }

    aoc2023::solve(day_number);
}

