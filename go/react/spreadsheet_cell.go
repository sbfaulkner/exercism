package react

import "fmt"

type SpreadsheetCell struct {
	value      int
	compute1   func(int) int
	compute2   func(int, int) int
	references []Cell
	callback   func(int)
}

// Value returns the current value of the cell.
func (cell SpreadsheetCell) Value() int {
	if cell.compute1 != nil {
		return cell.compute1(cell.references[0].Value())
	}

	if cell.compute2 != nil {
		return cell.compute2(cell.references[0].Value(), cell.references[1].Value())
	}

	return cell.value
}

// SetValue sets the value of the cell.
func (cell SpreadsheetCell) SetValue(value int) {
	fmt.Println("SetValue", cell, value)
	cell.value = value
	fmt.Println("  cell.callback", cell.callback)
	if cell.callback != nil {
		cell.callback(value)
	}
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (cell SpreadsheetCell) AddCallback(callback func(int)) Canceler {
	fmt.Println("AddCallback", callback)
	cell.callback = callback
	fmt.Println("  cell.callback", cell.callback)
	return cell
}

// Cancel removes the callback.
func (cell SpreadsheetCell) Cancel() {

}
