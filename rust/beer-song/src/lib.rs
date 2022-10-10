use std::fmt;

struct Bottles(u8);

impl fmt::Display for Bottles {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match *self {
            Bottles(0) => write!(f, "no more bottles"),
            Bottles(1) => write!(f, "1 bottle"),
            Bottles(n) => write!(f, "{n} bottles"),
        }
    }
}

struct Verse(u8);

impl fmt::Display for Verse {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let args = match *self {
            Verse(0) => (Bottles(0), "Go to the store and buy some more", Bottles(99)),
            Verse(1) => (Bottles(1), "Take it down and pass it around", Bottles(0)),
            Verse(n) => (Bottles(n), "Take one down and pass it around", Bottles(n - 1)),
        };

        write!(f, "{0} of beer on the wall, {0} of beer.\n{1}, {2} of beer on the wall.\n", args.0, args.1, args.2)
    }
}

fn capitalize(s: &str) -> String {
    s[0..1].to_uppercase() + &s[1..]
}

pub fn verse(n: u32) -> String {
    capitalize(&Verse(n as u8).to_string())
}

pub fn sing(start: u32, end: u32) -> String {
    (end..=start).rev().map(|n| verse(n)).collect::<Vec<String>>().join("\n")
}
