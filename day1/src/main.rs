#[macro_use]
extern crate nom;
extern crate num;

mod parser;

use nom::IResult;
use num::{Complex, Zero};

use std::collections::HashSet;
use std::io::{self, Read};

use parser::road;

#[derive(Copy, Clone, Debug)]
enum Direction {
    Left,
    Right,
}

#[derive(Copy, Clone, Debug)]
pub struct Walk {
    direction: Direction,
    steps: i32,
}

fn rotate(number: Complex<i32>, direction: Direction) -> Complex<i32> {
    let direction = match direction {
        Direction::Left => -Complex::i(),
        Direction::Right => Complex::i(),
    };
    number * direction
}

fn taxicab_distance(Complex { re, im }: Complex<i32>) -> i32 {
    re.abs() + im.abs()
}

fn calculate_position(road: &[Walk]) -> Complex<i32> {
    road.iter()
        .fold((Complex::i(), Complex::zero()),
              |(direction, position), walk| {
                  let new_direction = rotate(direction, walk.direction);
                  (new_direction, position + new_direction * Complex::from(walk.steps))
              })
        .1
}

fn first_duplicate(road: &[Walk]) -> Option<Complex<i32>> {
    let mut direction = Complex::i();
    let mut position = Complex::zero();
    let mut reached = HashSet::new();
    reached.insert((0, 0));
    for walk in road {
        direction = rotate(direction, walk.direction);
        for _ in 0..walk.steps {
            position = position + direction;
            if !reached.insert((position.re, position.im)) {
                return Some(position);
            }
        }
    }
    None
}

fn main() {
    let input = io::stdin();
    let mut out = Vec::new();
    input.lock().read_to_end(&mut out).unwrap();
    match road(&out) {
        IResult::Done(_, road) => {
            println!("Distance: {}", taxicab_distance(calculate_position(&road)));
            let duplicate = first_duplicate(&road).map(taxicab_distance);
            match duplicate {
                Some(position) => println!("First duplicate location: {}", position),
                None => println!("No duplicates found"),
            }
        }
        IResult::Error(_) => panic!("Syntax error"),
        IResult::Incomplete(_) => panic!("Incomplete input"),
    }
}

#[test]
fn destination() {
    let tests: [(&[u8], i32); 3] = [(b"R2, L3", 5), (b"R2, R2, R2", 2), (b"R5, L5, R5, R3", 12)];
    for &(route, answer) in &tests {
        assert_eq!(taxicab_distance(calculate_position(&road(route).unwrap().1)),
                   answer);
    }
}

#[test]
fn duplicate() {
    assert_eq!(first_duplicate(&[]), None);
    assert_eq!(first_duplicate(&road(b"R8, R4, R4, R8").unwrap().1).map(taxicab_distance),
               Some(4));
}
