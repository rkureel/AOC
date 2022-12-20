
#[derive(Debug)]
struct CPU {
    clock: u32,
    reg: i32,
    signal_str: i32,
    crt: Vec<Vec<char>>,
}

impl CPU {
    fn execute(&mut self, instr: Instruction) {
        if self.clock%40 == 20 {
            self.signal_str += self.clock as i32 * self.reg;
        }
        
        if self.reg%40 == (self.clock as i32 - 1)%40 || self.reg%40 == (self.clock as i32 - 2)%40 || self.reg%40 == (self.clock as i32)%40 {
            self.crt[((self.clock-1)/40) as usize][((self.clock-1)%40) as usize] = '#';
        }

        match instr.i_type {
            InstructionType::NOOP => {},
            InstructionType::ADDX => {
                self.clock += 1;
                if self.clock%40 == 20 {
                    self.signal_str += self.clock as i32 * self.reg;
                }
                
                if self.reg == (self.clock as i32 - 1)%40 || self.reg == (self.clock as i32 - 2)%40 || self.reg == (self.clock as i32)%40 {
                    self.crt[((self.clock-1)/40) as usize][((self.clock-1)%40) as usize] = '#';
                }
            },
        }
        self.clock += 1;
        self.reg += instr.value;
    }

    fn display_crt(&self) {
        for row in 0..6 {
            for col in 0..40 {
                print!("{}", self.crt[row][col]);
            }
            println!("");
        }
    }
}

#[derive(Debug)]
enum InstructionType {
    NOOP,
    ADDX,
}

fn get_instruction_type(s: &str) -> InstructionType {
    match s {
        "noop" => InstructionType::NOOP,
        "addx" => InstructionType::ADDX,
        _ => panic!(),
    }
}

#[derive(Debug)]
struct Instruction {
    i_type: InstructionType,
    value: i32,
}

fn parse_instruction(line: &str) -> Instruction {
    let line_split: Vec<&str> = line.split(" ").collect();
    match get_instruction_type(line_split[0]) {
        InstructionType::NOOP => Instruction{i_type:InstructionType::NOOP, value:0},
        InstructionType::ADDX => Instruction{i_type:InstructionType::ADDX, value:line_split[1].parse().unwrap()}
    }
}

fn parse_input(input: &String) -> Vec<Instruction> {
    input
        .split("\n")
        .filter(|l| !l.is_empty())
        .map(|l| parse_instruction(l))
        .collect()
}

pub fn part_1(input: &String) {
    let instructions: Vec<Instruction> = parse_input(input);
    let mut cpu = CPU{
        clock: 1,
        reg: 1,
        signal_str: 0,
        crt: vec![vec!['.'; 40];6],
    };
    for instruction in instructions {
        cpu.execute(instruction);
    }
    println!("Part 1: {}", cpu.signal_str);
}

pub fn part_2(input: &String) {
    let instructions: Vec<Instruction> = parse_input(input);
    let mut cpu = CPU{
        clock: 1,
        reg: 1,
        signal_str: 0,
        crt: vec![vec!['.'; 40];6],
    };
    for instruction in instructions {
        cpu.execute(instruction);
    }
    println!("Part 2:");
    cpu.display_crt();
}
