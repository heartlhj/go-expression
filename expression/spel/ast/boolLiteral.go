package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
)

type BoolLiteral struct {
	*Literal
}

func (l BoolLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
