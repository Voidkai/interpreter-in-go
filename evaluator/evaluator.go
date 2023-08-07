package evaluator

import (
	"interpreter-in-go/ast"
	"interpreter-in-go/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativateBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}

	return nil
}

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case operator == "+" || operator == "-" || operator == "*" || operator == "/" || operator == "<" || operator == ">":
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativateBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativateBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func nativateBoolToBooleanObject(b bool) object.Object {
	if b == true {
		return TRUE
	} else {
		return FALSE
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	case "<":
		return nativateBoolToBooleanObject(leftValue < rightValue)
	case ">":
		return nativateBoolToBooleanObject(leftValue > rightValue)
	case "==":
		if leftValue == rightValue {
			return TRUE
		} else {
			return FALSE
		}
	case "!=":
		if leftValue != rightValue {
			return TRUE
		} else {
			return FALSE
		}
	default:
		return NULL
	}
}
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return nil
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return nil
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
	}
	return result
}
