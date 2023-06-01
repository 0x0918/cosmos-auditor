package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func generateAST(path string) *ast.File {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	return f
}