package gendoc_test

import (
	"reflect"
	"testing"

	"github.com/po3rin/godocbot/gendoc"
)

func TestGenDoc(t *testing.T) {
	tests := []struct {
		code string
		want gendoc.Doc
	}{
		{
			code: "fmt.Println",
			want: gendoc.Doc{
				Definition: "Println func(a ...interface{}) (n int, err error)",
				Doc:        "Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.",
			},
		},
	}

	for _, tt := range tests {
		got, err := gendoc.GenDoc(tt.code)
		if err != nil {
			t.Fatalf("Failed to test %v\n got unexpected error", tt.code)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("Failed to test %v\n got: %+v\nwant: %+v", tt.code, got, tt.want)
		}
	}
}
