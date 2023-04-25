use std::fs::File;
use std::io::Read;
use std::path::Path;

use num_complex::Complex;

fn main() {
    // TODO change to day number
    let input_dir = Path::new("data/dayN/");

    let data = load_input(input_dir.join("example_input.txt").to_str().unwrap());
    println!("part1 : {}", part1(data.clone()));
    println!("part2 : {}", part2(data));

    let data = load_input(input_dir.join("input.txt").to_str().unwrap());
    println!("part1 : {}", part1(data.clone()));
    println!("part2 : {}", part2(data));
}

fn load_input(fname: &str) -> Vec<i32> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let mut commands: Vec<i32> = vec![];
    for line in contents.lines() {
        // TODO parse line
    }
    return commands;
}

fn part1(v: Vec<i32>) -> i32 {
    // TODO implement
    return 0;
}

fn part2(v: Vec<i32>) -> i32 {
    // TODO implement
    return 0;
}
