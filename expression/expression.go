package expression

import . "github.com/heartlhj/go-expression/expression/spel"

//获取值
type Expression interface {
	GetExpressionString() string

	GetValue() interface{}

	GetValueContext(context EvaluationContext) interface{}
}
