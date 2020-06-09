package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
)

type SpelNodeImpl struct {
	Children           []SpelNode
	Parent             SpelNode
	Pos                int
	exitTypeDescriptor string
}

func (this SpelNodeImpl) GetValueInternal(expressionState ExpressionState) TypedValue {
	return this.GetValueInternal(expressionState)
}

func (this SpelNodeImpl) GetValue(expressionState ExpressionState) interface{} {
	return this.GetValueInternal(expressionState)
}

func (this SpelNodeImpl) GetValueRef(state ExpressionState) ValueRef {
	return nil
}
func (this SpelNodeImpl) GetStartPosition() int {
	return this.Pos >> 16
}

func (this SpelNodeImpl) GetEndPosition() int {
	return this.Pos & 0xffff
}
