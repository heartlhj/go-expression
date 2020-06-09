package standard

type TokenKindType string

const (
	LPAREN TokenKindType = "("

	RPAREN TokenKindType = ")"

	COMMA TokenKindType = ","

	COLON TokenKindType = ":"

	HASH TokenKindType = "#"

	RSQUARE TokenKindType = "]"

	LSQUARE TokenKindType = "["

	LCURLY TokenKindType = "{"

	RCURLY TokenKindType = "}"

	DOT TokenKindType = "."

	PLUS TokenKindType = "+"

	STAR TokenKindType = "*"

	DIV     TokenKindType = "/"
	GE      TokenKindType = ">="
	GT      TokenKindType = ">"
	LE      TokenKindType = "<="
	LT      TokenKindType = "<"
	EQ      TokenKindType = "=="
	NE      TokenKindType = "!="
	PROJECT TokenKindType = "!["
	MOD     TokenKindType = "%"
	NOT     TokenKindType = "!"
	ASSIGN  TokenKindType = "="
	INC     TokenKindType = "++"
	DEC     TokenKindType = "--"

	MINUS  TokenKindType = "-"
	SELECT TokenKindType = "?["
	POWER  TokenKindType = "^"
	ELVIS  TokenKindType = "?:"

	SAFE_NAVI        TokenKindType = "?."
	BEAN_REF         TokenKindType = "@"
	FACTORY_BEAN_REF TokenKindType = "&"
	SYMBOLIC_OR      TokenKindType = "||"

	SYMBOLIC_AND TokenKindType = "&&"
	BETWEEN      TokenKindType = "between"

	SELECT_LAST TokenKindType = "$["

	IDENTIFIER TokenKindType = "IDENTIFIER"

	LITERAL_INT TokenKindType = "LITERAL_INT"

	LITERAL_LONG TokenKindType = "LITERAL_LONG"

	LITERAL_HEXINT TokenKindType = "LITERAL_HEXINT"

	LITERAL_HEXLONG TokenKindType = "LITERAL_HEXLONG"

	LITERAL_STRING TokenKindType = "LITERAL_STRING"

	LITERAL_REAL TokenKindType = "LITERAL_REAL"

	LITERAL_REAL_FLOAT TokenKindType = "LITERAL_REAL_FLOAT"
)

type TokenKind struct {
	TokenChars    []rune
	HasPayload    bool
	TokenKindType TokenKindType
}

func (t *TokenKind) InitTokenKind() *TokenKind {
	t.TokenChars = []rune(t.TokenKindType)
	t.HasPayload = len(t.TokenChars) == 0
	return t
}
