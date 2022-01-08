use std::{
    collections::{HashSet, LinkedList},
    vec,
};

fn main() {
    let input = include_str!("../input.txt");
    let grid = parse(input);
    println!("{}", part1(&grid));
    println!("{}", part2(&grid));
}

fn part1(grid: &Vec<Vec<i32>>) -> i32 {
    let rows = grid.len();
    let cols = grid[0].len();
    let mut sum = 0;
    for row in 0..rows {
        for col in 0..cols {
            let cur = grid[row][col];
            let ns = get_neighbours_part1(grid, row, col);
            // println!("{} {:?}", cur, &ns);
            if ns.into_iter().all(|v| cur < v) {
                sum += cur + 1;
                // println!("low: {}", cur);
            }
        }
    }

    sum
}

fn part2(grid: &Vec<Vec<i32>>) -> i32 {
    let rows = grid.len();
    let cols = grid[0].len();
    let mut basin_sizes: Vec<i32> = vec![];
    for row in 0..rows {
        for col in 0..cols {
            if grid[row][col] != 9 {
                let size = get_size_of_increasing_path(grid, row, col);
                basin_sizes.push(size);
            }
        }
    }
    basin_sizes.sort();
    basin_sizes.iter().rev().take(3).product()
}

fn get_size_of_increasing_path(grid: &Vec<Vec<i32>>, row: usize, col: usize) -> i32 {
    let mut size = 0;
    let mut q: LinkedList<(usize, usize)> = LinkedList::new();
    let mut done: HashSet<(usize, usize)> = HashSet::new();
    q.push_back((row, col));
    done.insert((row, col));
    while q.len() > 0 {
        let (cur_row, cur_col) = q.pop_front().unwrap();
        // print!("{:?} ", grid[cur_row][cur_col]);
        size += 1;
        for (nrow, ncol) in get_neighbours_part2(grid, cur_row.to_owned(), cur_col.to_owned()) {
            if !done.contains(&(nrow, ncol)) {
                done.insert((nrow, ncol));
                q.push_back((nrow, ncol));
            }
        }
    }
    // println!(" size {}", size);

    size
}

fn parse(input_str: &str) -> Vec<Vec<i32>> {
    let mut grid: Vec<Vec<i32>> = vec![];
    for line in input_str.lines() {
        let row: Vec<i32> = line
            .split("")
            .filter(|v| v.len() > 0)
            .map(|v| v.parse::<i32>().unwrap())
            .collect();
        grid.push(row)
    }
    // for g in &grid {
    //     println!("{:?}", g);
    // }
    grid
}

fn get_neighbours_part1(grid: &Vec<Vec<i32>>, row: usize, col: usize) -> Vec<i32> {
    let rows = grid.len() as i32;
    let cols = grid[0].len() as i32;
    let mut neighbours: Vec<i32> = vec![];
    let dr: [i32; 4] = [1, -1, 0, 0];
    let dc: [i32; 4] = [0, 0, 1, -1];
    for i in 0..4 {
        let new_row = row as i32 + dr[i];
        let new_col = col as i32 + dc[i];
        if new_row >= 0 && new_col >= 0 && new_row < rows && new_col < cols {
            neighbours.push(grid[new_row as usize][new_col as usize]);
        }
    }

    neighbours
}

fn get_neighbours_part2(grid: &Vec<Vec<i32>>, row: usize, col: usize) -> Vec<(usize, usize)> {
    let rows = grid.len() as i32;
    let cols = grid[0].len() as i32;
    let mut neighbours: Vec<(usize, usize)> = vec![];
    let dr: [i32; 4] = [1, -1, 0, 0];
    let dc: [i32; 4] = [0, 0, 1, -1];
    for i in 0..4 {
        let new_row = row as i32 + dr[i];
        let new_col = col as i32 + dc[i];
        if new_row >= 0 && new_col >= 0 && new_row < rows && new_col < cols {
            if grid[new_row as usize][new_col as usize] >= grid[row][col]
                && grid[new_row as usize][new_col as usize] != 9
            {
                neighbours.push((new_row as usize, new_col as usize));
            }
        }
    }

    neighbours
}

#[test]
fn test_get_neighbours() {
    let grid = vec![
        vec![2, 1, 9, 9, 9, 4, 3, 2, 1, 0],
        vec![3, 9, 8, 7, 8, 9, 4, 9, 2, 1],
    ];
    assert_eq!(get_neighbours_part2(&grid, 0, 1), vec![(0, 0)]);
    assert_eq!(get_neighbours_part2(&grid, 0, 0), vec![(1, 0)]);
}
