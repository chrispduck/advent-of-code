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

fn load_input(fname: &str) -> Vec<(String, i64)> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let mut commands: Vec<(String, i64)> = vec![];
    for line in contents.lines() {
        let line_vec: Vec<&str> = line.split(' ').collect();
        let direction: &str = line_vec[0];
        let distance: &str = line_vec[1];
        let distance = distance.parse::<i64>().unwrap();
        commands.push((direction.to_string(), distance));
    }

    return commands;
}

fn part1(v: &Vec<(String, i64)>) -> i64 {
    let mut position = Complex::new(0, 0);

    for (direction, distance) in v {
        position += to_complex(direction) * distance;
    }

    return position.re.abs() * position.im.abs();
}

fn part2(v: &Vec<(String, i64)>) -> i64 {
    let mut position = Complex::new(0, 0);
    let mut aim = 0;

    for (direction, distance) in v {
        match direction.as_str() {
            "forward" => {
                position += Complex::new(*distance, aim * *distance);
            },
            "up" => aim -= distance,
            "down" => aim += distance,
            _ => {}
        }
    }

    return position.re.abs() * position.im.abs();
}

fn to_complex(direction: &str) -> Complex<i64> {
    match direction {
        "up" => Complex::new(0, -1),
        "down" => Complex::new(0, 1),
        "forward" => Complex::new(1, 0),
        _ => Complex::new(0, 0),
    }
}