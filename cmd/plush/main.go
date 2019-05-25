package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/plush/v4"
)

const usage = `
Plush is a tool for managing Plush source code.

Usage:

	llush <command> [arguments]

The commands are:

	run		Executes .plush files
	fmt		plushfmt (reformat) plush sources
	ast		print the AST for a .plush file
`

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, "-h")
	}
	switch args[0] {
	case "run":
		args = args[1:]
	case "fmt":
		if err := fmtOptions.Flags.Parse(args[1:]); err != nil {
			log.Fatal(err)
		}
		format(fmtOptions.Flags.Args())
		return
	case "ast":
		printAST(args[1:])
		return
	case "-h":
		fmt.Println(strings.TrimSpace(usage))
		os.Exit(1)
	}
	run(args)
}

func run(args []string) {
	for _, a := range args {
		script, err := plush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}
		c := ast.NewContext(context.Background(), os.Stdout)

		res, err := script.Exec(c)
		if err != nil {
			log.Fatal(err)
		}

		if res == nil {
			return
		}

		if ri, ok := res.Value.([]interface{}); ok {
			for _, i := range ri {
				fmt.Println(i)
			}
			continue
		}
		fmt.Println(res)
	}
}
