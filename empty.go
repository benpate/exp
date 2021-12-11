package exp

type EmptyExpression struct{}

func Empty() Expression {
	return EmptyExpression{}
}

func (e EmptyExpression) And(exp Expression) Expression {
	return exp
}

func (e EmptyExpression) Or(exp Expression) Expression {
	return exp
}

func (e EmptyExpression) Match(fn MatcherFunc) bool {
	return true
}
