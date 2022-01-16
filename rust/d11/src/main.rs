use std::{
    collections::{HashSet, LinkedList},
    hash::Hash,
};

type Grid = [[i32; 10]; 10];

fn main() {
    let input = parse(include_str!("../input.txt"));
    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn parse(input: &str) -> Grid {
    let mut g: Grid = [[0; 10]; 10];
    for (row, line) in input.lines().enumerate() {
        for (col, c) in line.split("").filter(|v| v.len() > 0).enumerate() {
            if c.len() > 0 {
                g[row][col] = c.parse().unwrap();
            }
        }
    }

    g
}

fn part1(g: &Grid) -> i32 {
    let mut os = g.clone();
    let (rows, cols) = (os.len(), os[0].len());
    let mut flashes = 0;

    for _ in 0..100 {
        let mut flashed: HashSet<(usize, usize)> = HashSet::new();
        for row in 0..os.len() {
            for col in 0..os[row].len() {
                let mut stack = vec![(row, col)];
                while stack.len() > 0 {
                    let (crow, ccol) = stack.pop().unwrap();
                    os[crow][ccol] += 1;
                    if os[crow][ccol] >= 10 && !flashed.contains(&(crow, ccol)) {
                        flashed.insert((crow, ccol));
                        for n in get_neighbours((rows, cols), (crow, ccol)) {
                            stack.push(n);
                        }
                    }
                }
            }
        }
        for row in 0..os.len() {
            for col in 0..os[row].len() {
                if os[row][col] >= 10 {
                    os[row][col] = 0;
                    flashes += 1;
                }
            }
        }

        // for row in 0..os.len() {
        //     println!("{:?}", os[row]);
        // }
        // println!();
    }
    flashes
}

fn part2(g: &Grid) -> i32 {
    let mut os = g.clone();
    let (rows, cols) = (os.len(), os[0].len());
    let mut step = 0;
    loop {
        // same as above
        step += 1;
        let mut flashed: HashSet<(usize, usize)> = HashSet::new();
        for row in 0..os.len() {
            for col in 0..os[row].len() {
                let mut stack = vec![(row, col)];
                while stack.len() > 0 {
                    let (crow, ccol) = stack.pop().unwrap();
                    os[crow][ccol] += 1;
                    if os[crow][ccol] >= 10 && !flashed.contains(&(crow, ccol)) {
                        flashed.insert((crow, ccol));
                        for n in get_neighbours((rows, cols), (crow, ccol)) {
                            stack.push(n);
                        }
                    }
                }
            }
        }

        // flashed items should equal total number of elements
        // ie 100 (since it's a 10x10 grid)
        if flashed.len() == 100 {
            break;
        }
        for row in 0..os.len() {
            for col in 0..os[row].len() {
                if os[row][col] >= 10 {
                    os[row][col] = 0;
                }
            }
        }
    }
    step
}

fn get_neighbours(total: (usize, usize), current: (usize, usize)) -> Vec<(usize, usize)> {
    let (total_rows, total_cols) = total;
    let (row, col) = current;

    let mut ns = vec![];
    let dr = [1, -1, 0, 0, 1, -1, -1, 1];
    let dc = [0, 0, 1, -1, -1, 1, -1, 1];
    for i in 0..8 {
        let (new_row, new_col) = (row as i32 + dr[i], col as i32 + dc[i]);
        if new_col >= 0
            && new_row >= 0
            && (new_row as usize) < total_rows
            && (new_col as usize) < total_cols
        {
            ns.push((new_row as usize, new_col as usize));
        }
    }

    ns
}

#[test]
fn test_get_neighbours() {
    let got = get_neighbours((4, 4), (1, 1));
    assert_eq!(
        got,
        vec![
            (2, 1),
            (0, 1),
            (1, 2),
            (1, 0),
            (2, 0),
            (0, 2),
            (0, 0),
            (2, 2)
        ]
    );
}
