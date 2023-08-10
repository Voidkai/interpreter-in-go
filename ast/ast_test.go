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

func TestProgram_String(t *testing.T) {
	type fields struct {
		Statements []Statement
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "test1", fields: fields{Statements: []Statement{&LetStatement{
			Token: token.Token{Type: token.LET, Literal: "let"},
			Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
			Value: &Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
				Value: "anotherVar",
			},
		}}}, want: "let myVar = anotherVar;",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Statements: tt.fields.Statements,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() wrong. got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringLiteral_TokenLiteral(t *testing.T) {
	type fields struct {
		Token token.Token
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				Token: token.Token{
					Type:    token.STRING,
					Literal: "hello",
				},
				Value: "hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &StringLiteral{
				Token: tt.fields.Token,
				Value: tt.fields.Value,
			}
			if got := sl.TokenLiteral(); got != tt.want {
				t.Errorf("TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}
