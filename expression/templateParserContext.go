package expression

type TemplateParserContext struct {
	expressionPrefix string

	expressionSuffix string
}

func (t *TemplateParserContext) isTemplate() bool {
	return true
}

func (t *TemplateParserContext) getExpressionPrefix() string {
	return t.expressionPrefix
}

func (t *TemplateParserContext) getExpressionSuffix() string {
	return t.expressionSuffix
}
