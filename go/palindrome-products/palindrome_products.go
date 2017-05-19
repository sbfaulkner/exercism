package palindrome

import (
	"errors"
	"fmt"
)

const testVersion = 1

// Product is a palindromic product.
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Products returns the minimum and maximum palindromic products.
func Products(fmin, fmax int) (pmin Product, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}

	for a := fmin; a <= fmax; a++ {
		for b := a; b <= fmax; b++ {
			c := a * b

			if !isPalindrome(fmt.Sprint(c)) {
				continue
			}

			if pmin.Product == 0 {
				pmin.Product = c
				pmin.Factorizations = nil
			}

			if c > pmax.Product {
				pmax.Product = c
				pmax.Factorizations = nil
			}

			if c == pmin.Product {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{a, b})
			}

			if c == pmax.Product {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{a, b})
			}
		}
	}

	if pmin.Product == 0 {
		err = errors.New("no palindromes")
		return
	}

	return
}

func isPalindrome(text string) bool {
	return text == reverse(text)
}

func reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
