package wordy

const testVersion = 1

// Answer parses and evaluates a simple math word problem.
func Answer(input string) (answer int, ok bool) {
	e := evaluator{
		lex:  lex(input),
		eval: evaluateStart,
	}

	for e.eval != nil {
		e.eval = e.eval(&e)
	}

	if e.current.typ != tokenQuestionMark {
		return
	}

	return e.answer, true
}
