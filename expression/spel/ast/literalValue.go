package ast

import . "github.com/heartlhj/go-expression/expression/spel"

type LiteralValue struct {
	*SpelNodeImpl
}

func (l *LiteralValue) GetLiteralValue() TypedValue {
	return TypedValue{}
}
