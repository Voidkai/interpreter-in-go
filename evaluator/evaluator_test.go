package evaluator

import (
	"interpreter-in-go/lexer"
	"interpreter-in-go/object"
	"interpreter-in-go/parser"
	"testing"
)

func testObject(t *testing.T, evaluated object.Object, want object.Object) {
	switch result := evaluated.(type) {
	case *object.Integer:
		if result.Value != want.(*object.Integer).Value {
			t.Errorf("object has wrong value. got=%d, want=%d", result.Value, want.(*object.Integer).Value)
		}
	case *object.Boolean:
		if result.Value != want.(*object.Boolean).Value {
			t.Errorf("object has wrong value. got=%t, want=%t", result.Value, want.(*object.Boolean).Value)
		}
	default:
		t.Errorf("object is not Integer or Boolean. got=%T(%+v)", evaluated, evaluated)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func TestBangOperator(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want object.Object
	}{
		{name: "test1", args: args{input: "!true"}, want: &object.Boolean{Value: false}},
		{name: "test2", args: args{input: "!false"}, want: &object.Boolean{Value: true}},
		{name: "test3", args: args{input: "!5"}, want: &object.Boolean{Value: false}},
		{name: "test4", args: args{input: "!!true"}, want: &object.Boolean{Value: true}},
		{name: "test5", args: args{input: "!!false"}, want: &object.Boolean{Value: false}},
		{name: "test6", args: args{input: "!!5"}, want: &object.Boolean{Value: true}},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.args.input)
		testObject(t, evaluated, tt.want)
	}
}

func TestEvalExpression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want object.Object
	}{
		{name: "Boolean_test1", args: args{input: "true"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test2", args: args{input: "false"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test3", args: args{input: "1 < 2"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test4", args: args{input: "1 > 2"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test5", args: args{input: "1 == 2"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test6", args: args{input: "1 != 2"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test7", args: args{input: "true == true"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test8", args: args{input: "true != true"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test9", args: args{input: "false == false"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test10", args: args{input: "false != false"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test11", args: args{input: "true == false"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test12", args: args{input: "true != false"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test13", args: args{input: "false == true"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test14", args: args{input: "false != true"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test15", args: args{input: "(1 < 2) == true"}, want: &object.Boolean{Value: true}},
		{name: "Boolean_test16", args: args{input: "(1 < 2) == false"}, want: &object.Boolean{Value: false}},
		{name: "Boolean_test17", args: args{input: "(1+2)*3 == 9"}, want: &object.Boolean{Value: false}},
		{name: "Integer_test1", args: args{input: "5"}, want: &object.Integer{Value: 5}},
		{name: "Integer_test2", args: args{input: "10"}, want: &object.Integer{Value: 10}},
		{name: "Integer_test3", args: args{input: "true"}, want: &object.Boolean{Value: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(tt.args.input)
			testObject(t, evaluated, tt.want)
		})
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{name: "test1", input: "if (true) { 10 }", want: 10},
		{name: "test2", input: "if (false) { 10 }", want: nil},
		{name: "test3", input: "if (1) { 10 }", want: 10},
		{name: "test4", input: "if (1 < 2) { 10 }", want: 10},
		{name: "test5", input: "if (1 > 2) { 10 }", want: nil},
		{name: "test6", input: "if (1 > 2) { 10 } else { 20 }", want: 20},
		{name: "test7", input: "if (1 < 2) { 10 } else { 20 }", want: 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.want.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, want int64) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != want {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, want)
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{name: "test1", input: "return 10;", want: 10},
		{name: "test2", input: "return 10; 9;", want: 10},
		{name: "test3", input: "return 2*5; 9;", want: 10},
		{name: "test4", input: "9; return 2*5; 9;", want: 10},
		{name: "test5", input: "if (10 > 1) { if (10 > 1) { return 10; } return 1; }", want: 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.want)
	}
}

func TestErrorHanding(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantMessage string
	}{
		{name: "test1", input: "5 + true;", wantMessage: "type mismatch: INTEGER + BOOLEAN"},
		{name: "test2", input: "5 + true; 5;", wantMessage: "type mismatch: INTEGER + BOOLEAN"},
		{name: "test3", input: "-true", wantMessage: "unknown operator: -BOOLEAN"},
		{name: "test4", input: "true + false;", wantMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{name: "test5", input: "5; true + false; 5", wantMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{name: "test6", input: "if (10 > 1) { true + false; }", wantMessage: "unknown operator: BOOLEAN + BOOLEAN"},
		{name: "test7", input: "if (10 > 1) { if (10 > 1) { return true + false; } return 1; }", wantMessage: "unknown operator: BOOLEAN + BOOLEAN"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("%s: no error object returned. got=%T(%+v)", tt.name, evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.wantMessage {
			t.Errorf("wrong error message. want=%q, got=%q", tt.wantMessage, errObj.Message)
		}
	}

}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{name: "test1", input: "let a = 5; a;", expected: 5},
		{name: "test2", input: "let a = 5 * 5; a;", expected: 25},
		{name: "test3", input: "let a = 5; let b = a; b;", expected: 5},
		{name: "test4", input: "let a = 5; let b = a; let c = a + b + 5; c;", expected: 15},
		{name: "test4", input: "let a = 5; let b = a; let c = a + b + 5; c;", expected: 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T(%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}

}

func TestStringInfixExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "test1", input: `"Hello" + " " + "World!"`, expected: "Hello World!"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T(%+v)", evaluated, evaluated)
		}
		if str.Value != tt.expected {
			t.Errorf("String has wrong value. got=%q", str.Value)
		}
	}
}

func TestBuiltinFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{name: "test1", input: `len("")`, expected: 0},
		{name: "test2", input: `len("four")`, expected: 4},
		{name: "test3", input: `len("hello world")`, expected: 11},
		{name: "test4", input: `len(1)`, expected: "argument to `len` not supported, got INTEGER"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("object is not Error. got=%T(%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. want=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T(%+v)", evaluated, evaluated)
	}
	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}
	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)

}

func TestHashLiterals(t *testing.T) {
	input := `
	{
		"one": 10 - 9,
		"two": 1 + 1,
		"thr" + "ee": 6 / 2
}`
	evaluated := testEval(input)
	hash, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T(%+v)", evaluated, evaluated)
	}
	expected := map[object.HashKey]object.Object{
		(&object.String{Value: "one"}).HashKey():   &object.Integer{Value: 1},
		(&object.String{Value: "two"}).HashKey():   &object.Integer{Value: 2},
		(&object.String{Value: "three"}).HashKey(): &object.Integer{Value: 3},
	}
	if len(hash.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d， excepted = %d", len(hash.Pairs), len(expected))
	}
	for expectedKey, expectedValue := range expected {
		pair, ok := hash.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}
		testIntegerObject(t, pair.Value, expectedValue.(*object.Integer).Value)
	}
}

func TestHashIndexExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{name: "test1", input: `{"foo": 5}["foo"]`, expected: 5},
		{name: "test2", input: `{"foo": 5}["bar"]`, expected: nil},
		{name: "test3", input: `let key = "foo"; {"foo": 5}[key]`, expected: 5},
		{name: "test4", input: `{}["foo"]`, expected: nil},
		{name: "test5", input: `{5: 5}[5]`, expected: 5},
		{name: "test6", input: `{true: 5}[true]`, expected: 5},
		{name: "test7", input: `{false: 5}[false]`, expected: 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
