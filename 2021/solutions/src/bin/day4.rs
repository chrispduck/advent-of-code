use regex::Regex;
use std::fs::File;
use std::io::Read;
use std::path::Path;

fn main() {
    let input_dir = "data/day4/".to_string();

    let (numbers, score_cards) =
        load_input(format!("{}{}", input_dir, "example_input.txt").as_str());
    println!("part1 : {}", part1(numbers, score_cards));
    // println!("part2 : {}", part2(&v));
    //
    // let v = load_input(format!("{}{}", input_dir, "input.txt").as_str());
    // println!("part1 : {}", part1(&v));
    // println!("part2 : {}", part2(&v));
}

#[derive(Debug, Clone)]
struct ScoreCard {
    // fixed size array of 5 uint of 5 uint;
    // 5 rows of 5 columns
    numbers: [[u32; 5]; 5],
    ticked: [[bool; 5]; 5],
}

impl ScoreCard {
    fn new() -> ScoreCard {
        ScoreCard {
            numbers: [[0; 5]; 5],
            ticked: [[false; 5]; 5],
        }
    }
    // tick a number on the scorecard
    fn tick_number(&mut self, number: u32) {
        for i in 0..5 {
            for j in 0..5 {
                if self.numbers[i][j] == number {
                    self.ticked[i][j] = true;
                }
            }
        }
    }

    // check if a row is complete
    fn is_any_row_complete(&self) -> bool {
        for i in 0..5 {
            let mut row_complete = true;
            for j in 0..5 {
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
        for j in 0..5 {
            let mut column_complete = true;
            for i in 0..5 {
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
        for i in 0..5 {
            for j in 0..5 {
                if !self.ticked[i][j] {
                    count += 1;
                }
            }
        }
        return count;
    }
}

fn load_input(fname: &str) -> (Vec<u32>, Vec<ScoreCard>) {
    let data_path = Path::new(fname);
    let mut file = File::open(data_path).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();

    let (numbers_str, mut contents): (&str, &str) = contents.split_once('\n').unwrap();
    let numbers: Vec<u32> = numbers_str
        .split(',')
        .map(|x| x.parse::<u32>().unwrap())
        .collect();
    let mut boards: Vec<ScoreCard> = vec![];

    contents.split_once("\n\n").unwrap(); // skip the empty line
    println!("contents: {}", contents);

    while let Some((mut board_string, rest)) = contents.split_once("\n\n") {
        contents = rest;
        let mut board = ScoreCard::new();
        println!("board_string :{:?}", board_string);
        let mut board_lines = board_string.trim().lines();

        for i in 0..5 {
            if let Some(line) = board_lines.next() {
                // split line and parse into the board
                let re = Regex::new(" +").unwrap();
                println!("line: {:?}", line);
                let nums: Vec<&str> = re.split(line.trim()).filter(|x| !x.is_empty()).collect();
                println!("nums: {:?}", nums);
                for j in 0..5 {
                    board.numbers[i][j] = nums[j].parse::<u32>().unwrap();
                }
            }
        }
        boards.push(board);
    }
    return (numbers, boards);
}

fn part1(numbers: Vec<u32>, score_cards: Vec<ScoreCard>) -> u32 {
    println!("numbers: {:?}", numbers);
    println!("score_cards: {:?}", score_cards);

    let mut score_cards = score_cards.clone();
    for number in numbers {
        for score_card in &mut score_cards {
            score_card.tick_number(number);
            if score_card.is_complete() {
                println!("score_card: {:?}", score_card);
                println!("number: {:?}", number);
                return score_card.count_unmarked_nums() * number as u32;
            }
        }
    }
    return 0;
}

fn part2(v: &Vec<i32>) -> i32 {
    // TODO implement
    return 0;
}
