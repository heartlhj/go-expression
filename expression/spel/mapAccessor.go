package spel

import "reflect"

type MapAccessor struct {
}

func (this MapAccessor) CanRead(context EvaluationContext, target interface{}, name string) bool {
	m, ok := target.(map[string]interface{})
	return ok && m[name] != nil
}

func (this MapAccessor) Read(context EvaluationContext, target interface{}, name string) TypedValue {
	m, ok := target.(map[string]interface{})
	if !ok {
		panic("Target must be of type Map")
	}
	value := m[name]
	if value == nil {
		panic("Map does not contain a value for key")
	}
	return TypedValue{Value: value}
}

func (this MapAccessor) GetSpecificTargetClasses() interface{} {
	return reflect.Map
}
