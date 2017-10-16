package forth

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 1

var (
	errDivideByZero         = errors.New("divide by zero")
	errEmptyStack           = errors.New("empty stack")
	errCannotRedefineNumber = errors.New("cannot redefine number")
)

type evalFn func(*evaluator, []string) error

type entry struct {
	fn     evalFn
	params []string
}

type evaluator struct {
	words chan string
	data  []int
	dict  map[string]entry
}

func newEvaluator(input []string) *evaluator {
	e := evaluator{
		words: make(chan string, 2),
		data:  []int{},
		dict: map[string]entry{
			":":    {fn: define},
			"+":    {fn: add},
			"-":    {fn: subtract},
			"*":    {fn: multiply},
			"/":    {fn: divide},
			"DUP":  {fn: duplicate},
			"DROP": {fn: drop},
			"SWAP": {fn: swap},
			"OVER": {fn: over},
		},
	}

	go e.parse(input)

	return &e
}

func (e *evaluator) parse(input []string) {
	isSeparator := func(r rune) bool { return unicode.IsSpace(r) || unicode.IsControl(r) }

	for _, line := range input {
		for _, word := range strings.FieldsFunc(line, isSeparator) {
			e.emit(word)
		}
	}

	close(e.words)
}

func (e *evaluator) peek() (int, error) {
	if len(e.data) < 1 {
		return 0, errEmptyStack
	}

	return e.data[len(e.data)-1], nil
}

func (e *evaluator) pop() (int, error) {
	i, err := e.peek()
	if err != nil {
		return 0, err
	}

	e.data = e.data[:len(e.data)-1]

	return i, nil
}

func (e *evaluator) pop2() (int, int, error) {
	var i, j int
	var err error

	j, err = e.pop()
	if err != nil {
		return 0, 0, err
	}

	i, err = e.pop()
	if err != nil {
		return 0, 0, err
	}

	return i, j, nil
}

func (e *evaluator) push(i int) {
	e.data = append(e.data, i)
}

func (e *evaluator) next() string {
	return <-e.words
}

func (e *evaluator) emit(word string) {
	e.words <- strings.ToUpper(word)
}

func define(e *evaluator, _ []string) error {
	name := e.next()
	_, err := strconv.ParseInt(name, 10, 0)
	if err == nil {
		return errCannotRedefineNumber
	}

	code := []string{}

	for {
		word := e.next()
		if word == ";" {
			break
		}

		code = append(code, word)
	}

	e.dict[name] = entry{fn: execute, params: code}

	return nil
}

func execute(e *evaluator, code []string) error {
	for _, word := range code {
		if err := e.evaluate(word); err != nil {
			return err
		}
	}

	return nil
}

func (e *evaluator) performBinaryOperation(op func(i, j int) (int, error)) error {
	i, j, err := e.pop2()
	if err != nil {
		return err
	}

	k, err := op(i, j)
	if err != nil {
		return err
	}

	e.push(k)

	return nil
}

func add(e *evaluator, _ []string) error {
	return e.performBinaryOperation(func(i, j int) (int, error) { return i + j, nil })
}

func subtract(e *evaluator, _ []string) error {
	return e.performBinaryOperation(func(i, j int) (int, error) { return i - j, nil })
}

func multiply(e *evaluator, _ []string) error {
	return e.performBinaryOperation(func(i, j int) (int, error) { return i * j, nil })
}

func divide(e *evaluator, _ []string) error {
	return e.performBinaryOperation(func(i, j int) (int, error) {
		if j == 0 {
			return 0, errDivideByZero
		}
		return i / j, nil
	})
}

func duplicate(e *evaluator, _ []string) error {
	i, err := e.peek()
	if err != nil {
		return err
	}

	e.push(i)

	return nil
}

func drop(e *evaluator, _ []string) error {
	_, err := e.pop()
	if err != nil {
		return err
	}

	return nil
}

func swap(e *evaluator, _ []string) error {
	i, j, err := e.pop2()
	if err != nil {
		return err
	}

	e.push(j)
	e.push(i)

	return nil
}

func over(e *evaluator, _ []string) error {
	i, err := e.pop()
	if err != nil {
		return err
	}

	j, err := e.peek()
	if err != nil {
		return err
	}

	e.push(i)
	e.push(j)

	return nil
}

func (e *evaluator) evaluate(word string) error {
	if entry, ok := e.dict[word]; ok {
		err := entry.fn(e, entry.params)
		if err != nil {
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

func (e *evaluator) run() error {
	for {
		word := e.next()
		if word == "" {
			break
		}

		if err := e.evaluate(word); err != nil {
			return err
		}
	}

	return nil
}

// Forth evaluates a slice of input strings and returns the resulting stack as a slice of ints
func Forth(input []string) ([]int, error) {
	e := newEvaluator(input)

	err := e.run()
	if err != nil {
		return []int{}, err
	}

	return e.data, nil
}
