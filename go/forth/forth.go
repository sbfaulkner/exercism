package forth

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 1

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsControl(r)
}

func peek(stack *[]int) (value int, err error) {
	if len(*stack) < 1 {
		err = errors.New("empty stack")
		return
	}

	return (*stack)[len(*stack)-1], nil
}

func pop(stack *[]int) (value int, err error) {
	value, err = peek(stack)
	if err != nil {
		return
	}

	*stack = (*stack)[:len(*stack)-1]

	return
}

func push(stack *[]int, value int) {
	*stack = append(*stack, value)
}

func perform(stack *[]int, f func(x, y int) ([]int, error)) (err error) {
	var x, y int
	var result []int

	y, err = pop(stack)
	if err != nil {
		return
	}

	x, err = pop(stack)
	if err != nil {
		return
	}

	result, err = f(x, y)
	if err != nil {
		return
	}

	for _, r := range result {
		push(stack, r)
	}

	return
}

func add(stack *[]int) error {
	return perform(stack, func(x, y int) ([]int, error) { return []int{x + y}, nil })
}

func subtract(stack *[]int) error {
	return perform(stack, func(x, y int) ([]int, error) { return []int{x - y}, nil })
}

func multiply(stack *[]int) error {
	return perform(stack, func(x, y int) ([]int, error) { return []int{x * y}, nil })
}

func divide(stack *[]int) error {
	return perform(stack, func(x, y int) ([]int, error) {
		if y == 0 {
			return []int{}, errors.New("divide by zero")
		}

		return []int{x / y}, nil
	})
}

func duplicate(stack *[]int) error {
	value, err := peek(stack)
	if err != nil {
		return err
	}

	push(stack, value)

	return nil
}

func drop(stack *[]int) error {
	_, err := pop(stack)
	if err != nil {
		return err
	}

	return nil
}

func swap(stack *[]int) error {
	return perform(stack, func(x, y int) ([]int, error) {
		return []int{y, x}, nil
	})
}

func over(stack *[]int) error {
	return perform(stack, func(x, y int) ([]int, error) {
		return []int{x, y, x}, nil
	})
}

// Forth evaluates a slice of input strings and returns the resulting stack as a slice of ints
func Forth(input []string) ([]int, error) {
	stack := []int{}

	dictionary := map[string]func(stack *[]int) error{
		"+":    add,
		"-":    subtract,
		"*":    multiply,
		"/":    divide,
		"DUP":  duplicate,
		"DROP": drop,
		"SWAP": swap,
		"OVER": over,
	}

	for _, l := range input {
		for _, w := range strings.FieldsFunc(l, isSeparator) {
			if f := dictionary[strings.ToUpper(w)]; f != nil {
				if err := f(&stack); err != nil {
					return stack, err
				}
			} else {
				i, err := strconv.ParseInt(w, 10, 0)
				if err != nil {
					log.Println(err.Error())
					continue
				}

				stack = append(stack, int(i))
			}
		}
	}

	return stack, nil
}
