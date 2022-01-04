use std::collections::{HashMap, HashSet};

fn main() {
    let input = include_str!("../input.txt");
    let (numbers, boards) = parse_input(input);
    println!("{}", part2(numbers, boards));
}

fn parse_input(input: &str) -> (Vec<i32>, Vec<Board>) {
    let mut lines = input.lines();
    let numbers: Vec<i32> = lines
        .next()
        .unwrap()
        .split(",")
        .map(|v| v.parse::<i32>().unwrap())
        .collect();

    let mut boards: Vec<Board> = vec![];
    let mut current_board = Board::new();
    let mut row: usize = 0;
    lines.next();
    for line in lines {
        if line.is_empty() {
            boards.push(current_board);
            current_board = Board::new();
            row = 0;
        } else {
            line.split(" ")
                .filter(|v| !v.is_empty())
                .map(|ns| ns.parse::<i32>().unwrap())
                .enumerate()
                .for_each(|(idx, n)| {
                    current_board
                        .items
                        .entry(n)
                        .or_insert(Vec::new())
                        .push((row, idx))
                });
            row += 1;
        }
    }
    boards.push(current_board);

    (numbers, boards)
}

#[derive(Debug)]
struct Board {
    items: HashMap<i32, Vec<(usize, usize)>>,
    drawed: HashSet<(usize, usize)>,
}
impl Board {
    fn new() -> Self {
        Board {
            items: HashMap::new(),
            drawed: HashSet::new(),
        }
    }
    fn score(&self) -> i32 {
        let mut score = 0;
        for n in self.items.keys() {
            for (row, col) in self.items.get(n).unwrap() {
                if self.drawed.get(&(row.to_owned(), col.to_owned())).is_none() {
                    score += n;
                }
            }
        }

        score
    }
    fn has_won(&self) -> bool {
        // check colums
        for i in 0usize..5 {
            let mut found = true;
            for j in 0usize..5 {
                if self.drawed.get(&(j, i)).is_none() {
                    found = false;
                }
            }
            if found {
                return true;
            }
        }
        // check rows
        for i in 0usize..5 {
            let mut found = true;
            for j in 0usize..5 {
                if self.drawed.get(&(i, j)).is_none() {
                    found = false;
                }
            }
            if found {
                return true;
            }
        }

        false
    }
}

fn part2(numbers: Vec<i32>, mut boards: Vec<Board>) -> i32 {
    /*
    Board {
        Items n: idxs
        Drawed set of indices

        complete() -> bool
    }
    // Boards
    N1 -> 0,0 0,1
    N2 -> 1,1 1,2

    // found
    hashset -> 0,0 - 0,1 - 0,2
    */
    let mut score: i32 = 0;
    let total_boards = boards.len();
    let mut won_boards = 0;
    'mainloop: for n in numbers {
        for board in &mut boards {
            if !board.has_won() {
                if let Some(vals) = board.items.get(&n) {
                    vals.into_iter().for_each(|(row, col)| {
                        board.drawed.insert((row.to_owned(), col.to_owned()));
                    });
                }
                if board.has_won() {
                    won_boards += 1
                }
                if won_boards == total_boards {
                    println!("{} {}", board.score(), n);
                    score = board.score() * n;
                    break 'mainloop;
                }
            }
        }
    }

    score
}
