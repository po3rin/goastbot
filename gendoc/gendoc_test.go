package gendoc

import (
	"reflect"
	"testing"
)

func TestGenDoc(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "sync.Once.Dos",
			want:  "",
		},
	}

	for _, tt := range tests {
		got, err := GenDoc(tt.input)
		if err != nil {
			t.Fatalf("Failed to test %v\n got unexpected error: %v", tt.input, err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("Failed to test %v\n got:\n %+vwant:\n %+v", tt.input, got, tt.want)
		}
	}
}
