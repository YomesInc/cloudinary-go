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
	"strings"
	"text/template"
)

func main() {
	curDir, _ := os.Getwd()

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == ".go" {
				searchMixinsInFile(curDir, path)
			}
			return nil
		},
	)

	if err != nil {
		panic(err)
	}
}

func searchMixinsInFile(curDir string, fileName string) {
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

		var mixins []string

		for _, field := range s.Fields.List {
			if field.Tag != nil {
				sz := len(field.Tag.Value)
				tag := reflect.StructTag(field.Tag.Value[1 : sz-1])
				if tag.Get("mixin") != "" {
					mixins = append(mixins, tag.Get("mixin"))
				}
			}
		}

		if len(mixins) > 0 {
			generateMixins(
				curDir,
				filepath.Dir(fileName),
				currentPackage,
				currentTypeName,
				mixins,
			)
		}

		return false
	})
}

func generateMixins(curDir string, directory string, packageName string, typeName string, mixins []string) {
	for _, m := range mixins {
		t := template.Must(template.ParseFiles(fmt.Sprintf("%s/transformation/gen/%s.template", curDir, m)))

		mixinData := struct {
			PackageName string
			StructName  string
			Receiver    string
		}{
			packageName,
			typeName,
			strings.ToLower(typeName)[0:1],
		}

		f, _ := os.Create(fmt.Sprintf("%s/%s/%s_%s.go", curDir, directory, strings.ToLower(typeName), m))
		w := bufio.NewWriter(f)
		t.Execute(w, mixinData)
		w.Flush()
		f.Close()
	}
}
