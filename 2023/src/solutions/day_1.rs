use strum::IntoEnumIterator;
use strum_macros::EnumIter;


pub fn solve(input: &Vec<String>) {
    println!("Solving for day 1.");
    println!("Solution for part 1: {}", part_1(input));
    println!("Solution for part 2: {}", part_2(input));
}

fn part_1(input: &Vec<String>) -> u32 {
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
    sum
}


fn part_2(input: &Vec<String>) -> u32 {
    let sum: u32 = input.iter().map(|line| {
            let (first_digit, last_digit) = Digit::find_first_and_last_digits(line);
            10*first_digit as u32 + last_digit as u32
        })
        .sum();
    sum
}

#[derive(Clone, Copy, Debug, EnumIter)]
enum Digit {
    One = 1,
    Two = 2,
    Three = 3,
    Four = 4,
    Five = 5,
    Six = 6,
    Seven = 7,
    Eight = 8,
    Nine = 9,
}

impl Digit {

    fn find_first_and_last_digits(input: &str) -> (Self, Self) {
        let mut first_digit_index: usize = usize::MAX;
        let mut last_digit_index: usize = usize::MIN;

        let mut first_digit = Self::One;
        let mut last_digit = Self::One;

        for digit in Self::iter() {
            let first_index = match (input.find(&digit.as_spelling()), input.find(digit.as_character())) {
                (Some(s), Some(c)) => s.min(c),
                (Some(s), None) => s,
                (None, Some(c)) => c,
                (None, None) => continue,
            };

            let last_index = match (input.rfind(&digit.as_spelling()), input.rfind(digit.as_character())) {
                (Some(s), Some(c)) => s.max(c),
                (Some(s), None) => s,
                (None, Some(c)) => c,
                (None, None) => continue,
            };

            if first_index < first_digit_index {
                first_digit_index = first_index;
                first_digit = digit;
            }
            if last_index >= last_digit_index {
                last_digit_index = last_index;
                last_digit = digit;
            }
        }

        (first_digit, last_digit)
    }

    fn as_spelling(&self) -> String {
        match self {
            Self::One => String::from("one"),
            Self::Two => String::from("two"),
            Self::Three => String::from("three"),
            Self::Four => String::from("four"),
            Self::Five => String::from("five"),
            Self::Six => String::from("six"),
            Self::Seven => String::from("seven"),
            Self::Eight => String::from("eight"),
            Self::Nine => String::from("nine"),
        }
    }

    fn as_character(&self) -> char {
        match self {
            Self::One => '1',
            Self::Two => '2',
            Self::Three => '3',
            Self::Four => '4',
            Self::Five => '5',
            Self::Six => '6',
            Self::Seven => '7',
            Self::Eight => '8',
            Self::Nine => '9',
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn day_1_part_1() {
        let input = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";
        assert_eq!(super::part_1(&crate::generate_test_input(&input)), 142); 
    }

    #[test]
    fn day_1_part_2() {
        let input = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";
        assert_eq!(super::part_2(&crate::generate_test_input(&input)), 281); 
    }
}
