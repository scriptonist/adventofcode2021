fn main() {
    let input = include_str!("../input.txt");
    let grid = parse(input);
    println!("{}", part1(&grid));
}

fn part1(grid: &Vec<Vec<i32>>) -> i32 {
    let rows = grid.len();
    let cols = grid[0].len();
    let mut sum = 0;
    for row in 0..rows {
        for col in 0..cols {
            let cur = grid[row][col];
            let ns = get_neighbours(grid, row, col);
            // println!("{} {:?}", cur, &ns);
            if ns.into_iter().all(|v| cur < v) {
                sum += cur + 1;
                // println!("low: {}", cur);
            }
        }
    }

    sum
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

fn get_neighbours(grid: &Vec<Vec<i32>>, row: usize, col: usize) -> Vec<i32> {
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
