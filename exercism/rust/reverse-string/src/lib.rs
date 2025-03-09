use unicode_reverse::reverse_grapheme_clusters_in_place;

pub fn reverse(input: &str) -> String {
    let mut reversed_str = String::from(input);

    reverse_grapheme_clusters_in_place(&mut reversed_str);

    reversed_str
}
