fn main() {
    let input = include_str!("../input.txt").lines();
    println!("{}", part1(input));
}

fn part1(report: std::str::Lines) -> i32 {
    let report_vals: Vec<Vec<char>> = report.map(|s| s.chars().collect::<Vec<char>>()).collect();
    let bits_count = report_vals[0].len();
    let (mut epsilon_rate, mut gamma_rate) = (0, 0);
    let two: i32 = 2;

    for i in 0..bits_count {
        let (mut zeros_count, mut ones_count) = (0, 0);
        for val in &report_vals {
            match val[i] {
                '0' => zeros_count += 1,
                '1' => ones_count += 1,
                _ => unreachable!(),
            }
        }
        if zeros_count > ones_count {
            epsilon_rate += two.pow((bits_count - i - 1) as u32);
        } else {
            gamma_rate += two.pow((bits_count - i - 1) as u32);
        }
    }
    gamma_rate * epsilon_rate
}
