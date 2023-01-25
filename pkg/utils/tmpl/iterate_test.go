package tmpl

import (
	"math"
	"testing"
)

func Test_IterateRange(t *testing.T) {
	data := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
	}

	steps := 4
	out := IterateRange(data, steps)

	expected := int(math.Ceil(float64(len(data)) / float64(steps)))

	if len(out) != expected {
		t.Error("expected other length, got", len(out), "!=", expected)
	}

	for c, i := range out {
		for d, j := range i {
			if j != data[c*4+d] {
				t.Error("expected ", data[c*4+d], "got", j)
			}
		}
	}
}
