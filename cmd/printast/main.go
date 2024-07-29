package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		_, _ = fmt.Printf("usage: %s <FILE>\n", os.Args[0])

		return
	}

	file, err := parser.ParseFile(token.NewFileSet(), args[0], nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	err = ast.Print(token.NewFileSet(), file)
	if err != nil {
		log.Fatal(err)
	}
}
