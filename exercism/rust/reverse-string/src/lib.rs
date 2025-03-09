pub fn reverse(input: &str) -> String {
    let mut reversed_str = String::with_capacity(input.len() * 2);

    for c in input.chars().rev() {
        reversed_str.push(c);
    }

    reversed_str
}
