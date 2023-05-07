use ndarray::Array;
use std::fs::File;
use std::io::Read;
use std::path::Path;

use num_complex::Complex;

fn main() {
    // TODO change to day number
    let input_dir = Path::new("data/day7/");

    let data = load_input(input_dir.join("example_input.txt").to_str().unwrap());
    println!("part1 : {}", part1(&data));
    // println!("part2 : {}", part2(data));
    //
    let data = load_input(input_dir.join("input.txt").to_str().unwrap());
    println!("part1 : {}", part1(&data));
    // println!("part2 : {}", part2(data));
}

fn load_input(fname: &str) -> Vec<i64> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    let v = contents
        .split(',')
        .map(|s| s.trim().parse::<i64>().unwrap())
        .collect();
    return v;
}

fn part1(v: &Vec<i64>) -> i64 {
    let min: i64 = v.iter().min().unwrap().clone();
    let max: i64 = v.iter().max().unwrap().clone();
    let mut res = i64::MAX;
    for pos in min..max + 1 {
        let fuel: i64 = v.iter().map(|x| (x - pos).abs()).sum();
        println!("pos: {}, fuel: {}", pos, fuel);
        res = res.min(fuel);
    }
    return res;
}

fn part2(v: Vec<u64>) -> u64 {
    // TODO implement
    return 0;
}
