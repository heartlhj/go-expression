package spel

import "container/list"

//根据KEY获取MAP的value
type ExpressionState struct {
	RelatedContext EvaluationContext

	RootObject TypedValue

	ContextObjects list.List

	VariableScopes list.List
}

func (this *ExpressionState) LookupVariable(name string) TypedValue {
	variable := this.RelatedContext.LookupVariable(name)
	return TypedValue{Value: variable}
}

func (this *ExpressionState) PopActiveContextObjectNull() {
	front := this.ContextObjects.Front()
	this.ContextObjects.Remove(front)
}

func (this *ExpressionState) PushActiveContextObject(obj TypedValue) {
	this.ContextObjects.PushFront(obj)
}

func (this *ExpressionState) PopActiveContextObject() {
	front := this.ContextObjects.Front()
	this.ContextObjects.Remove(front)
}

func (this *ExpressionState) GetActiveContextObject() TypedValue {
	return this.ContextObjects.Front().Value.(TypedValue)
}

func (this *ExpressionState) GetEvaluationContext() EvaluationContext {
	return this.RelatedContext
}

func (this *ExpressionState) GetRootContextObject() TypedValue {
	return this.RootObject
}
