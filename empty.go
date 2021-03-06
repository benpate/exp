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

func (e EmptyExpression) AndEqual(name string, value interface{}) Expression {
	return e.And(New(name, OperatorEqual, value))
}

func (e EmptyExpression) AndNotEqual(name string, value interface{}) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

func (e EmptyExpression) AndLessThan(name string, value interface{}) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

func (e EmptyExpression) AndLessOrEqual(name string, value interface{}) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

func (e EmptyExpression) AndGreaterThan(name string, value interface{}) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

func (e EmptyExpression) AndGreaterOrEqual(name string, value interface{}) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

func (e EmptyExpression) AndIn(name string, value interface{}) Expression {
	return e.And(New(name, OperatorIn, value))
}

func (e EmptyExpression) AndNotIn(name string, value interface{}) Expression {
	return e.And(New(name, OperatorNotIn, value))
}
