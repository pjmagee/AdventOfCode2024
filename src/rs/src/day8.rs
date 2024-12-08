use crate::day::Day;
use serde_json::json;
use std::collections::HashMap;
use std::collections::HashSet;

pub struct Day8;

impl Day for Day8 {
    fn process( input: &str) -> String {
        json!({ "part_one": part_one(input).len(), "part_two": part_two(input).len() }).to_string()
    }
}

fn part_one(input: &str) -> HashSet<(isize, isize)> {

    let lines: Vec<String> = input.lines().map(|l| l.to_string()).collect();

    let rows = lines.len();
    let cols = if rows > 0 { lines[0].len() } else { 0 };

    // Store antenna positions by frequency
    let mut antennas_by_freq: HashMap<char, Vec<(isize, isize)>> = HashMap::new();

    for (r, line) in lines.iter().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if ch != '.' {
                // Record antenna position
                antennas_by_freq.entry(ch).or_default().push((r as isize, c as isize));
            }
        }
    }

    let mut antinodes: HashSet<(isize, isize)> = HashSet::new();

    // For each frequency, for each pair of antennas, find two antinodes
    for (_freq, positions) in antennas_by_freq.iter() {
        // For each pair of antennas
        for i in 0..positions.len() {
            for j in i+1..positions.len() {
                let (x_a, y_a) = positions[i];
                let (x_b, y_b) = positions[j];

                // Compute the two antinodes:
                // N1 = 2B - A
                let n1_x = 2 * x_b - x_a;
                let n1_y = 2 * y_b - y_a;
                // N2 = 2A - B
                let n2_x = 2 * x_a - x_b;
                let n2_y = 2 * y_a - y_b;

                // Check boundaries and insert if valid
                if n1_x >= 0 && n1_x < rows as isize && n1_y >= 0 && n1_y < cols as isize {
                    antinodes.insert((n1_x, n1_y));
                }

                if n2_x >= 0 && n2_x < rows as isize && n2_y >= 0 && n2_y < cols as isize {
                    antinodes.insert((n2_x, n2_y));
                }
            }
        }
    }

    antinodes
}

fn part_two(input: &str) -> HashSet<(isize, isize)> {

    let lines: Vec<String> = input.lines().map(|l| l.to_string()).collect();

    let rows = lines.len();
    let cols = if rows > 0 { lines[0].len() } else { 0 };

    // Map frequency -> vector of (row, col)
    let mut antennas_by_freq: HashMap<char, Vec<(isize, isize)>> = HashMap::new();
    for (r, line) in lines.iter().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if ch != '.' {
                antennas_by_freq.entry(ch).or_default().push((r as isize, c as isize));
            }
        }
    }

    let mut antinodes: HashSet<(isize, isize)> = HashSet::new();

    // For each frequency, find all lines
    for (_, positions) in antennas_by_freq.iter() {
        // If there's only one antenna for this frequency, it can't form a line or add antinodes
        // except itself is not considered since no line formed with another antenna of same frequency.
        if positions.len() == 1 {
            continue;
        }

        // A line is defined by direction and a line constant:
        // direction: (dx, dy) reduced by gcd and oriented consistently
        // normal: (-dy, dx), line constant d = normal.x*x1 + normal.y*y1
        let mut seen_lines: HashSet<(isize, isize, isize)> = HashSet::new();

        // Insert all antennas first as they are definitely antinodes if there's more than one antenna of same freq
        for &(x, y) in positions {
            antinodes.insert((x, y));
        }

        fn gcd(a: isize, b: isize) -> isize {
            if b == 0 {
                a.abs()
            } else {
                gcd(b, a % b)
            }
        }

        for i in 0..positions.len() {
            for j in i+1..positions.len() {
                let (x1, y1) = positions[i];
                let (x2, y2) = positions[j];

                let mut dx = x2 - x1;
                let mut dy = y2 - y1;

                // Reduce by gcd
                let g = gcd(dx, dy);
                dx /= g;
                dy /= g;

                // Ensure consistent orientation
                if dx < 0 || (dx == 0 && dy < 0) {
                    dx = -dx;
                    dy = -dy;
                }

                // Normal vector to (dx, dy) is (-dy, dx)
                let nx = -dy;
                let ny = dx;

                // line constant d = nx*x1 + ny*y1
                let d = nx * x1 + ny * y1;

                let line_id = (nx, ny, d);
                if seen_lines.contains(&line_id) {
                    // Line already processed
                    continue;
                }
                seen_lines.insert(line_id);

                // We have a unique line for this frequency. Now generate all points along this line.
                // Start from one known antenna on the line, say (x1, y1).
                // Move in the direction (dx, dy) and (-dx, -dy).
                //
                // We'll expand until out of bounds.

                // Forward direction
                let mut tx = x1;
                let mut ty = y1;
                // Move forward
                loop {
                    tx += dx;
                    ty += dy;
                    if tx < 0 || tx >= rows as isize || ty < 0 || ty >= cols as isize {
                        break;
                    }
                    antinodes.insert((tx, ty));
                }

                // Backward direction
                let mut tx = x1;
                let mut ty = y1;
                loop {
                    tx -= dx;
                    ty -= dy;
                    if tx < 0 || tx >= rows as isize || ty < 0 || ty >= cols as isize {
                        break;
                    }
                    antinodes.insert((tx, ty));
                }
            }
        }
    }

    antinodes
}

