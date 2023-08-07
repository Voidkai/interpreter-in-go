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

	return Eval(program)
}

func TestEvalIntegerExpression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want object.Object
	}{
		{name: "test1", args: args{input: "5"}, want: &object.Integer{Value: 5}},
		{name: "test2", args: args{input: "10"}, want: &object.Integer{Value: 10}},
		{name: "test3", args: args{input: "true"}, want: &object.Boolean{Value: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := Eval(tt.args.node); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Eval() = %v, want %v", got, tt.want)
			//}
			evaluated := testEval(tt.args.input)
			testObject(t, evaluated, tt.want)
		})
	}
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

func TestEvalBooleanExpression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want object.Object
	}{
		{name: "test1", args: args{input: "true"}, want: &object.Boolean{Value: true}},
		{name: "test2", args: args{input: "false"}, want: &object.Boolean{Value: false}},
		{name: "test3", args: args{input: "1 < 2"}, want: &object.Boolean{Value: true}},
		{name: "test4", args: args{input: "1 > 2"}, want: &object.Boolean{Value: false}},
		{name: "test5", args: args{input: "1 == 2"}, want: &object.Boolean{Value: false}},
		{name: "test6", args: args{input: "1 != 2"}, want: &object.Boolean{Value: true}},
		{name: "test7", args: args{input: "true == true"}, want: &object.Boolean{Value: true}},
		{name: "test8", args: args{input: "true != true"}, want: &object.Boolean{Value: false}},
		{name: "test9", args: args{input: "false == false"}, want: &object.Boolean{Value: true}},
		{name: "test10", args: args{input: "false != false"}, want: &object.Boolean{Value: false}},
		{name: "test11", args: args{input: "true == false"}, want: &object.Boolean{Value: false}},
		{name: "test12", args: args{input: "true != false"}, want: &object.Boolean{Value: true}},
		{name: "test13", args: args{input: "false == true"}, want: &object.Boolean{Value: false}},
		{name: "test14", args: args{input: "false != true"}, want: &object.Boolean{Value: true}},
		{name: "test15", args: args{input: "(1 < 2) == true"}, want: &object.Boolean{Value: true}},
		{name: "test16", args: args{input: "(1 < 2) == false"}, want: &object.Boolean{Value: false}},
		{name: "test17", args: args{input: "(1+2)*3 == 9"}, want: &object.Boolean{Value: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(tt.args.input)
			testObject(t, evaluated, tt.want)
		})
	}
}
