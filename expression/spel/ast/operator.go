package ast

import "reflect"

type Operator struct {
	*SpelNodeImpl
	operatorName          string
	leftActualDescriptor  string
	rightActualDescriptor string
}

func (s *SpelNodeImpl) getLeftOperand() SpelNode {
	return s.Children[0]
}

func (s *SpelNodeImpl) getRightOperand() SpelNode {
	return s.Children[1]
}

func (s *SpelNodeImpl) toDescriptorFromObject(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func (s *SpelNodeImpl) equalityCheck(left interface{}, right interface{}) bool {
	if s.toDescriptorFromObject(left) == s.toDescriptorFromObject(right) {
		if left == right {
			return true
		}
	}
	return false
}

func checkType(left interface{}, right interface{}) bool {
	if reflect.TypeOf(left).String() != reflect.TypeOf(right).String() {
		return false
	}
	return true
}
