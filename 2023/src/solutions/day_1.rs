pub fn solve(input: &str) {
    println!("Solving for day 1.");
    let mut sum: u32 = 0;
    for line in input.lines() {
        let mut first_digit: u32 = 0;
        let mut second_digit: u32 = 0;
        for c in line.chars() {
            if c.is_ascii_digit() {
                first_digit = c.to_digit(10)
                    .unwrap();
                break;
            }
        }
        for c in line.chars().rev() {
            if c.is_ascii_digit() {
                second_digit = c.to_digit(10)
                    .unwrap();
                break;
                
            }
        }
        let caliberation_value: u32 = first_digit*10 + second_digit;
        sum += caliberation_value;
    }
    dbg!(sum);
}
