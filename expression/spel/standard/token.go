package standard

type Token struct {
	Kind     TokenKind
	Data     string
	StartPos int
	EndPos   int
}

func (t Token) isNumericRelationalOperator() bool {
	return true
}

func (t Token) StringValue() string {
	if t.Data != "" {
		return t.Data
	}
	return ""
}

func (t Token) IsNumericRelationalOperator() bool {
	return t.Kind.TokenKindType == GT || t.Kind.TokenKindType == GE || t.Kind.TokenKindType == LT ||
		t.Kind.TokenKindType == LE || t.Kind.TokenKindType == EQ || t.Kind.TokenKindType == NE
}
