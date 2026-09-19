package parser

type Operand struct {
	Previous bool
	Value string
}

type Operation struct {
	Opertator string
	Left Operand
	Right Operand
	Pipe bool
	Output bool
	Sleep bool
	SleepTime int
}