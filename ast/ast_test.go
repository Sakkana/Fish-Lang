package ast

import (
	"Fish-Lang/token"
	"log"
	"testing"
)

// let statement
func TestStringLet(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Indentifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Indentifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVal"},
					Value: "anotherVar",
				},
			},
		},
	}

	template := "let myVar = anotherVar;"
	if program.String() != template {
		for i, c := range program.String() {
			log.Println(string(template[i]), " ---> ", string(c))
		}
		t.Errorf("program.String() wrong. got=%q",
			program.String())
	}

}
