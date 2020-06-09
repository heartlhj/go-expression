package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
	"github.com/heartlhj/go-expression/expression/support"
)

//不等于
type OpNE struct {
	*Operator
}

func (o *OpNE) GetValueInternal(expressionState ExpressionState) TypedValue {
	left := o.getLeftOperand().GetValueInternal(expressionState).Value
	right := o.getRightOperand().GetValueInternal(expressionState).Value
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)
	check := !o.equalityCheck(left, right)
	value := support.BooleanTypedValue{}
	return value.ForValue(check)
}
