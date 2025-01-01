package exp

// OrExpression compares a series of sub-expressions, using the OR logic
type OrExpression []Expression

// Or combines one or more expression parameters into an OrExpression
func Or(expressions ...Expression) OrExpression {

	var result Expression
	result = make(OrExpression, 0)

	for _, expression := range expressions {
		result = result.Or(expression)
	}

	return result.(OrExpression)
}

// Or appends a new expression into this compound expression
func (e OrExpression) Or(exp Expression) Expression {

	switch value := exp.(type) {
	case EmptyExpression:
		return e
	case OrExpression:
		return append(e, value...)
	default:
		return append(e, value)
	}
}

// And returns a fully populated AndExpression
func (e OrExpression) And(exp Expression) Expression {

	if _, ok := exp.(EmptyExpression); ok {
		return e
	}

	return And(e, exp)
}

func (e OrExpression) AndEqual(name string, value any) Expression {
	return e.And(New(name, OperatorEqual, value))
}

func (e OrExpression) AndNotEqual(name string, value any) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

func (e OrExpression) AndLessThan(name string, value any) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

func (e OrExpression) AndLessOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

func (e OrExpression) AndGreaterThan(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

func (e OrExpression) AndGreaterOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

func (e OrExpression) AndIn(name string, value any) Expression {
	return e.And(New(name, OperatorIn, value))
}

func (e OrExpression) AndNotIn(name string, value any) Expression {
	return e.And(New(name, OperatorNotIn, value))
}

func (e OrExpression) AndInAll(field string, values ...any) Expression {
	return e.And(New(field, OperatorInAll, values))
}

// Match implements the Expression interface.  It loops through all sub-expressions and returns TRUE if any of them match
func (e OrExpression) Match(fn MatcherFunc) bool {

	for _, expression := range e {
		if expression.Match(fn) {
			return true
		}
	}

	return false
}

func (e OrExpression) IsEmpty() bool {
	return len(e) == 0
}

func (e OrExpression) NotEmpty() bool {
	return len(e) > 0
}
