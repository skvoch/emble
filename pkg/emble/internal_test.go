package emble_test

import (
	"testing"

	types "github.com/skvoch/emble/internal/types"
	emble "github.com/skvoch/emble/pkg/emble"
)

func Test_BasicCase(t *testing.T) {

	cases := []struct {
		Name       string
		Structures types.InterfacePair
	}{
		{
			Name:       "Common structures",
			Structures: emble.MakeCommonStructures(),
		},
	}

	for _, c := range cases {

	}

}
