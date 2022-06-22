package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `package test

    // Hello
    type A struct {
        // Where
        B int // Are you
    }
    `

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			fmt.Println(t.Doc.Text())
		case *ast.StructType:
			for _, field := range t.Fields.List {
				fmt.Println(field.Names[0], field.Doc.Text())
				fmt.Println(field.Names[0], field.Comment.Text())
			}
		}
		return true
	})
}
