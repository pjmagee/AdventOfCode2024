
mod day;
use crate::day::Day;
mod day7;
use atty;
use std::io;
use std::io::Read;

fn main() {

    let mut input = String::new();

    if atty::is(atty::Stream::Stdin) {

        let data = r#"
        190: 10 19
        3267: 81 40 27
        83: 17 5
        156: 15 6
        7290: 6 8 6 15
        161011: 16 10 13
        192: 17 8 14
        21037: 9 7 18 13
        292: 11 6 16 20"#;
        input = String::from(data);

    } else {
        io::stdin()
            .read_to_string(&mut input)
            .expect("Failed to read from stdin");
    }

    let day_arg = std::env::args().nth(1);

    if day_arg.is_none() {
        println!("Usage: rs.exe <day> < <input>");
        return;
    }

    if input.len() == 0 {
        println!("Usage: rs.exe <day> < <input>");
        return;
    }

    let day = day_arg.unwrap().parse::<u32>();

    if day.is_err() {
        println!("Day must be a number");
        return;
    }

    let day = day.unwrap();

    let result = match day {
        7 => day7::Day7::process(&input),
        _ => format!("Day {} not implemented", day),
    };

    println!("{}", result);
}
