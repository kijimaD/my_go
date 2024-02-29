// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2024-02-29 20:01:37

// === Go Playground ===
// Execute the snippet with:                 Ctl-Return
// Provide custom arguments to compile with: Alt-Return
// Other useful commands:
// - remove the snippet completely with its dir and all files: (go-playground-rm)
// - upload the current buffer to playground.golang.org:       (go-playground-upload)

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

type MyStruct struct {
	// Field1のコメント
	Field1 int `json:"field_1"`
	// Field2のコメント
	Field2 string `json:"field_2"`
}

func getTagByName(s interface{}, fieldName string) string {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == fieldName {
			return field.Tag.Get("json")
		}
	}

	return ""
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "snippet.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						fmt.Printf("Struct Name: %s\n", typeSpec.Name.Name)
						for _, field := range structType.Fields.List {
							if field.Doc != nil {
								my := MyStruct{}
								tag := getTagByName(my, field.Names[0].Name)

								fmt.Printf("Field: %s, Tag: %s, Comment: %s", field.Names[0].Name, tag, field.Doc.Text())
							}
						}
					}
				}
			}
		}
	}
}
