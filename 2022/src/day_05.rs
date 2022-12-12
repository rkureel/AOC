use std::collections::VecDeque;

pub fn part_1(input: &String) {
    let mut reading_stack: bool = true;
    let stack_size: u32 = 9;
    let mut stacks: Vec<VecDeque<char>> = Vec::new();
    for _i in 0..stack_size {
        stacks.push(VecDeque::new());
    }
    for line in input.lines() {
        if line.len() == 0 {
            continue;
        }
        let line_vec: Vec<char> = line.chars().collect();
        if reading_stack {
            if line_vec[1] == '1' {
                reading_stack = false;
                continue;
            } else {
                for i in 0..stack_size as usize {
                    if line_vec[4*i+1]>='A' && line_vec[4*i+1]<='Z' {
                        stacks[i].push_front(line_vec[4*i+1]);
                    }
                }           
            }
        }
        else {
            let words: Vec<&str> = line.split(" ").collect();
            let quantity: u32 = words[1].parse().expect("Unable to parse");
            let from: u32 = words[3].parse::<u32>().expect("Unable to parse")-1;
            let to: u32 = words[5].parse::<u32>().expect("Unable to parse")-1;
            for _i in 0..quantity {
                let val: char = stacks[from as usize].pop_back().unwrap();
                stacks[to as usize].push_back(val);    
            }
        }
    }
    let mut message: String = String::from("");
    for i in 0..stack_size as usize {
        let val: char = stacks[i].pop_back().unwrap();
        message.push(val);
    }
    println!("Part 1: {}", message);
}

pub fn part_2(input: &String) {
    let mut reading_stack: bool = true;
    let stack_size: u32 = 9;
    let mut stacks: Vec<VecDeque<char>> = Vec::new();
    for _i in 0..stack_size {
        stacks.push(VecDeque::new());
    }
    for line in input.lines() {
        if line.len() == 0 {
            continue;
        }
        let line_vec: Vec<char> = line.chars().collect();
        if reading_stack {
            if line_vec[1] == '1' {
                reading_stack = false;
                continue;
            } else {
                for i in 0..stack_size as usize {
                    if line_vec[4*i+1]>='A' && line_vec[4*i+1]<='Z' {
                        stacks[i].push_front(line_vec[4*i+1]);
                    }
                }           
            }
        }
        else {
            let words: Vec<&str> = line.split(" ").collect();
            let quantity: u32 = words[1].parse().expect("Unable to parse");
            let from: u32 = words[3].parse::<u32>().expect("Unable to parse")-1;
            let to: u32 = words[5].parse::<u32>().expect("Unable to parse")-1;
            let mut temp: VecDeque<char> = VecDeque::new();
            for _i in 0..quantity {
                let val: char = stacks[from as usize].pop_back().unwrap();
                temp.push_back(val);
            }
            for _i in 0..quantity {
                let val: char = temp.pop_back().unwrap();
                stacks[to as usize].push_back(val);
            }
        }
    }
    let mut message: String = String::from("");
    for i in 0..stack_size as usize {
        let val: char = stacks[i].pop_back().unwrap();
        message.push(val);
    }
    println!("Part 2: {}", message);
}
