package spel

type ExpressionImpl struct {
}

func (e ExpressionImpl) GetExpressionString() string {
	return ""
}

func (e ExpressionImpl) GetValue() interface{} {
	return nil
}

func (e ExpressionImpl) GetValueContext(context EvaluationContext) interface{} {
	return nil
}
