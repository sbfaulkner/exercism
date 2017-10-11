package forth

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 1

type stack []int

func (s *stack) peek() (value int, err error) {
	if len(*s) < 1 {
		err = errors.New("empty stack")
		return
	}

	return (*s)[len(*s)-1], nil
}

func (s *stack) pop() (value int, err error) {
	value, err = s.peek()
	if err != nil {
		return
	}

	*s = (*s)[:len(*s)-1]

	return
}

func (s *stack) push(value int) {
	*s = append(*s, value)
}

func (s *stack) perform(f func(x, y int) ([]int, error)) (err error) {
	var x, y int
	var result []int

	y, err = s.pop()
	if err != nil {
		return
	}

	x, err = s.pop()
	if err != nil {
		return
	}

	result, err = f(x, y)
	if err != nil {
		return
	}

	for _, r := range result {
		s.push(r)
	}

	return
}

func add(s *stack) error {
	return s.perform(func(x, y int) ([]int, error) { return []int{x + y}, nil })
}

func subtract(s *stack) error {
	return s.perform(func(x, y int) ([]int, error) { return []int{x - y}, nil })
}

func multiply(s *stack) error {
	return s.perform(func(x, y int) ([]int, error) { return []int{x * y}, nil })
}

func divide(s *stack) error {
	return s.perform(func(x, y int) ([]int, error) {
		if y == 0 {
			return []int{}, errors.New("divide by zero")
		}

		return []int{x / y}, nil
	})
}

func duplicate(s *stack) error {
	value, err := s.peek()
	if err != nil {
		return err
	}

	s.push(value)

	return nil
}

func drop(s *stack) error {
	_, err := s.pop()
	if err != nil {
		return err
	}

	return nil
}

func swap(s *stack) error {
	return s.perform(func(x, y int) ([]int, error) {
		return []int{y, x}, nil
	})
}

func over(s *stack) error {
	return s.perform(func(x, y int) ([]int, error) {
		return []int{x, y, x}, nil
	})
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsControl(r)
}

type stackFn func(*stack) error

type evaluator struct {
	data stack
	dict map[string]stackFn
}

// Forth evaluates a slice of input strings and returns the resulting stack as a slice of ints
func Forth(input []string) ([]int, error) {
	e := evaluator{
		data: stack{},
		dict: map[string]stackFn{
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

	for _, l := range input {
		for _, w := range strings.FieldsFunc(l, isSeparator) {
			if fn := e.dict[strings.ToUpper(w)]; fn != nil {
				if err := fn(&e.data); err != nil {
					return e.data, err
				}
			} else {
				i, err := strconv.ParseInt(w, 10, 0)
				if err != nil {
					log.Println(err.Error())
					continue
				}

				e.data.push(int(i))
			}
		}
	}

	return e.data, nil
}
