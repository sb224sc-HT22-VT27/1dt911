package mathmetics_test

import (
	"lunit-test/calc/mathmetics"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases_wrapangle = []struct {
	name     string
	angle    float64
	expected float64
}{
	{"Less then 360", 200.0, 200.0},
	{"more then 360", 400.0, 40.0},
}

func TestWrapAngle(t *testing.T) {
	for _, tc := range testcases_wrapangle {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			got := mathmetics.WrapAngle_wrapper(tc.angle)

			assert.Equal(tc.expected, got)
		})
	}
}
