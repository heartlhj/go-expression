package expression

type ExpressionParser interface {
	ParseExpression(var1 string) Expression

	DoParseExpression(var1 string) Expression
}
