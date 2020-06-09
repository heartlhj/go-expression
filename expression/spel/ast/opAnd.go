package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
	. "github.com/heartlhj/go-expression/expression/support"
)

//与操作 eg:  "#name=='lisi' && #age>=18"
type OpAnd struct {
	*Operator
}

func (o *OpAnd) GetValueInternal(expressionState ExpressionState) TypedValue {
	if !getBooleanValue(expressionState, o.getLeftOperand()) {
		value := BooleanTypedValue{}
		return value.ForValue(false)
	}
	booleanValue := getBooleanValue(expressionState, o.getRightOperand())
	value := BooleanTypedValue{}
	return value.ForValue(booleanValue)
}

func getBooleanValue(state ExpressionState, operand SpelNode) bool {
	value := operand.GetValueInternal(state)
	if value.Value == nil {
		panic("Type conversion problem, cannot convert from [null] to bool")
	}
	return value.Value.(bool)
}
