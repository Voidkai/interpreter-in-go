package object

import (
	"interpreter-in-go/ast"
	"interpreter-in-go/token"
	"testing"
)

func TestBoolean_Inspect(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boolean{
				Value: tt.fields.Value,
			}
			if got := b.Inspect(); got != tt.want {
				t.Errorf("Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolean_Type(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   ObjectType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Boolean{
				Value: tt.fields.Value,
			}
			if got := b.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Inspect(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Message: tt.fields.Message,
			}
			if got := e.Inspect(); got != tt.want {
				t.Errorf("Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Type(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   ObjectType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Message: tt.fields.Message,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Inspect(t *testing.T) {
	type fields struct {
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Env        *Environment
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
				Parameters: []*ast.Identifier{{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"}},
				Body: &ast.BlockStatement{
					Token: token.Token{
						Type:    token.LBRACE,
						Literal: "{",
					},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "x",
							},
							Expression: &ast.InfixExpression{
								Token: token.Token{
									Type:    token.PLUS,
									Literal: "+",
								},
								Left: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "x",
									},
									Value: "x",
								},
								Operator: "+",
								Right: &ast.IntegerLiteral{
									Token: token.Token{
										Type:    token.INT,
										Literal: "2",
									},
									Value: 2,
								},
							},
						},
					},
				},
				Env: &Environment{},
			},
			want: "fn(x) {\n(x + 2)\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				Parameters: tt.fields.Parameters,
				Body:       tt.fields.Body,
				Env:        tt.fields.Env,
			}
			if got := f.Inspect(); got != tt.want {
				t.Errorf("Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Type(t *testing.T) {
	type fields struct {
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Env        *Environment
	}
	tests := []struct {
		name   string
		fields fields
		want   ObjectType
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				Parameters: []*ast.Identifier{{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"}},
				Body: &ast.BlockStatement{
					Token: token.Token{
						Type:    token.LBRACE,
						Literal: "{",
					},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Token: token.Token{
								Type:    token.IDENT,
								Literal: "x",
							},
							Expression: &ast.InfixExpression{
								Token: token.Token{
									Type:    token.PLUS,
									Literal: "+",
								},
								Left: &ast.Identifier{
									Token: token.Token{
										Type:    token.IDENT,
										Literal: "x",
									},
									Value: "x",
								},
								Operator: "+",
								Right: &ast.IntegerLiteral{
									Token: token.Token{
										Type:    token.INT,
										Literal: "2",
									},
									Value: 2,
								},
							},
						},
					},
				},
				Env: &Environment{},
			},
			want: "RETURN_VALUE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				Parameters: tt.fields.Parameters,
				Body:       tt.fields.Body,
				Env:        tt.fields.Env,
			}
			if got := f.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_Inspect(t *testing.T) {
	type fields struct {
		Value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integer{
				Value: tt.fields.Value,
			}
			if got := i.Inspect(); got != tt.want {
				t.Errorf("Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_Type(t *testing.T) {
	type fields struct {
		Value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   ObjectType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integer{
				Value: tt.fields.Value,
			}
			if got := i.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNull_Inspect(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Null{}
			if got := n.Inspect(); got != tt.want {
				t.Errorf("Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNull_Type(t *testing.T) {
	tests := []struct {
		name string
		want ObjectType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Null{}
			if got := n.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnValue_Inspect(t *testing.T) {
	type fields struct {
		Value Object
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &ReturnValue{
				Value: tt.fields.Value,
			}
			if got := rv.Inspect(); got != tt.want {
				t.Errorf("Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnValue_Type(t *testing.T) {
	type fields struct {
		Value Object
	}
	tests := []struct {
		name   string
		fields fields
		want   ObjectType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv := &ReturnValue{
				Value: tt.fields.Value,
			}
			if got := rv.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
