package spel

import "reflect"

var (
	CACHED_COMMON_TYPES = []reflect.Kind{
		reflect.Bool, reflect.Int, reflect.Interface, reflect.Int8, reflect.Float64, reflect.Func, reflect.String,
		reflect.Float32,
	}
)

type TypeDescriptor struct {
	commonTypesCache map[reflect.Kind]TypeDescriptor
	resolvableType   ResolvableType
}

func (this *TypeDescriptor) initTypeDescriptor() {
	for _, preCachedClass := range CACHED_COMMON_TYPES {
		this.commonTypesCache[preCachedClass] = this.valueOf(preCachedClass)
	}
}

func (this TypeDescriptor) ForObject(source interface{}) (TypeDescriptor, error) {
	if source != nil {
		return this.valueOf((source.(reflect.Type)).Kind()), nil
	}
	return TypeDescriptor{}, nil
}

func (this TypeDescriptor) valueOf(objType reflect.Kind) TypeDescriptor {
	descriptor := this.commonTypesCache[objType]
	if reflect.DeepEqual(descriptor, TypeDescriptor{}) {
		return TypeDescriptor{resolvableType: ResolvableType{objType}}
	}
	return descriptor
}
