package ast

import (
	. "github.com/heartlhj/go-expression/expression/spel"
	"reflect"
)

type IndexedType string

//多维数组
type Indexer struct {
	*SpelNodeImpl
	cachedReadName     string
	cachedReadAccessor PropertyAccessor
	indexedType        reflect.Kind
}

func (this Indexer) GetValueRef(state ExpressionState) ValueRef {
	context := state.GetActiveContextObject()
	target := context.Value
	targetDescriptor := context.GetTypeDescriptor()
	var indexValue TypedValue
	var index interface{}
	_, ok := target.(map[interface{}]interface{})
	reference, isOK := this.Children[0].(PropertyOrFieldReference)
	if ok && isOK {
		index = reference.Name
		indexValue = TypedValue{Value: index}
	} else {
		defer state.PopActiveContextObject()
		state.PushActiveContextObject(state.GetRootContextObject())
		indexValue = this.Children[0].GetValueInternal(state)
		index = indexValue.Value
		if index == nil {
			panic("No index")
		}
	}

	if target == nil {
		panic("Cannot index into a null value")
	}
	kind := reflect.TypeOf(target).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		this.indexedType = kind
		index, _ := index.(int)
		return ArrayIndexingValueRef{Array: target, Index: index, TypeDescriptor: targetDescriptor}
	} else if kind == reflect.String {
		this.indexedType = kind
		index, _ := index.(int)
		return StringIndexingLValue{Target: target, Index: index, TypeDescriptor: targetDescriptor}
	}
	return nil
}

func (this Indexer) GetValueInternal(state ExpressionState) TypedValue {
	return this.GetValueRef(state).GetValue()
}

func (this *Indexer) setValue(state ExpressionState) TypedValue {
	return this.GetValueRef(state).GetValue()
}

type ArrayIndexingValueRef struct {
	Array interface{}

	Index int

	TypeDescriptor TypeDescriptor
}

func (this ArrayIndexingValueRef) GetValue() TypedValue {
	arry := reflect.ValueOf(this.Array)
	//获取下标为Index的数据
	len := arry.Len()
	if this.Index >= len {
		panic("The index is invalid")
	}
	value := arry.Index(this.Index)
	return TypedValue{Value: value}
}

type StringIndexingLValue struct {
	Target interface{}

	Index int

	TypeDescriptor TypeDescriptor
}

func (this StringIndexingLValue) GetValue() TypedValue {
	target, _ := this.Target.(string)
	if this.Index >= len(target) {
		panic("The index is invalid")
	}
	//获取下标为Index的数据
	value := string(target[this.Index])
	return TypedValue{Value: value}
}
