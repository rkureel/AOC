pub fn part_1(input: &String) {
    let mut grid: Vec<Vec<u8>> = Vec::new();
    let mut visible: Vec<Vec<u8>> = Vec::new();
    for (row, line) in input.lines().enumerate() {
        if line.len() == 0 {
            break;
        }
        grid.push(Vec::new());
        visible.push(Vec::new());
        for char in line.chars() {
            grid[row].push(char as u8 - '0' as u8);
            visible[row].push(0);
        }
    }
    let mut visible_count: u32 = 0;
    for row in 0..grid.len() {
        let mut max_so_far: u8 = 0;
        for col in 0..grid[0].len() {
            if row == 0 || row == grid.len()-1 || col == 0 || col == grid[0].len()-1 {
                visible[row][col] = 1;
            }
            if max_so_far<grid[row][col] {
                visible[row][col] = 1;
                max_so_far = grid[row][col];
            }
        }

        max_so_far = 0;
        for col in (0..grid[0].len()).rev() {
            if max_so_far<grid[row][col] {
                visible[row][col] = 1;
                max_so_far = grid[row][col];
            }
        }
    }
    for col in 0..grid[0].len() {
        let mut max_so_far: u8 = 0;
        for row in 0..grid.len() {
            if max_so_far<grid[row][col] {
                visible[row][col] = 1;
                max_so_far = grid[row][col];
            }
        }

        max_so_far = 0;
        for row in (0..grid.len()).rev() {
            if max_so_far<grid[row][col] {
                visible[row][col] = 1;
                max_so_far = grid[row][col];
            }
        }
    }
    for row in 0..grid.len() {
        for col in 0..grid[0].len() {
            visible_count += visible[row][col] as u32;
        }
    }
    println!("Part 1: {}", visible_count);
}

pub fn part_2(input: &String) {
    let mut grid: Vec<Vec<u8>> = Vec::new();
    for (row, line) in input.lines().enumerate() {
        if line.len() == 0 {
            break;
        }
        grid.push(Vec::new());
        for char in line.chars() {
            grid[row].push(char as u8 - '0' as u8);
        }
    }
    let mut max_scenic_score: i32 = 0;
    for row in 0..grid.len() as i32 {
        for col in 0..grid[0].len() as i32 {
            let mut up = row-1;
            let mut down = row+1;
            let mut right = col+1;
            let mut left = col-1;
            while up>=0 && grid[up as usize][col as usize]<grid[row as usize][col as usize] {
                up-=1;
            }
            while down<grid.len() as i32 && grid[down as usize][col as usize]<grid[row as usize][col as usize] {
                down+=1;
            }
            while right<grid[0].len() as i32 && grid[row as usize][right as usize]<grid[row as usize][col as usize] {
                right+=1;
            }
            while left>=0 && grid[row as usize][left as usize]<grid[row as usize][col as usize] {
                left-=1;
            }
            if up < 0 {
                up = 0;
            }
            if left < 0 {
                left = 0;
            }
            if right as usize >= grid[0].len() {
                right = (grid[0].len()-1) as i32;
            }
            if down as usize >= grid.len() {
                down = (grid.len()-1) as i32;
            }
            let score: i32 = (row-up)*(col-left)*(right-col)*(down-row);
            if score > max_scenic_score {
                max_scenic_score = score;
            }
        }
    }
    println!("Part 2: {}", max_scenic_score);
}   
