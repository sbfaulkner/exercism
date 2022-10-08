// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

fn success_rate(speed: u8) -> f64 {
    if speed < 5 {
        1.0
    } else if speed < 9 {
        0.9
    } else {
        0.77
    }
}

pub fn production_rate_per_hour(speed: u8) -> f64 {
    speed as f64 * 221.0 * success_rate(speed)
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}
