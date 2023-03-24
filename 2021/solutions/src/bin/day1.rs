use std::fs::File;
use std::io::Read;
use std::path::Path;

// use std::env::current_dir;

fn main() {
    let fname = "data/day1/example_input.txt";
    let v = load_input(fname);
    println!("part1 : {}", part1(&v));
    println!("part2 : {}", part2(&v));


    let fname = "data/day1/input.txt";
    let v = load_input(fname);
    println!("part1 : {}", part1(&v));
    println!("part2 : {}", part2(&v));
}

fn load_input(fname: &str)-> Vec<i32>{
    let data_path = Path::new(fname);
    // println!("current_dir: {:?}", current_dir());
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    // println!("contents: {}", contents);


    // Using collect to create a vector
    let _v: Vec<i32> = contents
        .lines()
        .map(|s| s.parse::<i32>().unwrap())
        .collect();


    // Another way to create a vector
    let mut v2: Vec<i32> = vec![];
    for line in contents.lines() {
        let i = line.parse::<i32>().unwrap();
        v2.push(i);
    }
    // println!("v: {:?}", v);

    return v2
}

fn part1(v: &Vec<i32>) -> i32 {
    let mut sum = 0;
    for i in 1..v.len() {
        if v[i-1] < v[i]{
            sum+=1;
        }
    }
    return sum;
}

fn part2(v: &Vec<i32>) -> i32 {
    let mut sum = 0;
    for i in 1..(v.len()-2) {
        let lhs = v[i-1] + v[i] + v[i+1];
        let rhs = v [i] + v[i+1] + v[i+2];
        if lhs < rhs {
            sum+=1;
        }
        // println!("{}", i)
    }
    return sum
}