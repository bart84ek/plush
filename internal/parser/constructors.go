package parser

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
)

func toScript(i interface{}) (ast.Script, error) {
	ii := reduce(i)
	bb := &bytes.Buffer{}
	for _, x := range ii {
		bb.WriteString(fmt.Sprintf("%s\n", x))
	}
	return lush.Parse("", bb.Bytes())
}

func newLushS(i interface{}) (string, error) {
	fmt.Printf("### internal/parser/constructors.go:23 i (%T) -> %q %+v\n", i, i, i)
	ii, ok := i.([]interface{})
	if !ok {
		return "", fmt.Errorf("expected []interface{} got %T", i)
	}
	fmt.Printf("### internal/parser/constructors.go:28 ii (%T) -> %q %+v\n", ii, ii, ii)
	var line string
	for _, x := range ii {
		line += fmt.Sprintf("%s", x)
	}

	// for _, x := range ii {
	// 	fmt.Printf("### internal/parser/constructors.go:30 x (%T) -> %q %+v\n", x, x, x)
	// 	switch t := x.(type) {
	// 	case []interface{}:
	// 		return newLushS(t)
	// 	default:
	// 		if t == nil {
	// 			continue
	// 		}
	// 		return fmt.Sprintf("%s", t), nil
	// 	}
	// }
	return "", nil
}

var index int

func newLushP(i interface{}) (string, error) {
	bb := &bytes.Buffer{}

	defer func() { index++ }()
	ii := reduce(i)

	// fmt.Fprintf(bb, "let lushy%d = ", index)
	var line string
	for _, x := range ii {
		line += fmt.Sprintf("%s", x)
	}

	nl := strings.HasSuffix(line, "-")
	line = strings.TrimSuffix(line, "-")
	for _, skip := range []string{"for", "if"} {
		if strings.HasPrefix(strings.TrimSpace(line), skip) {
			fmt.Fprintln(bb, line)
			return bb.String(), nil
		}
	}

	vn := fmt.Sprintf("lushy%d", index)
	fmt.Fprintf(bb, "let %s = %s\n", vn, line)
	if nl {
		fmt.Fprintf(bb, "fmt.Print(%s)\n", vn)
	} else {
		fmt.Fprintf(bb, "fmt.Println(%s)\n", vn)
	}
	fmt.Fprintf(bb, "\n")
	return bb.String(), nil
}

func html(in []byte) (string, error) {
	return fmt.Sprintf("fmt.Print(` %s`)", string(in)), nil
}

func reduce(i interface{}) []interface{} {
	var res []interface{}
	ii, ok := i.([]interface{})
	if !ok {
		log.Fatal(i)
	}

	for _, x := range ii {
		switch t := x.(type) {
		case []interface{}:
			res = append(res, reduce(t)...)
		default:
			if t == nil {
				continue
			}
			res = append(res, t)
		}
	}
	return res
}
