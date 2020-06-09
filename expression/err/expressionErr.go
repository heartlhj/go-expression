package err

type ExpressionErr struct {
	Code string
	Msg  string
}

func (e ExpressionErr) Error() string {
	return e.Code + e.Msg
}
