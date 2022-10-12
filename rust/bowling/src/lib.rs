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
}

pub struct BowlingGame {
    frames: Vec<Frame>,
    frame: u16,
}

impl BowlingGame {
    pub fn new() -> Self {
        BowlingGame {
            frames: vec![Frame { rolls: vec![] }; 10],
            frame: 0,
        }
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        println!("Roll {} pins in frame {}", pins, self.frame);

        if self.is_complete() {
            println!("Game is complete");
            return Err(Error::GameComplete);
        }

        let frame = &mut self.frames[self.frame as usize];

        let remaining = match frame.rolls.len() {
            2 if self.frame == 9 && frame.rolls[0] == 10 && frame.rolls[1] != 10 => 10 - frame.rolls[1],
            2 if self.frame == 9 => 10,
            1 if frame.rolls[0] == 10 => 10,
            _ => 10 - frame.score(),
        };

        if pins > remaining {
            println!("Not enough pins left");
            return Err(Error::NotEnoughPinsLeft);
        }

        frame.rolls.push(pins);

        if self.frame == 9 && (frame.rolls.len() == 3 || frame.rolls.len() == 2 && frame.score() < 10)
        || self.frame < 9 && (frame.rolls.len() == 2 || frame.score() == 10) {
            self.frame += 1;
        }

        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        println!("Score {:?}", self.frames);

        match self.is_complete() {
            true => Some(
                self.frames.iter().enumerate().fold(0, |acc, (i, frame)| {
                    let mut score = frame.score();

                    // println!("Frame {}: {:?} (score={})", i, frame, score);

                    // if frame.rolls.len() == 1 {
                    //     score += self.frames[i + 1].rolls[0];
                    //     score += self.frames[i + 1].rolls[1];
                    // } else
                    if i < 9 && frame.score() == 10 && frame.rolls.len() < 3 {
                        score += self.frames[i + 1].rolls[0];
                        if frame.rolls.len() < 2 {
                            if self.frames[i+1].rolls.len() > 1 {
                                score += self.frames[i + 1].rolls[1];
                            } else {
                                score += self.frames[i + 2].rolls[0];
                            }
                        }
                    }

                    // println!("Frame {}: {:?} (score={})", i, frame, score);

                    acc + score
                })
            ),
            false => None,
        }
    }

    fn is_complete(&self) -> bool {
        println!("Complete (frame={})", self.frame);
        self.frame == 10
    }
}
