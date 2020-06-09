package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
)

type LongLiteral struct {
	*Literal
}

func (l LongLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
