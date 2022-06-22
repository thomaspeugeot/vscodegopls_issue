package main

// imports
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Estruct should not be generated
//
// swagger:ignore
type Estruct struct {
	Name string
}

// Main docs
func main() {
	fset := token.NewFileSet() // positions are relative to fset

	d, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range d {
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				fmt.Printf("%s:\tFuncDecl %s\t%s\n", fset.Position(n.Pos()), x.Name, x.Doc.Text())
			case *ast.TypeSpec:
				fmt.Printf("%s:\tTypeSpec %s\t%s\n", fset.Position(n.Pos()), x.Name, x.Doc.Text())
			case *ast.Field:
				fmt.Printf("%s:\tField %s\t%s\n", fset.Position(n.Pos()), x.Names, x.Doc.Text())
			case *ast.GenDecl:
				fmt.Printf("%s:\tGenDecl %s\n", fset.Position(n.Pos()), x.Doc.Text())
			}

			return true
		})
	}
}
