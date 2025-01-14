package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Resolve_Default(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		tokenType    Type
		tokenLiteral string
	}{
		{S_START, "<%"},
		{C_START, "<%#"},
		{E_START, "<%="},
		{E_END, "%>"},
	}

	for _, tt := range tests {
		tok := Resolve(tt.tokenType)
		r.Equal(tt.tokenType, tok)
		r.Equal(tt.tokenLiteral, string(tok))
	}
}

func Test_SetTemplatingDelimiters(t *testing.T) {
	r := require.New(t)
	SetTemplatingDelimiters("{{", "}}")

	tests := []struct {
		tokenType    Type
		tokenLiteral string
	}{
		{S_START, "{{"},
		{C_START, "{{#"},
		{E_START, "{{="},
		{E_END, "}}"},
	}

	for _, tt := range tests {
		tok := Resolve(tt.tokenType)
		r.Equal(tt.tokenLiteral, string(tok))
	}
}

func Test_SetTemplatingDelimiters_LengthErrors(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		start string
		end   string
		error error
	}{
		{"{", "}", &delimiterLengthError{"{", templateDelimitersLen}},
		{"###", "###", &delimiterLengthError{"###", templateDelimitersLen}},
		{"{%", "}", &delimiterLengthError{"}", templateDelimitersLen}},
		{"{{{", "%}", &delimiterLengthError{"{{{", templateDelimitersLen}},
		{"{%", "%}", nil},
	}
	for _, tt := range tests {
		err := SetTemplatingDelimiters(tt.start, tt.end)
		r.Equal(tt.error, err)
	}
}
