package clock

import (
	"fmt"
)

// Clock represents an instance of a clock
type Clock struct {
	h int
	m int
}

func mod(x, y int) int {
	return ((x % y) + y) % y
}

// New creates a new clock instance
func New(h int, m int) Clock {
	t := mod(h*60+m, 24*60)
	return Clock{h: t / 60, m: t % 60}
}

// Add adds the specified number of minutes returning a new clock
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

// Subtract subtracts the specified number of minutes returning a new clock
func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
