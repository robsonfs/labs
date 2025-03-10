use time::{ext::NumericalDuration};
use time::PrimitiveDateTime as DateTime;

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    match start.checked_add(1_000_000_000.seconds()) {
        Some(date) => date,
        None => panic!("Overflow error")
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use time::{PrimitiveDateTime};

    #[test]
    #[should_panic(expected = "Overflow error")]
    fn test_after_overflow_panics() {
        // Using the maximum possible date for PrimitiveDateTime,
        // adding 1 billion seconds should overflow and panic.
        let start = PrimitiveDateTime::MAX;
        // This call is expected to panic with "Overflow error".
        let _ = after(start);
    }
}
