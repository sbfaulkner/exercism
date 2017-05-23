package react

type spreadsheetCell struct {
	value     int
	callbacks []func(int)
}

// Value returns the current value of the cell.
func (cell *spreadsheetCell) Value() int {
	return cell.value
}

// SetValue sets the value of the cell.
func (cell *spreadsheetCell) SetValue(value int) {
	if cell.value != value {
		cell.value = value
		for _, callback := range cell.callbacks {
			if callback != nil {
				callback(value)
			}
		}
	}
}

type spreadsheetCallbackCanceler struct {
	cell  *spreadsheetCell
	index int
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (cell *spreadsheetCell) AddCallback(callback func(int)) Canceler {
	cell.callbacks = append(cell.callbacks, callback)
	return spreadsheetCallbackCanceler{cell: cell, index: len(cell.callbacks) - 1}
}

// Cancel removes the callback.
func (canceler spreadsheetCallbackCanceler) Cancel() {
	for i := range canceler.cell.callbacks {
		if i == canceler.index {
			canceler.cell.callbacks[i] = nil
			return
		}
	}
}
