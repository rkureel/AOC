use std::collections::HashSet;

struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn to_str(&self)->String {
        format!("({}, {})", self.x, self.y)
    }

    fn shift(&mut self, delta_x: i32, delta_y:i32) {
        self.x += delta_x;
        self.y += delta_y;
    }

    fn move_to(&mut self, x: i32, y: i32) {
        self.x = x;
        self.y = y;
    }
}

pub fn part_1(input: &String) {
    let mut positions: HashSet<String> = HashSet::new();
    let mut head: Point = Point{x: 0, y: 0};
    let mut tail: Point = Point{x: 0, y: 0};
    positions.insert(tail.to_str());
    
    for line in input.lines() {
        if line.len() == 0 {
            break;
        }
        let line_vec: Vec<&str> = line.split(" ").collect();
        let delta_x: i32;
        let delta_y: i32;
        let delta: i32 = line_vec[1].parse().unwrap();
        match line_vec[0] {
            "L" => {
                delta_x = -1;
                delta_y = 0;
            },
            "R" => {
                delta_x = 1;
                delta_y = 0;
            },
            "U" => {
                delta_x = 0;
                delta_y = 1;
            },
            "D" => {
                delta_x = 0;
                delta_y = -1;
            },
            _ => panic!(),
        }
        for _ in 0..delta {
            head.shift(delta_x, delta_y);
            if (head.x-tail.x).abs()>1 {
                tail.move_to(head.x - (head.x-tail.x)/2, head.y);
            }
            else if (head.y-tail.y).abs()>1 {
                tail.move_to(head.x, head.y - (head.y-tail.y)/2);
            }
            positions.insert(tail.to_str());
        }
    }
    println!("Part 1: {}", positions.len());
}

pub fn part_2(input: &String) {
    let mut positions: HashSet<String> = HashSet::new();
    let mut points: Vec<Point> = Vec::new();
    for _ in 0..10 {
        points.push(Point{x: 0, y: 0});
    }
    positions.insert(points[9].to_str());
    
    for line in input.lines() {
        if line.len() == 0 {
            break;
        }
        let line_vec: Vec<&str> = line.split(" ").collect();
        let delta_x: i32;
        let delta_y: i32;
        let delta: i32 = line_vec[1].parse().unwrap();
        match line_vec[0] {
            "L" => {
                delta_x = -1;
                delta_y = 0;
            },
            "R" => {
                delta_x = 1;
                delta_y = 0;
            },
            "U" => {
                delta_x = 0;
                delta_y = 1;
            },
            "D" => {
                delta_x = 0;
                delta_y = -1;
            },
            _ => panic!(),
        }
        for _ in 0..delta {
            points[0].shift(delta_x, delta_y);
            for i in 1..10 {
                if (points[i-1].x-points[i].x).abs()>1{
                    let new_x: i32 = points[i-1].x - (points[i-1].x-points[i].x)/2;
                    let new_y: i32 = if points[i].y==points[i-1].y {
                        points[i-1].y
                    } else {
                        points[i].y + if points[i-1].y>points[i].y {1} else {-1}
                    };
                    points[i].move_to(new_x, new_y);
                } 
                else if (points[i-1].y - points[i].y).abs()>1 {
                    let new_x: i32 = if points[i].x==points[i-1].x {
                        points[i-1].x
                    } else {
                        points[i].x + if points[i-1].x>points[i].x {1} else {-1}
                    };
                    let new_y: i32 = points[i-1].y - (points[i-1].y-points[i].y)/2;
                    points[i].move_to(new_x, new_y);
                }
            }
        
            positions.insert(points[9].to_str());
        }
    }
    println!("Part 2: {}", positions.len());
}
