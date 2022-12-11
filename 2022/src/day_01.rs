pub fn part_1(input: &String) {
    let mut max_sum: i32 = 0;
    let mut cur_sum: i32 = 0;
    for line in input.lines() {
        if line.chars().count()>0 {
            let calories: i32 = line.trim().parse()
            .expect("Can't parse number!");
            cur_sum += calories;
        }
        else {
            if cur_sum>max_sum {
                max_sum = cur_sum;
            }
            cur_sum = 0;
        }
    }
    if cur_sum>max_sum {
        max_sum = cur_sum;
    }
    println!("Part 1: {max_sum}");
}

pub fn part_2(input: &String) {
    let mut calories_vec: Vec<i32> = Vec::new();
    let mut cur_sum: i32 = 0;
    for line in input.lines() {
        if line.chars().count()>0 {
            let calories:i32 = line.trim().parse()
            .expect("Can't parse number!");
            cur_sum += calories;
        }
        else {
            calories_vec.push(cur_sum);
            cur_sum = 0;
        }
    }
    if cur_sum > 0 {
        calories_vec.push(cur_sum);
    }
    calories_vec.sort();
    let n = calories_vec.len();
    let total = calories_vec[n-1] + calories_vec[n-2] + calories_vec[n-3];
    println!("Part 2: {total}");
}
