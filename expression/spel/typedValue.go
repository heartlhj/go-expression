package spel

import "reflect"

type TypedValue struct {
	Value interface{}

	typeDescriptor TypeDescriptor
}

func (this TypedValue) GetTypeDescriptor() TypeDescriptor {
	if reflect.DeepEqual(this.typeDescriptor, TypeDescriptor{}) && this.Value == nil {
		object, _ := TypeDescriptor{}.ForObject(this.Value)
		this.typeDescriptor = object
	}
	return this.typeDescriptor
}
