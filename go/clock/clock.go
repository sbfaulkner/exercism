package clock

import "fmt"

const testVersion = 4

// Clock is a basic type to support simple time operations
type Clock struct {
	hour   int
	minute int
}

// New constructs a new Clock type
func New(hour, minute int) Clock {
	clock := Clock{hour, minute}
	clock.normalize()
	return clock
}

// String converts a Clock type to a 24-hour clock string representation
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add adds a given number of minute to a Clock type
func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes
	clock.normalize()
	return clock
}

// normalize is a helper function to clean up negative hours or minutes for a Clock type
func (clock *Clock) normalize() {
	minutes := clock.hour*60 + clock.minute
	minutes %= 24 * 60
	minutes += 24 * 60

	clock.minute = minutes % 60
	clock.hour = (minutes / 60) % 24
}
