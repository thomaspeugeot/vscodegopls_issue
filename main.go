package main

// imports
import (
	"embed"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func copyModels(embeddedDir embed.FS, source, destination string) {

	fs.WalkDir(embeddedDir, source, func(path string, d fs.DirEntry, err error) error {
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
				log.Fatalln(err.Error())
			}
			return ioutil.WriteFile(filepath.Join(destination, path), data, 0777)
		}
	})
}

// Main docs
func main() {
	tmpDir := os.TempDir()
	copyModels(embeddedModelsDir, "go", tmpDir)
	defer os.RemoveAll(tmpDir)

	fset := token.NewFileSet()
	pkgsParser, errParser := parser.ParseDir(fset, filepath.Join(tmpDir, "go", "models"), nil, parser.ParseComments)
	if errParser != nil {
		panic(errParser)
	}
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}

}
