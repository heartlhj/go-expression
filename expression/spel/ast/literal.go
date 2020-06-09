package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
)

type GetLiteralValue interface {
	GetLiteralValue() TypedValue
}

type Literal struct {
	*SpelNodeImpl
	OriginalValue string
	Value         TypedValue
}
