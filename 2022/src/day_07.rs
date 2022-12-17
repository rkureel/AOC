use std::collections::{VecDeque, HashMap};

fn get_sizes<'a>(dir: String, listings: &HashMap<String, Vec<Vec<&'a str>>>, sizes: &mut HashMap<String, u128>){
    sizes.insert(dir.clone(), 0);
    let line_split: Vec<Vec<&str>> = listings.get(&dir as &str).unwrap().to_vec();
    for line in line_split {
        if line[0] == "dir" {
            let mut new_path: String = dir.clone();
            new_path.push_str(line[1]);
            new_path.push_str("/");
            if !sizes.contains_key(&new_path) {
                get_sizes(new_path.clone(), &listings, sizes); 
            }
            sizes.insert(dir.clone(), *sizes.get(&dir).unwrap() + *sizes.get(&new_path).unwrap());
        }
        else {
            let size: u128 = line[0].parse().unwrap();
            sizes.insert(dir.clone(), *sizes.get(&dir).unwrap()+size);
        }
    }
}

fn get_path(dirs: &VecDeque<&str>) -> String {
    let mut path: String = String::new();
    for dir in dirs {
        path.push_str(&dir);
        path.push('/');
    }
    return path;
}

pub fn part_1(input: &String) {
    let mut listings: HashMap<String, Vec<Vec<&str>>> = HashMap::new();
    let mut dirs: VecDeque<&str> = VecDeque::new();
    for line in input.lines() {
        if line.len() == 0 {
            continue;
        } 
        let line_split: Vec<&str> = line.split(" ").collect();
        if line_split[0] == "$" {
            if line_split[1] == "cd" {
                if line_split[2] == ".." {
                    dirs.pop_back();
                }
                else if line_split[2] == "/" {
                    dirs.clear();
                    dirs.push_back("/");
                }
                else {
                    dirs.push_back(line_split[2]);
                }
            }
            else if line_split[1] == "ls" {
                continue;
            }
        }
        else {
            let dir = get_path(&dirs);
            listings.entry(dir.clone()).or_insert(Vec::new());
            let mut vec: Vec<Vec<&str>> = listings.get(&dir).unwrap().to_vec();
            vec.push(line_split);
            listings.insert(dir, vec);
        }
    }
    let mut sizes: HashMap<String, u128> = HashMap::new();
    get_sizes(String::from("//"), &listings, &mut sizes);
    let mut sum: u128 = 0;
    for dir in listings.keys() {
        if *sizes.get(dir).unwrap() <= 100000 {
            sum += *sizes.get(dir).unwrap();
        }
    }
    println!("Part 1: {}", sum);
}

pub fn part_2(input: &String) {
    let mut listings: HashMap<String, Vec<Vec<&str>>> = HashMap::new();
    let mut dirs: VecDeque<&str> = VecDeque::new();
    for line in input.lines() {
        if line.len() == 0 {
            continue;
        } 
        let line_split: Vec<&str> = line.split(" ").collect();
        if line_split[0] == "$" {
            if line_split[1] == "cd" {
                if line_split[2] == ".." {
                    dirs.pop_back();
                }
                else if line_split[2] == "/" {
                    dirs.clear();
                    dirs.push_back("/");
                }
                else {
                    dirs.push_back(line_split[2]);
                }
            }
            else if line_split[1] == "ls" {
                continue;
            }
        }
        else {
            let dir = get_path(&dirs);
            listings.entry(dir.clone()).or_insert(Vec::new());
            let mut vec: Vec<Vec<&str>> = listings.get(&dir).unwrap().to_vec();
            vec.push(line_split);
            listings.insert(dir, vec);
        }
    }
    let mut sizes: HashMap<String, u128> = HashMap::new();
    get_sizes(String::from("//"), &listings, &mut sizes);
    let free: u128 = 70000000 - *sizes.get("//").unwrap();
    let req: u128 = 30000000 - free;
    let mut minval: u128 = 70000000;
    for dir in listings.keys() {
        if *sizes.get(dir).unwrap() >= req && *sizes.get(dir).unwrap()<minval {
            minval = *sizes.get(dir).unwrap();
        }
    }
    println!("Part 2: {}", minval);
}
