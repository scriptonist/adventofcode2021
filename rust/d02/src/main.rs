fn main() {
    let course = include_str!("../input.txt")
        .lines()
        .map(|action| action.split(" ").take(2).collect::<Vec<&str>>())
        .into_iter()
        .map(|item| {
            let (direction, val) = (item[0], item[1].parse::<i32>().unwrap());
            (direction, val)
        })
        .collect::<Vec<(&str, i32)>>();

    println!("{}", part1(&course));
    println!("{}", part2(&course));
}

fn part1(course: &Vec<(&str, i32)>) -> i32 {
    let (mut horizontal, mut depth) = (0, 0);
    for (direction, val) in course.iter() {
        match *direction {
            "forward" => horizontal += val,
            "down" => depth += val,
            "up" => depth -= val,
            _ => unreachable!(),
        }
    }

    horizontal * depth
}

fn part2(course: &Vec<(&str, i32)>) -> i32 {
    let (mut horizontal, mut depth, mut aim) = (0, 0, 0);
    for (direction, val) in course.iter() {
        match *direction {
            "forward" => {
                horizontal += val;
                depth += aim * val;
            }
            "down" => aim += val,
            "up" => aim -= val,
            _ => unreachable!(),
        }
    }

    horizontal * depth
}
