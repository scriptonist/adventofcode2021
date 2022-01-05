use std::{io::BufRead, vec};

fn main() {
    let input = parse(include_str!("../input.txt"));
    println!("{}", part1(&input));
}

fn part1(inputs: &Vec<InputItem>) -> i32 {
    let want_lens = [2, 4, 3, 7];
    let mut count = 0;
    for line in inputs {
        for digit in line.digit {
            if want_lens.contains(&(digit.len() as i32)) {
                count += 1;
            }
        }
    }

    count
}

fn part2(inputs: &Vec<InputItem>) -> i32 {
    
}

#[derive(Debug)]
struct InputItem {
    signals: Vec<String>,
    digit: Vec<String>,
}

fn parse(input_str: &str) -> Vec<InputItem> {
    let mut inputs: Vec<InputItem> = vec![];
    for line in input_str.lines() {
        let mut parts = line.split("|").filter(|v| !v.is_empty());
        let (signals, digit) = (parts.next().unwrap(), parts.next().unwrap());
        inputs.push(InputItem {
            signals: signals
                .split(" ")
                .filter(|v| !v.is_empty())
                .map(|v| v.to_owned())
                .collect::<Vec<String>>(),
            digit: digit
                .split(" ")
                .filter(|v| !v.is_empty())
                .map(|v| v.to_owned())
                .collect::<Vec<String>>(),
        })
    }

    inputs
}
