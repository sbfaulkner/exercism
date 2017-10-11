package forth

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 1

type evalFn func(*evaluator) error

type evaluator struct {
	data []int
	dict map[string]evalFn
}

func newEvaluator() *evaluator {
	return &evaluator{
		data: []int{},
		dict: map[string]evalFn{
			"+":    add,
			"-":    subtract,
			"*":    multiply,
			"/":    divide,
			"DUP":  duplicate,
			"DROP": drop,
			"SWAP": swap,
			"OVER": over,
		},
	}
}

func (e *evaluator) peek() (value int, err error) {
	if len(e.data) < 1 {
		err = errors.New("empty stack")
		return
	}

	return e.data[len(e.data)-1], nil
}

func (e *evaluator) pop() (value int, err error) {
	value, err = e.peek()
	if err != nil {
		return
	}

	e.data = e.data[:len(e.data)-1]

	return
}

func (e *evaluator) push(value int) {
	e.data = append(e.data, value)
}

func (e *evaluator) perform(f func(x, y int) ([]int, error)) (err error) {
	var x, y int
	var result []int

	y, err = e.pop()
	if err != nil {
		return
	}

	x, err = e.pop()
	if err != nil {
		return
	}

	result, err = f(x, y)
	if err != nil {
		return
	}

	for _, r := range result {
		e.push(r)
	}

	return
}

func add(e *evaluator) error {
	return e.perform(func(x, y int) ([]int, error) { return []int{x + y}, nil })
}

func subtract(e *evaluator) error {
	return e.perform(func(x, y int) ([]int, error) { return []int{x - y}, nil })
}

func multiply(e *evaluator) error {
	return e.perform(func(x, y int) ([]int, error) { return []int{x * y}, nil })
}

func divide(e *evaluator) error {
	return e.perform(func(x, y int) ([]int, error) {
		if y == 0 {
			return []int{}, errors.New("divide by zero")
		}

		return []int{x / y}, nil
	})
}

func duplicate(e *evaluator) error {
	value, err := e.peek()
	if err != nil {
		return err
	}

	e.push(value)

	return nil
}

func drop(e *evaluator) error {
	_, err := e.pop()
	if err != nil {
		return err
	}

	return nil
}

func swap(e *evaluator) error {
	return e.perform(func(x, y int) ([]int, error) {
		return []int{y, x}, nil
	})
}

func over(e *evaluator) error {
	return e.perform(func(x, y int) ([]int, error) {
		return []int{x, y, x}, nil
	})
}

func (e *evaluator) evaluate(word string) error {
	if fn := e.dict[strings.ToUpper(word)]; fn != nil {
		if err := fn(e); err != nil {
			return err
		}
	} else {
		value, err := strconv.ParseInt(word, 10, 0)
		if err != nil {
			return err
		}

		e.push(int(value))
	}

	return nil
}

// Forth evaluates a slice of input strings and returns the resulting stack as a slice of ints
func Forth(input []string) ([]int, error) {
	isSeparator := func(r rune) bool { return unicode.IsSpace(r) || unicode.IsControl(r) }

	e := newEvaluator()

	for _, line := range input {
		for _, word := range strings.FieldsFunc(line, isSeparator) {
			if err := e.evaluate(word); err != nil {
				return []int{}, err
			}
		}
	}

	return e.data, nil
}
