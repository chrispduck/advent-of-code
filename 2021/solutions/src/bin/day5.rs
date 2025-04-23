use std::collections::HashMap;
use std::fs::File;
use std::io::Read;
use std::path::Path;

use num_complex::Complex;

fn main() {
    let input_dir = Path::new("data/day5/");

    let commands = load_input(input_dir.join("example_input.txt").to_str().unwrap());
    println!("part1 : {}", part1(commands.clone()));
    println!("part2 : {}", part2(commands));

    let commands = load_input(input_dir.join("input.txt").to_str().unwrap());
    println!("part1 : {}", part1(commands.clone()));
    println!("part2 : {}", part2(commands));
}

fn split_two_u64(s: &str) -> (u64, u64) {
    let parts: Vec<&str> = s.split(",").map(|s| s.trim()).collect();
    let (x, y) = (parts[0], parts[1]);
    let (x, y) = (x.parse::<u64>().unwrap(), y.parse::<u64>().unwrap());
    return (x, y);
}

fn load_input(fname: &str) -> Vec<(Complex<u64>, Complex<u64>)> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let mut commands: Vec<(Complex<u64>, Complex<u64>)> = vec![];
    for line in contents.lines() {
        let parts: Vec<&str> = line.split("->").map(|s| s.trim()).collect();
        let (from, to) = (parts[0], parts[1]);
        let (x1, y1) = split_two_u64(from);
        let (x2, y2) = split_two_u64(to);
        let from_complex: Complex<u64> = Complex::new(x1, y1);
        let to_complex: Complex<u64> = Complex::new(x2, y2);
        commands.push((from_complex, to_complex));
    }
    return commands;
}

fn iterate_between_coords_(
    from: Complex<u64>,
    to: Complex<u64>,
    enable_diag: bool,
) -> Vec<Complex<u64>> {
    let mut coords: Vec<Complex<u64>> = vec![];

    // Convert to signed integers for easier calculations with negative deltas
    let from_re = from.re as i64;
    let from_im = from.im as i64;
    let to_re = to.re as i64;
    let to_im = to.im as i64;

    let dx: i64 = to_re - from_re;
    let dy: i64 = to_im - from_im;

    if dx == 0 {
        // Vertical line
        let step = if dy >= 0 { 1 } else { -1 };
        for delta_y in 0..=dy.abs() {
            let new_y = (from_im + delta_y * step) as u64;
            coords.push(Complex::new(from.re, new_y));
        }
    } else if dy == 0 {
        // Horizontal line
        let step = if dx >= 0 { 1 } else { -1 };
        for delta_x in 0..=dx.abs() {
            let new_x = (from_re + delta_x * step) as u64;
            coords.push(Complex::new(new_x, from.im));
        }
    } else if enable_diag && dx.abs() == dy.abs() {
        // Diagonal line
        let step_x = if dx >= 0 { 1 } else { -1 };
        let step_y = if dy >= 0 { 1 } else { -1 };

        for delta in 0..=dx.abs() {
            let new_x = (from_re + delta * step_x) as u64;
            let new_y = (from_im + delta * step_y) as u64;
            coords.push(Complex::new(new_x, new_y));
        }
    }

    return coords;
}

fn add_to_hashmap(h: &mut HashMap<(u64, u64), u64>, c: Complex<u64>) {
    let c = (c.re, c.im);
    if !h.contains_key(&c) {
        h.insert(c, 0);
    }
    *h.get_mut(&c).unwrap() += 1;
}

fn compute_score(h: &HashMap<(u64, u64), u64>) -> u64 {
    h.values().filter(|&&x| x > 1).count() as u64
}

fn part1(v: Vec<(Complex<u64>, Complex<u64>)>) -> u64 {
    let mut count = HashMap::new();
    for (from, to) in v {
        let coords = iterate_between_coords_(from, to, false);
        for c in coords {
            add_to_hashmap(&mut count, c);
        }
    }
    return compute_score(&count);
}

fn part2(v: Vec<(Complex<u64>, Complex<u64>)>) -> u64 {
    let mut count = HashMap::new();
    for (from, to) in v {
        let coords = iterate_between_coords_(from, to, true);
        for c in coords {
            add_to_hashmap(&mut count, c);
        }
    }
    return compute_score(&count);
}