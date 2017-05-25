package react

import (
	"fmt"
)

var callbackID int

// SpreadsheetCell defines a cell in a spreadsheet.
type SpreadsheetCell struct {
	value      int
	dependents []spreadsheetCellDependency
	callbacks  map[int]func(int)
	dirty      bool
}

type spreadsheetCellDependency struct {
	cell   *SpreadsheetCell
	update func(int)
}

// NewCell creates a new SpreadsheetCell.
func NewCell(value int) *SpreadsheetCell {
	return &SpreadsheetCell{value: value, callbacks: map[int]func(int){}}
}

// Value returns the current value of the cell.
func (cell *SpreadsheetCell) Value() int {
	return cell.value
}

// SetValue sets the value of the cell.
func (cell *SpreadsheetCell) SetValue(value int) {
	cell.setValue(value)
	cell.performCallbacks()
}

func (cell *SpreadsheetCell) setValue(value int) {
	if cell.value != value {
		cell.value = value
		cell.dirty = true

		for _, dependent := range cell.dependents {
			dependent.update(value)
		}
	}
}

func (cell *SpreadsheetCell) performCallbacks() {
	if cell.dirty {
		for _, callback := range cell.callbacks {
			callback(cell.value)
		}

		for _, dependent := range cell.dependents {
			dependent.cell.performCallbacks()
		}

		cell.dirty = false
	}
}

type spreadsheetCallbackCanceler struct {
	cell  *SpreadsheetCell
	index int
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (cell *SpreadsheetCell) AddCallback(callback func(int)) Canceler {
	callbackID++
	cell.callbacks[callbackID] = callback
	return spreadsheetCallbackCanceler{cell: cell, index: callbackID}
}

// Cancel removes the callback.
func (canceler spreadsheetCallbackCanceler) Cancel() {
	delete(canceler.cell.callbacks, canceler.index)
}

func (cell *SpreadsheetCell) addDependent(dependent *SpreadsheetCell, update func(int)) {
	cell.dependents = append(cell.dependents, spreadsheetCellDependency{cell: dependent, update: update})
}
