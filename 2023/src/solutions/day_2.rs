use std::cmp::max;

pub fn solve(input: &Vec<String>) {
    println!("Solving for day 2");
    let games: Vec<Game> = parse(input);
    
    println!("Part 1 solution: {}", part_1(&games));
    println!("Part 2 solution: {}", part_2(&games));
}

fn parse(input: &Vec<String>) -> Vec<Game> {
    input
        .iter()
        .map(|line| {
            let (game_id, content) = line.split_once(": ").unwrap();
            let game_id: u32 = game_id
                .split_ascii_whitespace()
                .last()
                .unwrap()
                .parse()
                .unwrap();
            
            let cube_sets: Vec<CubeSet> = content
                .split(';')
                .map(|turn| {
                    let mut cube_set = CubeSet{n_red: 0, n_green: 0, n_blue: 0};
                    turn
                        .trim()
                        .split(",")
                        .for_each(|cube| {
                            let mut cube = cube.trim().split(" ");
                            let cube_count: u8 = cube.next().unwrap().parse().unwrap();
                            let cube_color = match cube.next().unwrap() {
                                "red" => Some(Color::Red),
                                "green" => Some(Color::Green),
                                "blue" => Some(Color::Blue),
                                _ => None
                            };
                            if let Some(color) = cube_color {
                                cube_set.set_n_color(&color, cube_count);
                            }
                        });
                    cube_set
                })
                .collect();

            Game { id: game_id, cube_sets: cube_sets }
        })
        .collect()
}

fn part_1(games: &Vec<Game>) -> u32 {
    const MAX_N_RED: u8 = 12;
    const MAX_N_GREEN: u8 = 13;
    const MAX_N_BLUE: u8 = 14;

    let possible_games_sum: u32 = games
        .iter()
        .filter_map(|game| {
            let is_game_valid = game.cube_sets.iter().all(|cube_set| {
                cube_set.n_red <= MAX_N_RED 
                && cube_set.n_green <= MAX_N_GREEN
                && cube_set.n_blue <= MAX_N_BLUE
            });

            if is_game_valid {
                Some(game.id)
            } else {
                None
            }
        })
        .sum();
    possible_games_sum
}

fn part_2(games: &Vec<Game>) -> u32 {
    let power_sum: u32 = games.iter().map(|game| {
        let mut max_observed_red: u8 = u8::MIN;
        let mut max_observed_green: u8 = u8::MIN;
        let mut max_observed_blue: u8 = u8::MIN;

        game.cube_sets
            .iter()
            .for_each(|cube_set| {
                max_observed_red = max(cube_set.n_red, max_observed_red);
                max_observed_green = max(cube_set.n_green, max_observed_green);
                max_observed_blue = max(cube_set.n_blue, max_observed_blue);
            });
        u32::from(max_observed_red) * u32::from(max_observed_blue) * u32::from(max_observed_green)
    
    })
    .sum();
    power_sum 
}

#[derive(Debug)]
struct Game {
    id: u32,
    cube_sets: Vec<CubeSet>,
}

#[derive(Debug)]
struct CubeSet {
    n_red: u8,
    n_green: u8,
    n_blue: u8,
}

impl CubeSet {
    pub fn set_n_color(&mut self, color: &Color, n: u8) {
        match color {
            Color::Red => self.n_red = n,
            Color::Green => self.n_green = n,
            Color::Blue => self.n_blue = n,
        }
    }
}

#[derive(Debug)]
enum Color {
    Red,
    Green,
    Blue,
}

#[cfg(test)]
mod tests {

    static INPUT: &str = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

    
    #[test]
    fn day_2_part_1() {
        assert_eq!(super::part_1(&super::parse(&crate::generate_test_input(INPUT))), 8); 
    }

    #[test]
    fn day_2_part_2() {
        assert_eq!(super::part_2(&super::parse(&crate::generate_test_input(INPUT))), 2286); 
    }
}
