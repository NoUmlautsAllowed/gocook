package form

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	type q struct {
		Query  string `form:"query"`
		Tags   string `form:"tags"`
		Number int    `form:"number"`
		Other  string
	}

	tests := []struct {
		name  string
		want  string
		input q
	}{
		{
			name:  "basic",
			want:  "number=1&query=abc&tags=x%2Cy",
			input: q{Query: "abc", Tags: `x,y`, Number: 1},
		},
		{
			name:  "empty",
			want:  "number=0&query=abc",
			input: q{Query: "abc", Number: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Marshal(tt.input); err != nil {
				t.Errorf("Marshal() error = %v", err)
			} else if string(got) != tt.want {
				t.Errorf("Marshal() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestMarshalError(t *testing.T) {
	v := 1
	got, err := Marshal(v)
	if err == nil {
		t.Errorf("Marshal() error = %v, want err", err)
	}
	if got != nil {
		t.Errorf("Marshal() = %v, want nil", got)
	}
}
