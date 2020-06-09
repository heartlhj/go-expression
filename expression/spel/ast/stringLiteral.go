package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
)

/**

 */
type StringLiteral struct {
	*Literal
}

func (l StringLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
