use nom::digit;

use {Direction, Walk};

named!(pub road< Vec<Walk> >, terminated!(separated_list!(tag!(","), ws!(walk)), eof!()));
named!(walk< Walk >, do_parse!(
    direction: alt!(
        tag!("L") => {|_| Direction::Left} |
        tag!("R") => {|_| Direction::Right}) >>
    steps: map_res!(digit, |x| String::from_utf8_lossy(x).parse()) >>
    (Walk {
        direction: direction,
        steps: steps,
    })
));
