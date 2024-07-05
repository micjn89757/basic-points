package main

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/java"
)


func main() {
	parser := sitter.NewParser()
	parser.SetLanguage(java.GetLanguage())

	sourceCode := []byte("public static void main(String[] args) {int i = 1;}")

	tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		panic(err)
	}

	n := tree.RootNode()

	fmt.Println(n)
	fmt.Println(n.ChildCount())

	fmt.Println(n.NamedChild(0))
	fmt.Println(n.Child(0))

}