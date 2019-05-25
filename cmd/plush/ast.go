package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/plush/v4"
)

const astUsage = `
usage: plush ast [files]
`

func printAST(args []string) {
	if len(args) == 0 {
		fmt.Println(strings.TrimSpace(astUsage))
		os.Exit(1)
	}
	for _, a := range args {
		script, err := plush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(script, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
