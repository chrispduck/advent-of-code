use std::collections::HashMap;
use std::fs::File;
use std::io::Read;
use std::path::Path;

const N_DAYS_PT1: usize = 80;
const N_DAYS_PT2: usize = 256;
const NEW_FISH_TIMER: u64 = 8;
const OLD_FISH_TIMER: u64 = 6;

fn main() {
    let input_dir = Path::new("data/day6/");

    let data = load_input(input_dir.join("example_input.txt").to_str().unwrap());
    println!("part1 : {}", part1(data.clone(), N_DAYS_PT1));
    println!("part2 : {}", part2(data, N_DAYS_PT2));
    //
    let data = load_input(input_dir.join("input.txt").to_str().unwrap());
    println!("part1 : {}", part1(data.clone(), N_DAYS_PT1));
    println!("part2 : {}", part2(data, N_DAYS_PT2));
}

fn load_input(fname: &str) -> HashMap<u64, u64> {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let nums: Vec<u64> = contents
        .split(',')
        .map(|s| s.trim().parse::<u64>().unwrap())
        .collect();
    let mut timers: HashMap<u64, u64> = HashMap::new();
    for num in nums {
        let n_option = timers.get(&num);
        if let Some(value) = n_option {
            timers.insert(num, value + 1);
        } else {
            timers.insert(num, 1);
        }
    }
    return timers;
}

fn compute_n_fish(v: HashMap<u64, u64>) -> u64 {
    v.values().sum()
}

fn add(h: &mut HashMap<u64, u64>, key: u64, val: u64) {
    let n_option = h.get(&key);
    if let Some(value) = n_option {
        h.insert(key, value + val);
    } else {
        h.insert(key, val);
    }
}

fn part1(v: HashMap<u64, u64>, n_days: usize) -> u64 {
    return simulate(v, n_days);
}

fn part2(v: HashMap<u64, u64>, n_days: usize) -> u64 {
    return simulate(v, n_days);
}

fn simulate(v: HashMap<u64, u64>, n_days: usize) -> u64 {
    let mut v_today = v;
    for _ in 0..n_days {
        let mut v_next_day = HashMap::new();
        for (timer, n_fish) in v_today {
            if timer == 0 {
                add(&mut v_next_day, NEW_FISH_TIMER, n_fish);
                add(&mut v_next_day, OLD_FISH_TIMER, n_fish);
            } else {
                add(&mut v_next_day, timer - 1, n_fish);
            }
        }
        v_today = v_next_day.clone()
    }

    return compute_n_fish(v_today);
}
