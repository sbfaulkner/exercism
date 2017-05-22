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
	return &SpreadsheetCell{value: value}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (spreadsheet Spreadsheet) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	v := compute(cell.Value())
	c := &SpreadsheetCell{value: v}
	cell.(*SpreadsheetCell).AddCallback(func(value int) { c.SetValue(compute(value)) })
	return c
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (spreadsheet Spreadsheet) CreateCompute2(cell1 Cell, cell2 Cell, compute func(int, int) int) ComputeCell {
	v := compute(cell1.Value(), cell2.Value())
	c := &SpreadsheetCell{value: v}
	cell1.(*SpreadsheetCell).AddCallback(func(value1 int) { c.SetValue(compute(value1, cell2.Value())) })
	cell2.(*SpreadsheetCell).AddCallback(func(value2 int) { c.SetValue(compute(cell1.Value(), value2)) })
	return c
}
