package evaluator

import (
	"Interpreter_in_Go/ast"
	"Interpreter_in_Go/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.ReturnStatement:
		reVal := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: reVal}

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return boolNativeToBoolObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		lt := Eval(node.Left)
		rt := Eval(node.Right)
		return evalInfixExpression(node.Operator, lt, rt)

	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.IfExpression:
		return evalConditionalExpression(node)
	default:
		return nil
	}
}

func evalProgram(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
		if returnVal, ok := result.(*object.ReturnValue); ok {
			return returnVal.Value
		}
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, stmt := range block.Statements {
		result = Eval(stmt)
		if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
			return result
		}
	}
	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalPrefixNegationExpression(right)
	default:
		return NULL
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return boolNativeToBoolObject(left == right)
	case operator == "!=":
		return boolNativeToBoolObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	ltVal := left.(*object.Integer).Value
	rtVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: ltVal + rtVal}
	case "-":
		return &object.Integer{Value: ltVal - rtVal}
	case "*":
		return &object.Integer{Value: ltVal * rtVal}
	case "/":
		return &object.Integer{Value: ltVal / rtVal}

	case "<":
		return boolNativeToBoolObject(ltVal < rtVal)
	case ">":
		return boolNativeToBoolObject(ltVal > rtVal)
	case "==":
		return boolNativeToBoolObject(ltVal == rtVal)
	case "!=":
		return boolNativeToBoolObject(ltVal != rtVal)
	default:
		return NULL
	}
}

func evalConditionalExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func evalPrefixNegationExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case FALSE:
		return TRUE
	case TRUE:
		return FALSE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func boolNativeToBoolObject(value bool) *object.Boolean {
	if value {
		return TRUE
	} else {
		return FALSE
	}
}
