package main

// imports
import (
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
)

func parseModel(embeddedDir embed.FS, source string, pkg *ast.Package) {

	fs.WalkDir(embeddedDir, source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)

		if d.IsDir() {
			return nil
		} else {
			var data, err1 = embeddedModelsDir.ReadFile(path)
			if err1 != nil {
				log.Fatalln(err.Error())
			}
			fset := token.NewFileSet()
			astFile, errParser := parser.ParseFile(fset, path, data, parser.ParseComments)

			pkg.Files[path] = astFile
			if errParser != nil {
				panic(errParser)
			}
			_ = astFile
			return nil
		}
	})
}

// Main docs
func main() {
	pkg := new(ast.Package)
	pkg.Files = make(map[string]*ast.File)
	parseModel(embeddedModelsDir, "go", pkg)
	log.Println("Nb of files ", len(pkg.Files))
}
