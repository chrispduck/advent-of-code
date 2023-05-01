use std::fs::File;
use std::io::Read;
use std::path::Path;

use num_complex::Complex;

fn main() {
    let fname = "data/day2/example_input.txt";
    let v = load_input(fname);
    println!("part1 : {}", part1(&v));
    println!("part2 : {}", part2(&v));

    let fname = "data/day2/input.txt";
    let v = load_input(fname);
    println!("part1 : {}", part1(&v));
    println!("part2 : {}", part2(&v));
}

fn load_input(fname: &str) -> Vec<Complex<u64>> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let mut commands: Vec<Complex<u64>> = vec![];
    for line in contents.lines() {
        let line_vec: Vec<&str> = line.split(' ').collect();
        let direction: &str = line_vec[0];
        let distance: &str = &line_vec[1].to_lowercase();
        // println!("direction: {}, distance: {}", direction, distance);
        let complex = to_complex(direction);
        let distance = distance.parse::<u64>().unwrap();
        let complex = complex * distance;
        commands.push(complex);
    }
    // println!("commands: {:?}", commands);

    return commands;
}

fn part1(v: &Vec<Complex<u64>>) -> u64 {
    let mut position = Complex::new(0, 0);
    for command in v {
        position += command;
    }
    return position.re.abs() * position.im.abs();
}

fn part2(v: &Vec<Complex<u64>>) -> u64 {
    let mut position = Complex::new(0, 0);
    let mut aim = Complex::new(1, 0);

    for command in v {
        if command.im != 0 {
            aim.im += command.im;
        }
        if command.re != 0 {
            position += aim * command.re;
        }
        // println!("position: {:?}, aim: {}", position, aim);
    }
    return position.re.abs() * position.im.abs();
}

fn to_complex(direction: &str) -> Complex<u64> {
    match direction {
        "up" => Complex::new(0, -1),
        "down" => Complex::new(0, 1),
        "forward" => Complex::new(1, 0),
        _ => Complex::new(0, 0),
    }
}
