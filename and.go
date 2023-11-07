package exp

// AndExpression combines a series of sub-expressions using AND logic
type AndExpression []Expression

// And combines one or more expression parameters into an AndExpression
func And(expressions ...Expression) AndExpression {

	var result Expression
	result = make(AndExpression, 0)

	for _, expression := range expressions {
		result = result.And(expression)
	}

	return result.(AndExpression)
}

func (e AndExpression) And(exp Expression) Expression {

	switch value := exp.(type) {
	case EmptyExpression:
		return e
	case AndExpression:
		return append(e, value...)
	default:
		return append(e, value)
	}
}

func (e AndExpression) Or(exp Expression) Expression {

	if _, ok := exp.(EmptyExpression); ok {
		return e
	}
	return Or(e, exp)
}

func (e AndExpression) AndEqual(name string, value any) Expression {
	return e.And(New(name, OperatorEqual, value))
}

func (e AndExpression) AndNotEqual(name string, value any) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

func (e AndExpression) AndLessThan(name string, value any) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

func (e AndExpression) AndLessOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

func (e AndExpression) AndGreaterThan(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

func (e AndExpression) AndGreaterOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

func (e AndExpression) AndIn(name string, value any) Expression {
	return e.And(New(name, OperatorIn, value))
}

func (e AndExpression) AndNotIn(name string, value any) Expression {
	return e.And(New(name, OperatorNotIn, value))
}

// Match implements the Expression interface.  It loops through all sub-expressions and returns TRUE if all of them match
func (e AndExpression) Match(fn MatcherFunc) bool {

	for _, expression := range e {

		if !expression.Match(fn) {
			return false
		}
	}

	return true
}

func (e AndExpression) IsEmpty() bool {
	return len(e) == 0
}

func (e AndExpression) NotEmpty() bool {
	return len(e) > 0
}
