fn bottles(n: u32) -> String {
    match n {
        0 => "no more bottles".to_string(),
        1 => "1 bottle".to_string(),
        _ => format!("{} bottles", n),
    }
}

fn take_down(n: u32) -> String {
    match n {
        0 => "Go to the store and buy some more".to_string(),
        1 => "Take it down and pass it around".to_string(),
        _ => "Take one down and pass it around".to_string(),
    }
}

fn capitalize(s: &str) -> String {
    s[0..1].to_uppercase() + &s[1..]
}

pub fn verse(n: u32) -> String {
    capitalize(
        format!(
            "{0} of beer on the wall, {0} of beer.\n{1}, {2} of beer on the wall.\n",
            bottles(n),
            take_down(n),
            bottles((n + 99) % 100),
        ).as_str()
    )
}

pub fn sing(start: u32, end: u32) -> String {
    (end..=start).rev().map(|n| verse(n)).collect::<Vec<String>>().join("\n")
}
