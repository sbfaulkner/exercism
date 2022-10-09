#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

fn pop_args(stack: &mut Vec<i32>) -> Option<(i32, i32)> {
    let b = stack.pop()?;
    let a = stack.pop()?;
    Some((a, b))
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack: Vec<i32> = vec![];

    for input in inputs {
        match input {
            CalculatorInput::Value(n) => stack.push(*n),
            CalculatorInput::Add => {
                let (a, b) = pop_args(&mut stack)?;
                stack.push(a + b)
            },
            CalculatorInput::Subtract => {
                let (a, b) = pop_args(&mut stack)?;
                stack.push(a - b)
            },
            CalculatorInput::Multiply => {
                let (a, b) = pop_args(&mut stack)?;
                stack.push(a * b)
            },
            CalculatorInput::Divide => {
                let (a, b) = pop_args(&mut stack)?;
                stack.push(a / b)
            },
        }
    }

    match stack.len() {
        1 => stack.pop(),
        _ => None,
    }
}
