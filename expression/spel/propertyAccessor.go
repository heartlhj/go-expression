package spel

//寄存器
type PropertyAccessor interface {
	CanRead(context EvaluationContext, target interface{}, name string) bool

	Read(context EvaluationContext, target interface{}, name string) TypedValue

	GetSpecificTargetClasses() interface{}
}
