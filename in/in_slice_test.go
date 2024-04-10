package check

import (
	"testing"

	"github.com/leisurelicht/go-utils/gen"
)

func Test_InSlice(t *testing.T) {
	tests := []struct {
		name  string
		val   any
		slice any
		want  bool
	}{
		{"string", "a", []string{"a", "b", "c"}, true},
		{"int", 1, []int{1, 2, 3}, true},
		{"int64", int64(1), []int64{1, 2, 3}, true},
		{"flat64", float64(1), []float64{1, 2, 3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InSlice(tt.val, tt.slice)
			if got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_InStrSlice(t *testing.T) {
	tests := []struct {
		name  string
		val   string
		slice []string
		want  bool
	}{
		{"all blank", "", []string{}, false},
		{"blank target", "", []string{"a", "b", "c"}, false},
		{"black slice", "a", []string{}, false},
		{"in", "a", []string{"a", "b", "c"}, true},
		{"not in", "d", []string{"a", "b", "c"}, false},
		{"in long slice", "ab", gen.RandomStrings(1000, 2, []string{"ab"}, []string{}), true},
		{"not in long slice", "at", gen.RandomStrings(1000, 2, []string{}, []string{"at"}), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InStringSlice(tt.val, tt.slice)
			if got != tt.want {
				t.Errorf("InStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
