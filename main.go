package main

// imports
import (
	_ "embed"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func copyModels(destination string) {

	fs.WalkDir(embeddedModelsDir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)

		if d.IsDir() {
			err := os.Mkdir(filepath.Join(destination, path), 0755)
			if err != nil {
				log.Fatalln(err.Error())
			}
			return nil
		} else {
			var data, err1 = embeddedModelsDir.ReadFile(path)
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(destination, path), data, 0777)
		}
		return nil
	})
}

// Main docs
func main() {
	// copyModels("tata")

	fset := token.NewFileSet()
	pkgsParser, errParser := parser.ParseDir(fset, filepath.Join("go", "models"), nil, parser.ParseComments)
	if errParser != nil {
		panic(errParser)
	}
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}

}
