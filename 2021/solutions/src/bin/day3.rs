use std::fs::File;
use std::io::Read;
use std::ops::BitXor;
use std::path::Path;

fn main() {
    let input_dir = "data/day3/".to_string();

    let v = load_input(format!("{}{}", input_dir, "example_input.txt").as_str());
    println!("part1 : {}", part1(&v, 5));
    println!("part2 : {}", part2(&v, 5));

    let v = load_input(format!("{}{}", input_dir, "input.txt").as_str());
    println!("part1 : {}", part1(&v, 12));
    println!("part2 : {}", part2(&v, 12));
}

fn load_input(fname: &str) -> Vec<u32> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let mut data: Vec<u32> = vec![];
    for line in contents.lines() {
        data.push(u32::from_str_radix(&line, 2).unwrap());
    }
    return data;
}

fn part1(v: &Vec<u32>, n_bit: u8) -> u32 {
    let n_lines = v.len();
    let mut gamma: u32 = 0;
    for exponent in 0..n_bit {
        let mask = 2_u32.pow(exponent as u32);
        let mut count: u32 = 0;
        for line in v {
            if *line & mask != 0 {
                count += 1;
            }
        }
        if count as f64 / n_lines as f64 > 0.5 {
            gamma += mask;
        }
    }

    let epsilon: u32 = gamma.bitxor(2_u32.pow(n_bit as u32) - 1);
    return gamma * epsilon;
}

fn part2(nums: &Vec<u32>, n_bits: u8) -> u32 {
    // oxygen rating
    let width = n_bits as usize;
    let oxy = find_rating(&nums, width, true);
    let co2 = find_rating(&nums, width, false);
    let life_support_rating = oxy * co2;
    return life_support_rating;
}

fn find_rating(nums: &[u32], width: usize, most_common: bool) -> u32 {
    let mut nums = nums.to_owned();

    for i in (0..width).rev() {
        let ones = nums.iter().filter(|n| **n & 1 << i > 0).count();
        let zeros = nums.len() - ones;
        if most_common && ones >= zeros || !most_common && ones < zeros {
            nums.retain(|n| *n & 1 << i > 0);
        } else {
            nums.retain(|n| *n & 1 << i == 0);
        }
        if nums.len() == 1 {
            break;
        }
    }
    nums[0]
}
