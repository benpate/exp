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

func (e EmptyExpression) AndEqual(name string, value any) Expression {
	return e.And(New(name, OperatorEqual, value))
}

func (e EmptyExpression) AndNotEqual(name string, value any) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

func (e EmptyExpression) AndLessThan(name string, value any) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

func (e EmptyExpression) AndLessOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

func (e EmptyExpression) AndGreaterThan(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

func (e EmptyExpression) AndGreaterOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

func (e EmptyExpression) AndIn(name string, value any) Expression {
	return e.And(New(name, OperatorIn, value))
}

func (e EmptyExpression) AndNotIn(name string, value any) Expression {
	return e.And(New(name, OperatorNotIn, value))
}

func (e EmptyExpression) AndInAll(name string, value ...any) Expression {
	return e.And(New(name, OperatorInAll, value))
}

func (e EmptyExpression) IsEmpty() bool {
	return true
}

func (e EmptyExpression) NotEmpty() bool {
	return false
}

func (e EmptyExpression) Fields() []string {
	return make([]string, 0)
}
