package wordy

import (
	"strconv"
)

type evalFn func(*evaluator) evalFn

type evaluator struct {
	lex     *lexer
	current token
	answer  int
	eval    evalFn
}

func (e *evaluator) getToken() token {
	e.current = e.lex.getToken()
	return e.current
}

func (e *evaluator) operand() int {
	o, _ := strconv.Atoi(e.current.val)
	return o
}

func (e *evaluator) acceptWord(valid ...string) bool {
	t := e.getToken()

	if t.typ != tokenWord {
		return false
	}

	if len(valid) == 0 {
		return true
	}

	for _, v := range valid {
		if t.val == v {
			return true
		}
	}

	return false
}

func (e *evaluator) acceptNumber() bool {
	t := e.getToken()
	return t.typ == tokenNumber
}

func evaluateStart(e *evaluator) evalFn {
	if !e.acceptWord("what") {
		return nil
	}

	if !e.acceptWord("is") {
		return nil
	}

	return evaluateProblem
}

func evaluateProblem(e *evaluator) evalFn {
	if !e.acceptNumber() {
		return nil
	}

	e.answer = e.operand()

	return evaluateOperator
}

func evaluateOperator(e *evaluator) evalFn {
	if !e.acceptWord() {
		return nil
	}

	switch e.current.val {
	case "plus":
		return evaluatePlus
	case "minus":
		return evaluateMinus
	case "multiplied":
		return evaluateMultiplied
	case "divided":
		return evaluateDivided
	default:
		return nil
	}
}

func evaluatePlus(e *evaluator) evalFn {
	if !e.acceptNumber() {
		return nil
	}

	e.answer += e.operand()

	return evaluateOperator
}

func evaluateMinus(e *evaluator) evalFn {
	if !e.acceptNumber() {
		return nil
	}

	e.answer -= e.operand()

	return evaluateOperator
}

func evaluateMultiplied(e *evaluator) evalFn {
	if !e.acceptWord("by") {
		return nil
	}

	if !e.acceptNumber() {
		return nil
	}

	e.answer *= e.operand()

	return evaluateOperator
}

func evaluateDivided(e *evaluator) evalFn {
	if !e.acceptWord("by") {
		return nil
	}

	if !e.acceptNumber() {
		return nil
	}

	e.answer /= e.operand()

	return evaluateOperator
}
