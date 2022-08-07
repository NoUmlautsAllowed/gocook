package tmpl

import "testing"

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

	out := IterateRange(data, 4)

	if len(out) != len(data)/4 {
		t.Error("expected other length")
	}

	for c, i := range out {
		for d, j := range i {
			if j != data[c*4+d] {
				t.Error("expected ", data[c*4+d], "got", j)
			}
		}
	}
}
