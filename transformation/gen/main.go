package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
)

type writeableDefinition struct {
	Path         string
	PackageName  string
	StructEntity string
}

func main() {
	var definitions []writeableDefinition
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == ".go" {
				definitions = append(definitions, searchMixinsInFile(path)...)
			}
			return nil
		},
	)

	if err != nil {
		panic(err)
	}

	if len(definitions) > 0 {
		curDir, _ := os.Getwd()
		writeDefinitions(definitions, curDir)
	}
}

func searchMixinsInFile(fileName string) []writeableDefinition {
	definitions := []writeableDefinition{}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var currentTypeName string
	var currentPackage string

	ast.Inspect(file, func(x ast.Node) bool {
		pkg, ok := x.(*ast.File)
		if ok {
			currentPackage = pkg.Name.String()
		}

		ss, ok := x.(*ast.TypeSpec)

		if ok {
			currentTypeName = ss.Name.String()
		}

		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range s.Fields.List {
			if field.Tag != nil {
				sz := len(field.Tag.Value)
				tag := reflect.StructTag(field.Tag.Value[1 : sz-1])
				if tag.Get("mixin") != "" {
					definitions = append(definitions, writeableDefinition{filepath.Dir(fileName), currentPackage, fmt.Sprintf("%s.%s", currentPackage, currentTypeName)})
				}
			}
		}

		return false
	})

	return definitions
}

func writeDefinitions(definitions []writeableDefinition, baseDir string) {
	t := template.Must(template.ParseFiles(fmt.Sprintf("%s/transformation/gen/generator.template", baseDir)))
	f, _ := os.Create(fmt.Sprintf("%s/transformation/gen/generator.go", baseDir))
	w := bufio.NewWriter(f)

	templateData := struct {
		Definitions []writeableDefinition
	}{Definitions: definitions}

	t.Execute(w, templateData)
	w.Flush()
	f.Close()
}
