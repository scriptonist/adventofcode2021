fn main() {
    let input = include_str!("../input.txt");
    println!("{}", part1(input));
    println!("{}", part2(input));
}

fn part1(input: &str) -> i32 {
    let mut score = 0;
    for line in input.lines() {
        let mut s: Vec<char> = vec![];
        for c in line.chars() {
            match c {
                '(' | '[' | '{' | '<' => s.push(c),
                ')' | ']' | '}' | '>' => {
                    if let Some(top_c) = s.pop() {
                        if get_closer(&top_c) != c {
                            score += get_score_p1(&c);
                        }
                    }
                }
                _ => unreachable!(),
            }
        }
    }
    score
}

fn part2(input: &str) -> i64 {
    let cleaned = clean(input);
    let mut scores: Vec<i64> = vec![];
    for line in cleaned {
        scores.push(get_line_score(line));
    }
    scores.sort();
    scores.iter().nth(scores.len() / 2).unwrap().to_owned()
}

fn get_line_score(line: &str) -> i64 {
    let mut score: i64 = 0;
    let mut s: Vec<char> = vec![];
    for c in line.chars() {
        match c {
            '(' | '[' | '{' | '<' => s.push(c),
            ')' | ']' | '}' | '>' => {
                let opening = s.pop().unwrap();
                if c != get_closer(&opening) {
                    s.push(opening);
                }
            }
            _ => unreachable!(),
        }
    }
    for c in s.iter().rev() {
        score = (score * 5) + get_score_p2(&get_closer(c)) as i64;
        // println!("{} {}", get_closer(c), score);
    }
    // println!("{:?}", s);
    // println!("{}", score);
    return score;
}

#[test]
fn test_get_line_score() {
    assert_eq!(get_line_score("<{([{{}}[<[[[<>{}]]]>[]]"), 294);
    assert_eq!(get_line_score("[({(<(())[]>[[{[]{<()<>>"), 288957);
    assert_eq!(get_line_score("[(()[<>])]({[<{<<[]>>("), 5566);
    assert_eq!(get_line_score("(((({<>}<{<{<>}{[]{[]{}"), 1480781);
    assert_eq!(get_line_score("{<[[]]>}<{[{[{[]{()[[[]"), 995444);
}

fn clean(input: &str) -> Vec<&str> {
    let mut cleaned = vec![];
    'main: for line in input.lines() {
        let mut s: Vec<char> = vec![];
        for c in line.chars() {
            match c {
                '(' | '[' | '{' | '<' => s.push(c),
                ')' | ']' | '}' | '>' => {
                    if let Some(top_c) = s.pop() {
                        if get_closer(&top_c) != c {
                            continue 'main;
                        }
                    }
                }
                _ => unreachable!(),
            }
        }
        cleaned.push(line)
    }
    cleaned
}

#[test]
fn test_clean() {
    let input = include_str!("../test.txt");
    assert_eq!(clean(input).len(), 5);
}

fn get_score_p1(c: &char) -> i32 {
    match c {
        ')' => 3,
        ']' => 57,
        '}' => 1197,
        '>' => 25137,
        _ => unreachable!(),
    }
}

fn get_score_p2(c: &char) -> i32 {
    match c {
        ')' => 1,
        ']' => 2,
        '}' => 3,
        '>' => 4,
        _ => unreachable!(),
    }
}

fn get_closer(c: &char) -> char {
    match c {
        '(' => ')',
        '[' => ']',
        '{' => '}',
        '<' => '>',
        _ => unreachable!(),
    }
}
