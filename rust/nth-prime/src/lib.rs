pub fn nth(n: u32) -> u32 {
    let mut count = 0;
    let mut candidate = 1;

    while count < n+1 {
        candidate += 1;
        if is_prime(candidate) {
            count += 1;
        }
    }

    candidate
}

fn is_prime(n: u32) -> bool {
    let divisor = (n as f64).sqrt() as u32;

    for d in 2..=divisor {
        if n % d == 0 {
            return false;
        }
    }

    true
}
