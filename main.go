package main

// imports
import (
	"embed"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
)

func parseModel(embeddedDir embed.FS, source string) {

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
			pkgsParser, errParser := parser.ParseFile(fset, path, data, parser.ParseComments)
			if errParser != nil {
				panic(errParser)
			}
			_ = pkgsParser
			return nil
		}
	})
}

// Main docs
func main() {
	parseModel(embeddedModelsDir, "go")
}
