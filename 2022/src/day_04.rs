pub fn part_1(input: &String) {
    let mut count: u32 = 0;
    for line in input.lines() {
        if line.len() == 0 {
            break;
        }
        let ranges: Vec<&str> = line.trim().split(",").collect();
        let left_numbers: Vec<&str> = ranges[0].split("-").collect();
        let right_numbers: Vec<&str> = ranges[1].split("-").collect();
        let left_start: u32 = left_numbers[0].parse().expect("Unable to parse");
        let left_end: u32 = left_numbers[1].parse().expect("Unable to parse");
        let right_start: u32 = right_numbers[0].parse().expect("Unable to parse");
        let right_end: u32 = right_numbers[1].parse().expect("Unable to parse");
        if left_start<=right_start && left_end >= right_end {
            count += 1;
        }
        else if right_start <= left_start && right_end >= left_end {
            count += 1;
        }
    }
    println!("Part 1: {}", count);
}

pub fn part_2(input: &String) {
    let mut count: u32 = 0;
    for line in input.lines() {
        if line.len() == 0 {
            break;
        }
        let ranges: Vec<&str> = line.trim().split(",").collect();
        let left_numbers: Vec<&str> = ranges[0].split("-").collect();
        let right_numbers: Vec<&str> = ranges[1].split("-").collect();
        let left_start: u32 = left_numbers[0].parse().expect("Unable to parse");
        let left_end: u32 = left_numbers[1].parse().expect("Unable to parse");
        let right_start: u32 = right_numbers[0].parse().expect("Unable to parse");
        let right_end: u32 = right_numbers[1].parse().expect("Unable to parse");
        if right_start >= left_start && right_start <= left_end {
            count += 1;
        }
        else if left_start >= right_start && left_start <= right_end {
            count += 1;
        }
    }
    println!("Part 2: {}", count);
}
