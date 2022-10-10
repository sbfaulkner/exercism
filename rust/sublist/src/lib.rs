#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> Comparison {
    if first_list.eq(second_list) {
        Comparison::Equal
    } else if contained(first_list, second_list) {
        Comparison::Sublist
    } else if contained(second_list, first_list) {
        Comparison::Superlist
    } else {
        Comparison::Unequal
    }
}

fn contained<T: PartialEq>(list: &[T], by: &[T]) -> bool {
    list.is_empty() || by.windows(list.len()).any(|window| window.eq(list))
}
