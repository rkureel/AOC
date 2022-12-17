use std::collections::HashSet;

pub fn part_1(input: &String) {
    let vec_string: Vec<char> = input.chars().collect();
    let mut index: u32 = 0;
    for i in 3..input.len() {
        index = i as u32;
        if vec_string[i] != vec_string[i-1] && vec_string[i] != vec_string[i-2] && vec_string[i] != vec_string[i-3] && vec_string[i-1] != vec_string[i-2] && vec_string[i-1] != vec_string[i-3] && vec_string[i-2] != vec_string[i-3] {
            break;
        }
    }
    println!("Part 1: {}", index+1);
}

pub fn part_2(input: &String) {
    let vec_string: Vec<char> = input.chars().collect();
    let mut index: u32 = 0;
    for i in 13..input.len() {
        index = i as u32;
        let mut set: HashSet<char> = HashSet::new();
        for j in 0..14 as usize {
            set.insert(vec_string[i-j]);
        }
        if set.len() == 14 {
            break;
        }
    }
    println!("Part 2: {}", index+1);
}
