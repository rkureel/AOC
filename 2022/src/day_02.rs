pub fn part_1(input: &String) {
    let mut score: u32 = 0;
    for line in input.lines() {
        let char_vec: Vec<char> = line.chars().collect(); 
        let opp_choice: char = char_vec[0];
        let player_choice: char = char_vec[2];

        // if the opponent chooses rock
        if opp_choice == 'A' {
            if player_choice == 'X' {
                score += 3;
            }
            else if player_choice == 'Y' {
                score += 6;
            }
        }
        // if the opponent chooses paper
        else if opp_choice == 'B' {
            if player_choice == 'Y' {
                score += 3;
            }
            else if player_choice == 'Z' {
                score += 6;
            }
        }

        // if the opponent chooses scissors
        else {
            if player_choice == 'X'  {
                score += 6;
            }
            else if player_choice == 'Z' {
                score += 3;
            }
        }

        if player_choice == 'X' {
            score += 1;
        } 
        else if player_choice == 'Y' {
            score += 2;
        }
        else {
            score += 3;
        }
    }
    println!("Part 1: {}", score);
}

pub fn part_2(input: &String) {
    let mut score: u32 = 0;
    for line in input.lines() {
        let char_vec: Vec<char> = line.chars().collect(); 
        let opp_choice: char = char_vec[0];
        let result: char = char_vec[2];
        let player_choice: char;
        // if the opponent chooses rock
        if opp_choice == 'A' {
            if result == 'X' {
                player_choice = 'Z';
                score += 0;
            }
            else if result == 'Y' {
                player_choice = 'X';
                score += 3;
            }
            else {
                player_choice = 'Y';
                score += 6;
            }
        }
        // if the opponent chooses paper
        else if opp_choice == 'B' {
            if result == 'X' {
                player_choice = 'X';
                score += 0;
            }
            else if result == 'Y' {
                player_choice = 'Y';
                score += 3;
            }
            else {
                player_choice = 'Z';
                score += 6;
            }
        }

        // if the opponent chooses scissors
        else {
            if result == 'X' {
                player_choice = 'Y';
                score += 0;
            }
            else if result == 'Y' {
                player_choice = 'Z';
                score += 3;
            }
            else {
                player_choice = 'X';
                score += 6;
            }
        }

        if player_choice == 'X' {
            score += 1;
        } 
        else if player_choice == 'Y' {
            score += 2;
        }
        else {
            score += 3;
        }
    }
    println!("Part 2: {}", score);
}

