package main

// Немного магии с ast
// пакет позволяет анализировать или формировать новый go код
// Это основа кодогенерации

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	// Для работы ast нам нужно инициализровать набор файлов
	fset := token.NewFileSet()
	// Парс файла
	f, err := parser.ParseFile(fset, `.\7\10_ast\e.go`, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// Decls содержит в себе список всех объявленных переменных и функций
	for _, decl := range f.Decls {
		// Если мы нашли функцию
		fdecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		// Проверим, может быть это пример
		if isExample(fdecl) {
			// Проанализируем сигнатуру
			output, found := findExampleOutput(fdecl.Body, f.Comments)
			if found {
				fmt.Printf("%s needs output ‘%s’\n", fdecl.Name.Name, output)
			}
		}
	}
}

func findExampleOutput(block *ast.BlockStmt, comments []*ast.CommentGroup) (string, bool) {
	var last *ast.CommentGroup
	for _, group := range comments {
		if (block.Pos() < group.Pos()) && (block.End() > group.End()) {
			last = group
		}
	}
	if last != nil {
		text := last.Text()
		marker := "Output: "
		if strings.HasPrefix(text, marker) {
			return strings.TrimRight(text[len(marker):], "\n"), true
		}
	}
	return "", false
}

func isExample(fdecl *ast.FuncDecl) bool {
	return strings.HasPrefix(fdecl.Name.Name, "Example")
}
