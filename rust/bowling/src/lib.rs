#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

#[derive(Clone, Debug)]
struct Frame {
    rolls: Vec<u16>,
}

impl Frame {
    fn score(&self) -> u16 {
        self.rolls.iter().sum()
    }

    fn len(&self) -> usize {
        self.rolls.len()
    }
}

pub struct BowlingGame {
    frames: Vec<Frame>,
    frame: usize,
}

impl BowlingGame {
    pub fn new() -> Self {
        BowlingGame {
            frames: vec![Frame { rolls: vec![] }; 10],
            frame: 0,
        }
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if self.is_complete() {
            return Err(Error::GameComplete);
        }

        let frame = &mut self.frames[self.frame];

        let pins_left = if self.frame < 9 || frame.len() == 0 || frame.len() == 1 && frame.score() < 10 {
            10 - frame.score()
        } else if frame.len() == 1 || frame.len() == 2 && frame.score() < 20 {
            20 - frame.score()
        } else {
            30 - frame.score()
        };

        if pins > pins_left {
            return Err(Error::NotEnoughPinsLeft);
        }

        frame.rolls.push(pins);

        if self.frame < 9 && (frame.len() == 2 || frame.score() == 10)
        || frame.len() == 2 && frame.score() < 10
        || frame.len() == 3 {
            self.frame += 1;
        }

        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        match self.is_complete() {
            true => Some(
                self.frames.iter().enumerate().fold(0, |acc, (i, frame)| {
                    let mut score = frame.score();

                    if i < 9 && frame.score() == 10 {
                        score += self.frames[i + 1].rolls[0];
                        if frame.len() == 1 {
                            if self.frames[i + 1].len() == 1 {
                                score += self.frames[i + 2].rolls[0];
                            } else {
                                score += self.frames[i + 1].rolls[1];
                            }
                        }
                    }

                    acc + score
                })
            ),
            false => None,
        }
    }

    fn is_complete(&self) -> bool {
        self.frame == 10
    }
}
