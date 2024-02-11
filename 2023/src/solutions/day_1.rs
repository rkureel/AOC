pub fn solve(input: &Vec<String>) {
    println!("Solving for day 1.");
    part_1(input);
}

fn part_1(input: &Vec<String>) {
    let mut sum: u32 = 0;
    for line in input {
        let digits: Vec<u32> = line.chars()
            .filter_map(|c| c.to_digit(10))
            .collect();

        if let Some(first_digit) = digits.first() {    
            if let Some(last_digit) = digits.last() {
                sum += first_digit*10 + last_digit;
            }
        }
    }
    println!("Solution for part 1: {}", sum);
}

