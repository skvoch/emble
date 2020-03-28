package emble

import (
	"strconv"
	"time"

	"math/rand"

	types "github.com/skvoch/emble/internal/types"
)

func randomInt() int {
	return rand.Int()
}

func randomString() string {
	value := rand.Int()

	return strconv.Itoa(value)
}

func randomFloat64() float64 {
	return rand.Float64()
}

// MakeCommonStructures - make two stuctures with different time
func MakeCommonStructures() types.InterfacePair {
	ValueInt := randomInt()
	ValueFloat := randomFloat64()
	ValueString := randomString()

	expect := &types.CommonStructure{
		ValueInt:    ValueInt,
		ValueFloat:  ValueFloat,
		ValueString: ValueString,
		ValueTime:   time.Now(),
	}

	actual := &types.CommonStructure{
		ValueInt:    ValueInt,
		ValueFloat:  ValueFloat,
		ValueString: ValueString,
		ValueTime:   time.Now(),
	}

	return types.InterfacePair{
		Expect: expect,
		Actual: actual,
	}
}
