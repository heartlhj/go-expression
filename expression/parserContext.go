package expression

type ParserContext interface {
	isTemplate() bool

	getExpressionPrefix() string

	getExpressionSuffix() string
}
