use crate::day::Day;
use serde_json::json;

pub struct Day7;

impl Day for Day7 {
    fn process( input: &str) -> String {

        let data = parse_input(input);

        let part1_result = solve_calibration_part1(&data);
        let part2_result = solve_calibration_part2(&data);

        let result = json!({
            "part_one": part1_result,
            "part_two": part2_result,
        });

        result.to_string()
    }
}

fn parse_input(input: &str) -> Vec<(i64, Vec<i64>)> {

    let mut result = Vec::new();

    for line in input.lines() {

        let line = line.trim();
        if line.is_empty() {
            continue;
        }

        let mut parts = line.split(':');
        let target_part = parts.next().unwrap().trim();
        let nums_part = parts.next().unwrap_or("").trim();

        let target: i64 = match target_part.parse() {
            Ok(t) => t,
            Err(_) => continue,
        };

        let numbers: Vec<i64> = nums_part
            .split_whitespace()
            .filter_map(|s| s.parse().ok())
            .collect();

        if !numbers.is_empty() {
            result.push((target, numbers));
        }
    }

    result
}

fn can_form_target(numbers: Vec<i64>, target: i64, allow_concat: bool) -> bool {

    fn recurse(nums: Vec<i64>, target: i64, allow_concat: bool) -> bool {

        if nums.len() == 1 {
            return nums[0] == target;
        }

        let a = nums[0];
        let b = nums[1];
        let rest = &nums[2..];

        // +
        {
            let mut new_nums = Vec::new();
            new_nums.push(a + b);
            new_nums.extend_from_slice(rest);
            if recurse(new_nums, target, allow_concat) {
                return true;
            }
        }

        // *
        {
            let mut new_nums = Vec::new();
            new_nums.push(a * b);
            new_nums.extend_from_slice(rest);
            if recurse(new_nums, target, allow_concat) {
                return true;
            }
        }

        // ||
        if allow_concat {
            let concat_str = format!("{}{}", a, b);
            if let Ok(concat_val) = concat_str.parse::<i64>() {
                let mut new_nums = Vec::new();
                new_nums.push(concat_val);
                new_nums.extend_from_slice(rest);
                if recurse(new_nums, target, allow_concat) {
                    return true;
                }
            }
        }

        false
    }

    recurse(numbers, target, allow_concat)
}

fn solve_calibration_part1(data: &[(i64, Vec<i64>)]) -> i64 {
    let mut total = 0;
    for (target, numbers) in data {
        if can_form_target(numbers.clone(), *target, false) {
            total += target;
        }
    }
    total
}

fn solve_calibration_part2(data: &[(i64, Vec<i64>)]) -> i64 {
    let mut total = 0;
    for (target, numbers) in data {
        if can_form_target(numbers.clone(), *target, true) {
            total += target;
        }
    }
    total
}

