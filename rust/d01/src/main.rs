fn main() {
    let raw_input = include_str!("../input.txt");
    let nums: Vec<i32> = raw_input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect();

    println!("{}", part1(&nums));
    println!("{}", part2(&nums));
}

fn part1(nums: &Vec<i32>) -> i32 {
    let mut result = 0;
    for (idx, n) in nums.iter().enumerate() {
        if idx < nums.len() - 1 {
            if nums[idx + 1] > *n {
                result += 1;
            }
        }
    }
    result
}

fn part2(nums: &Vec<i32>) -> i32 {
    let mut result = 0;
    let mut prev = [nums[0], nums[1], nums[2]].iter().sum::<i32>();

    for i in 1..(nums.len() - 2) {
        let window: [i32; 3] = [nums[i], nums[i + 1], nums[i + 2]];
        let window_sum = window.iter().sum::<i32>();
        if window_sum > prev {
            result += 1;
        }
        prev = window_sum;
    }
    result
}
