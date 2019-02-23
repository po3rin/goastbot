package gendoc

import (
	"go/ast"
	"go/parser"
)

// Doc has definition & doc string.
type Doc struct {
	Definition string
	Doc        string
}

// GenDoc generate doc.
func GenDoc(code string) (Doc, error) {
	expr, _ := parser.ParseExpr("1+1")
	ast.Print(nil, expr)
	return Doc{}, nil
}
