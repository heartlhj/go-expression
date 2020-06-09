package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
)

//map参数
const (
	THIS = "this"
	ROOT = "root"
)

type VariableReference struct {
	*SpelNodeImpl
	Name string
}

func (v VariableReference) GetValueInternal(state ExpressionState) TypedValue {
	return state.LookupVariable(v.Name)
}
