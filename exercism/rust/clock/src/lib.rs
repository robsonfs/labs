use std::fmt;

type Integer = i32;

trait EuclideanMod {
    fn euclidean_mod(self, rhs: Self) -> Self;
}

impl EuclideanMod for Integer {
    fn euclidean_mod(self, rhs: Self) -> Self {
        ((self % rhs) + rhs) % rhs
    }
}


pub struct Clock {
    hours: Integer,
    minutes: Integer,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}

impl Clock {
    pub fn new(hours: Integer, minutes: Integer) -> Self {
        let mut hours = if hours < 0 {
            24 - hours.abs().euclidean_mod(24)
        } else {
            hours.euclidean_mod(24)
        };

        let minutes = if minutes == minutes.euclidean_mod(60) {
            minutes
        } else if minutes < 0 {
            hours -= 1;
            minutes.euclidean_mod(60)
        } else {
            hours += minutes / 60;
            minutes.euclidean_mod(60)
        };

        Self {
            hours: hours.euclidean_mod(24),
            minutes: minutes.euclidean_mod(60),
        }
    }

    pub fn add_minutes(&mut self, minutes: Integer) -> Self {
        todo!("Add {minutes} minutes to existing Clock time");
    }
}
