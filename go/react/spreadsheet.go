package react

const testVersion = 5

// New creates a new Spreadsheet.
func New() Spreadsheet {
	return Spreadsheet{}
}

type Spreadsheet struct{}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (spreadsheet Spreadsheet) CreateInput(value int) InputCell {
	return SpreadsheetCell{value: value}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (spreadsheet Spreadsheet) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	return SpreadsheetCell{compute1: compute, references: []Cell{cell}}
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (spreadsheet Spreadsheet) CreateCompute2(cell1 Cell, cell2 Cell, compute func(int, int) int) ComputeCell {
	return SpreadsheetCell{compute2: compute, references: []Cell{cell1, cell2}}
}
