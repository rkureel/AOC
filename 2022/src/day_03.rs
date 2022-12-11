use std::collections::HashSet;

pub fn part_1(input: &String) {
    let mut priority: u32 = 0;
    for line in input.lines() {
        if line.len() == 0 {
            break;
        }
        let line_vec: Vec<char> = line.chars().collect();
        let mut item_set: HashSet<char> = HashSet::new();
        for i in 0..line_vec.len()/2 {
            item_set.insert(line_vec[i]); 
        }
        for i in line_vec.len()/2..line_vec.len() {
            if item_set.contains(&line_vec[i]) {
                let mut item_p: u32 = 0;
                if line_vec[i]>='a' && line_vec[i]<='z' {
                    item_p = line_vec[i] as u32 - 'a' as u32; 
                } 
                else if line_vec[i] >= 'A' && line_vec[i] <= 'Z'{
                    item_p = line_vec[i] as u32 - 'A' as u32 + 26;
                }
                priority += item_p+1;
                break;
            }
        }
    }
    println!("Part 1: {}", priority);
}

pub fn part_2(input: &String) {
    let mut priority: u32 = 0;
    let mut line_no: u32 = 1;
    let mut set_one: HashSet<char> = HashSet::new();
    let mut set_two: HashSet<char> = HashSet::new();
    let mut set_three: HashSet<char> = HashSet::new();
    for line in input.lines() {
        if line.len() == 0 {
            break;
        }
        let item_set: &mut HashSet<char>;
        if line_no == 1 {
            item_set = &mut set_one;
        }
        else if line_no == 2 {
            item_set = &mut set_two;
        }
        else {
            item_set = &mut set_three;
        }

        for c in line.chars() {
            item_set.insert(c);
        } 
        
        if line_no == 3 {
            for item in set_one.intersection(&set_two) {
                if set_three.contains(item) {
                    if *item >= 'a' && *item <= 'z' {
                        priority += *item as u32 - 'a' as u32 + 1;
                    }
                    else {
                        priority += *item as u32 - 'A' as u32 + 26 + 1;
                    }
                }
            }
        }

        line_no += 1;
        if line_no == 4 {
            line_no = 1;
            set_one.clear();
            set_two.clear();
            set_three.clear();
        }
    }
    println!("Part 2: {}", priority);
}
