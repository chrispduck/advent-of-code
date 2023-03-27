use std::fs::File;
use std::io::Read;
use std::path::Path;

use num_complex::Complex;

fn main() {
    // TODO change to day number
    let input_dir = "data/dayN/";

    let v = load_input(input_dir + "example_input.txt");
    println!("part1 : {}", part1(&v));
    println!("part2 : {}", part2(&v));

    let v = load_input(input_dir + "input.txt");
    println!("part1 : {}", part1(&v));
    println!("part2 : {}", part2(&v));
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

fn part1(v: &Vec<i32>) -> i32 {
    // TODO implement
    return 0;
}

fn part2(v: &Vec<i32>) -> i32 {
    // TODO implement
    return 0;
}
