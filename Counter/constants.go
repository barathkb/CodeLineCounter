package Counter

// LineType represents the type of line in the code file
type LineType int

const (
	Blank LineType = iota
	Comment
	Code
	Import
	Variable
)
