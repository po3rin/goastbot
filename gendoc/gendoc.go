package gendoc

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Doc has definition & doc string.
type Doc struct {
	Definition string
	Doc        string
}

var template = `package main
func main() {
	%v
}`

// GenDoc generate doc.
func GenDoc(code string) (Doc, error) {
	fset := token.NewFileSet()
	src := fmt.Sprintf(template, code)
	f, _ := parser.ParseFile(fset, "main.go", src, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}
	return Doc{}, nil
}
