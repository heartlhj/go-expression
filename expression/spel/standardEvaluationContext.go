package spel

//map存储
type StandardEvaluationContext struct {
	Variables map[string]interface{}

	propertyAccessors []PropertyAccessor
}

func (this *StandardEvaluationContext) AddPropertyAccessor(resolver PropertyAccessor) {
	this.AddBeforeDefault(this.initPropertyAccessors(), resolver)
}

func (this *StandardEvaluationContext) AddBeforeDefault(resolvers []PropertyAccessor, resolver PropertyAccessor) {
	resolvers = append(resolvers, resolver)
	this.propertyAccessors = resolvers
}
func (this *StandardEvaluationContext) SetVariable(var1 string, var2 map[string]interface{}) {
	this.Variables[var1] = var2
}

func (this *StandardEvaluationContext) SetVariables(var2 map[string]interface{}) {
	this.Variables = var2
}

func (this StandardEvaluationContext) LookupVariable(name string) interface{} {
	return this.Variables[name]
}

func (this StandardEvaluationContext) GetPropertyAccessors() []PropertyAccessor {
	return this.initPropertyAccessors()
}

func (this StandardEvaluationContext) initPropertyAccessors() []PropertyAccessor {
	accessors := this.propertyAccessors
	if accessors == nil {
		accessors = make([]PropertyAccessor, 1)
		accessors[0] = ReflectivePropertyAccessor{}
		this.propertyAccessors = accessors

	}
	return accessors
}
