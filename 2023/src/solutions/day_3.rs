pub fn solve(input: &Vec<String>) {
    let engine: Engine = parse(input);
    dbg!(engine);
}

fn parse(input: &Vec<String>) -> Engine {
    let mut numbers: Vec<Number> = Vec::new();
    let mut symbols: Vec<Symbol> = Vec::new();
    let mut line_count: u32 = 0;
    input.iter().for_each(|line| {
        let mut number_start_index: u32 = u32::MAX;
        let mut number_value: u32 = 0;
        for (index, c) in line.char_indices() {
            let index_u32: u32 = u32::try_from(index).unwrap();
            if c.is_ascii_digit() {
                if number_start_index == u32::MAX {
                    number_start_index = index_u32;
                }
                number_value = number_value*10 + c.to_digit(10).unwrap();
            } else {
                if number_start_index != u32::MAX {
                    let number_position: Postition = Postition { x_start: number_start_index, x_end: index_u32, y: line_count };
                    let number: Number = Number { value: number_value, position: number_position };
                    number_start_index = u32::MAX;
                    number_value = 0;
                    numbers.push(number);
                }
                if c != '.' {
                    let symbol_position: Postition = Postition { x_start: index_u32, x_end: index_u32+1, y: line_count };
                    let symbol: Symbol = Symbol { value: c, position: symbol_position };
                    symbols.push(symbol);
                }
            }
        }
        
        if number_start_index != u32::MAX {
            let number_position: Postition = Postition { x_start: number_start_index, x_end: u32::try_from(line.len()).unwrap(), y: line_count };
            let number: Number = Number { value: number_value, position: number_position };
            numbers.push(number);
        }

        line_count += 1;
    });
    Engine { numbers, symbols }
}

#[derive(Debug)]
struct Engine {
    numbers: Vec<Number>,
    symbols: Vec<Symbol>,
}

#[derive(Debug)]
struct Number {
    value: u32,
    position: Postition, 
}

#[derive(Debug)]
struct Symbol {
    value: char,
    position: Postition,
}

#[derive(Debug)]
struct Postition {
    x_start: u32,
    x_end: u32,
    y: u32,
}
