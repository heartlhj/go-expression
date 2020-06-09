package expression

type SpelExpressionParser struct {
	*TemplateAwareExpressionParser
}

func (s *TemplateAwareExpressionParser) DoParseExpression(expressionString string) Expression {
	parser := InternalSpelExpressionParser{}
	return parser.DoParseExpression(expressionString)
}
