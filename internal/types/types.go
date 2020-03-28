package types

import "time"

// InterfacePair - has to structures for testing
type InterfacePair struct {
	Expect interface{}
	Actual interface{}
}

// CommonStructure - just simple structure
type CommonStructure struct {
	ValueInt    int
	ValueFloat  float64
	ValueString string
	ValueTime   time.Time
}
