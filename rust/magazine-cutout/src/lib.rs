use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut words: HashMap<&str, i32> = HashMap::new();

    for word in magazine {
        words.entry(word).and_modify(|count| *count += 1).or_insert(1);
    }

    for word in note {
        if *words.entry(word).and_modify(|count| *count -= 1).or_insert(-1) < 0 {
            return false;
        }
    }

    true
}
