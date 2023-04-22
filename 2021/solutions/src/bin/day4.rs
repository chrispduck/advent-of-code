use regex::Regex;
use std::fs::File;
use std::io::Read;
use std::path::Path;

const N_ROWS: usize = 5;
const N_COLS: usize = 5;

fn main() {
    let input_dir = Path::new("data/day4/");

    let (numbers, score_cards) = load_input(input_dir.join("example_input.txt").to_str().unwrap());
    println!("part1 : {}", part1(numbers.clone(), score_cards.clone()));
    println!("part2 : {}", part2(numbers.clone(), score_cards.clone()));

    let (numbers, score_cards) = load_input(input_dir.join("input.txt").to_str().unwrap());
    println!("part1 : {}", part1(numbers.clone(), score_cards.clone()));
    println!("part2 : {}", part2(numbers.clone(), score_cards.clone()));
}

#[derive(Debug, Clone)]
struct ScoreCard {
    // fixed size array of 5 uint of 5 uint;
    // 5 rows of 5 columns
    numbers: [[u32; N_COLS]; N_ROWS],
    ticked: [[bool; N_COLS]; N_ROWS],
}

impl ScoreCard {
    fn new() -> ScoreCard {
        ScoreCard {
            numbers: [[0; N_COLS]; N_ROWS],
            ticked: [[false; N_COLS]; N_ROWS],
        }
    }
    // tick a number on the scorecard
    fn tick_number(&mut self, number: u32) -> bool {
        let mut ticked = false;
        for i in 0..N_ROWS {
            for j in 0..N_COLS {
                if self.numbers[i][j] == number {
                    self.ticked[i][j] = true;
                    ticked = true;
                }
            }
        }
        return ticked;
    }

    // check if a row is complete
    fn is_any_row_complete(&self) -> bool {
        for i in 0..N_ROWS {
            let mut row_complete = true;
            for j in 0..N_COLS {
                if !self.ticked[i][j] {
                    row_complete = false;
                }
            }
            if row_complete {
                return true;
            }
        }
        return false;
    }

    // check if a column is complete
    fn is_any_column_complete(&self) -> bool {
        for j in 0..N_COLS {
            let mut column_complete = true;
            for i in 0..N_ROWS {
                if !self.ticked[i][j] {
                    column_complete = false;
                }
            }
            if column_complete {
                return true;
            }
        }
        return false;
    }

    // check if the scorecard is complete
    fn is_complete(&self) -> bool {
        return self.is_any_row_complete() || self.is_any_column_complete();
    }

    fn count_unmarked_nums(&self) -> u32 {
        let mut count = 0;
        for i in 0..N_ROWS {
            for j in 0..N_COLS {
                if !self.ticked[i][j] {
                    count += self.numbers[i][j];
                }
            }
        }
        return count;
    }
}

fn load_input(filename: &str) -> (Vec<u32>, Vec<ScoreCard>) {
    let data_path = Path::new(filename);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    // first line is the numbers
    let (numbers_str, mut contents): (&str, &str) = contents.split_once('\n').unwrap();
    let numbers: Vec<u32> = numbers_str
        .split(',')
        .map(|x| x.parse::<u32>().unwrap())
        .collect();
    let mut boards: Vec<ScoreCard> = vec![];

    // read the scorecards
    while let Some((board_string, rest)) = contents.split_once("\n\n") {
        contents = rest;
        let mut board = ScoreCard::new();
        let mut board_lines = board_string.trim().lines();

        let re = Regex::new(" +").unwrap();
        for i in 0..N_ROWS {
            if let Some(line) = board_lines.next() {
                let nums: Vec<&str> = re.split(line.trim()).filter(|x| !x.is_empty()).collect();
                for j in 0..N_COLS {
                    board.numbers[i][j] = nums[j].parse::<u32>().unwrap();
                }
            }
        }
        boards.push(board);
    }
    return (numbers, boards);
}

fn part1(numbers: Vec<u32>, mut score_cards: Vec<ScoreCard>) -> u32 {
    // return score of first card that wins
    for number in numbers {
        for score_card in &mut score_cards {
            score_card.tick_number(number);
            if score_card.is_complete() {
                return score_card.count_unmarked_nums() * number as u32;
            }
        }
    }
    return 0;
}

fn part2(numbers: Vec<u32>, mut score_cards: Vec<ScoreCard>) -> u32 {
    // return the score of the card that would win last
    for number in numbers {
        let len_before = score_cards.len();
        for score_card in &mut score_cards {
            score_card.tick_number(number);
            if score_card.is_complete() && len_before == 1 {
                return score_card.count_unmarked_nums() * number as u32;
            }
        }
        score_cards.retain(|x| !x.is_complete());
    }
    return 0;
}
