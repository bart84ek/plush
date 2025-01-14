package ast

import (
	"bytes"
	"strings"
)

type CallExpression struct {
	TokenAble
	Callee    Expression
	Function  Expression
	Arguments []Expression
	Block     *BlockStatement
	ElseBlock *BlockStatement
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	if ce.Block != nil {
		out.WriteString(" {\n")
		out.WriteString(ce.Block.String())
		out.WriteString("}")
	}
	if ce.ElseBlock != nil {
		out.WriteString(" else { ")
		out.WriteString(ce.ElseBlock.String())
		out.WriteString(" }")
	}
	if ce.Block != nil || ce.ElseBlock != nil {
		out.WriteString("\n")
	}

	return out.String()
}
