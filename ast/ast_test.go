package ast

import (
	"interpreter-in-go/token"
	"testing"
)

func TestLetStatement_TokenLiteral(t *testing.T) {
	type fields struct {
		Token token.Token
		Name  *Identifier
		Value Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: "test1", fields: fields{
			Token: token.Token{
				Type:    token.IDENT,
				Literal: "x",
			},
			Name: &Identifier{
				Token: token.Token{
					Type:    token.IDENT,
					Literal: "x",
				},
				Value: "x",
			},
			Value: nil,
		}, want: "x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LetStatement{
				Token: tt.fields.Token,
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := ls.TokenLiteral(); got != tt.want {
				t.Errorf("TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}
