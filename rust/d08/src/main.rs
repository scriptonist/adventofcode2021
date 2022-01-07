use std::collections::{HashMap, HashSet};

fn main() {
    let input = parse(include_str!("../input.txt"));
    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(inputs: &Vec<InputItem>) -> i32 {
    let want_lens = [2, 4, 3, 7];
    let mut count = 0;
    for line in inputs {
        for digit in &line.digit {
            if want_lens.contains(&(digit.len() as i32)) {
                count += 1;
            }
        }
    }

    count
}

fn get_number(n: [bool; 7]) -> i32 {
    match n {
        [true, true, true, false, true, true, true] => 0,
        [false, false, true, false, false, true, false] => 1,
        [true, false, true, true, true, false, true] => 2,
        [true, false, true, true, false, true, true] => 3,
        [false, true, true, true, false, true, false] => 4,
        [true, true, false, true, false, true, true] => 5,
        [true, true, false, true, true, true, true] => 6,
        [true, false, true, false, false, true, false] => 7,
        [true, true, true, true, true, true, true] => 8,
        [true, true, true, true, false, true, true] => 9,
        _ => unreachable!(),
    }
}

fn part2(inputs: &Vec<InputItem>) -> i32 {
    let mut s = 0;
    let mut seven_segment_region_map: HashMap<char, usize> = HashMap::new();
    for item in inputs {
        let one = item
            .signals
            .iter()
            .filter(|v| v.len() == 2)
            .nth(0)
            .unwrap()
            .chars()
            .collect::<HashSet<char>>();
        let seven = item
            .signals
            .iter()
            .filter(|v| v.len() == 3)
            .nth(0)
            .unwrap()
            .chars()
            .collect::<HashSet<char>>();
        let four = item
            .signals
            .iter()
            .filter(|v| v.len() == 4)
            .nth(0)
            .unwrap()
            .chars()
            .collect::<HashSet<char>>();

        let eight = item
            .signals
            .iter()
            .filter(|v| v.len() == 7)
            .nth(0)
            .unwrap()
            .chars()
            .collect::<HashSet<char>>();

        /*
        1)
                .
                .

                .
                .
        7)
            ....
                .
                .

                .

        4)
            .   .
            .   .
             ....
                .
                .

        - Get the upper segement from 1 and 7


        */

        let upper = seven.difference(&one).nth(0).unwrap();

        let three_five_two: Vec<HashSet<char>> = item
            .signals
            .iter()
            .filter(|v| v.len() == 5)
            .take(3)
            .map(|item| item.chars().collect::<HashSet<char>>())
            .collect();
        // set of letters which are in three_five_two
        let chars_three_five_two = three_five_two
            .iter()
            .cloned()
            .reduce(|a, v| {
                a.iter()
                    .chain(v.iter())
                    .map(|v| v.to_owned())
                    .collect::<HashSet<char>>()
            })
            .unwrap();

        let common_chars: HashSet<char> = chars_three_five_two
            .iter()
            .filter(|v| {
                three_five_two
                    .iter()
                    .filter(|c| c.contains(v))
                    .map(|v| v.to_owned())
                    .collect::<Vec<HashSet<char>>>()
                    .len()
                    == 3
            })
            .map(|v| v.to_owned())
            .collect();

        let middle = common_chars
            .iter()
            .filter(|v| v != &upper)
            .filter(|v| four.contains(v))
            .nth(0)
            .unwrap();

        let lower = common_chars
            .iter()
            .filter(|v| v != &upper && v != &middle)
            .nth(0)
            .unwrap();

        let upper_left = four
            .iter()
            .filter(|v| !one.contains(v) && v != &middle)
            .nth(0)
            .unwrap();

        let lower_left = eight
            .iter()
            .filter(|v| {
                v != &upper && v != &lower && v != &upper_left && v != &middle && !one.contains(v)
            })
            .nth(0)
            .unwrap();

        let two = three_five_two
            .iter()
            .filter(|v| v.contains(lower_left))
            .nth(0)
            .unwrap();

        let upper_right = two.iter().filter(|v| one.contains(v)).nth(0).unwrap();
        let lower_right = one.iter().filter(|v| v != &upper_right).nth(0).unwrap();
        // println!(
        //     "7:{:?} 4:{:?} 1: {:?} {:?} {:?} {} {} {} {} {}",
        //     seven,
        //     four,
        //     one,
        //     common_chars,
        //     middle,
        //     lower,
        //     upper_left,
        //     lower_left,
        //     upper_right,
        //     lower_right
        // );

        // (0,1,2,3,4,5,6,7)
        seven_segment_region_map.insert(*upper, 0);
        seven_segment_region_map.insert(*upper_left, 1);
        seven_segment_region_map.insert(*upper_right, 2);
        seven_segment_region_map.insert(*middle, 3);
        seven_segment_region_map.insert(*lower_left, 4);
        seven_segment_region_map.insert(*lower_right, 5);
        seven_segment_region_map.insert(*lower, 6);

        let mut number: Vec<String> = vec![];
        for n in &item.digit {
            let mut seven: [bool; 7] = [false, false, false, false, false, false, false];
            for c in n.chars() {
                seven[seven_segment_region_map[&c]] = true;
            }
            // println!("{:?}", seven_segment_region_map);
            // println!("{:?} {:?}", n, seven);
            number.push(get_number(seven).to_string());
        }
        let parsed_number = number.join("").parse::<i32>().unwrap();
        // println!("{}", parsed_number);
        s += parsed_number;
    }

    s
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
