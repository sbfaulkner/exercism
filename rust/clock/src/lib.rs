use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    time: i32,
}

const MINUTES_PER_DAY: i32 = 24 * 60;

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Clock{
            time: (hours * 60 + minutes).rem_euclid(MINUTES_PER_DAY),
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock{
            time: (self.time + minutes).rem_euclid(MINUTES_PER_DAY),
        }
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.time / 60, self.time % 60)
    }
}
