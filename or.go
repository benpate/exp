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

// Or is a part of the Expression interface.
// It combines another expression into a new OrExpression
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

// And is a part of the Expression interface
// It combines this OrExpression with another expression into a new AndExpression
func (e OrExpression) And(exp Expression) Expression {

	if _, ok := exp.(EmptyExpression); ok {
		return e
	}

	return And(e, exp)
}

// AndEqual is a part of the Expression interface.
// It creates a new AndExpression using the Equal comparison
func (e OrExpression) AndEqual(name string, value any) Expression {
	return e.And(New(name, OperatorEqual, value))
}

// AndNotEqual is a part of the Expression interface.
// It creates a new AndExpression using the NotEqual comparison
func (e OrExpression) AndNotEqual(name string, value any) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

// AndLessThan is a part of the Expression interface.
// It creates a new AndExpression using the LessThan comparison
func (e OrExpression) AndLessThan(name string, value any) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

// AndLessOrEqual is a part of the Expression interface.
// It creates a new AndExpression using the LessOrEqual comparison
func (e OrExpression) AndLessOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

// AndGreaterThan is a part of the Expression interface.
// It creates a new AndExpression using the GreaterThan comparison
func (e OrExpression) AndGreaterThan(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

// AndGreaterOrEqual is a part of the Expression interface.
// It creates a new AndExpression using the GreaterOrEqual comparison
func (e OrExpression) AndGreaterOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

// AndIn is a part of the Expression interface.
// It creates a new AndExpression using the In comparison
func (e OrExpression) AndIn(name string, value any) Expression {
	return e.And(New(name, OperatorIn, value))
}

// AndNotIn is a part of the Expression interface.
// It creates a new AndExpression using the NotIn comparison
func (e OrExpression) AndNotIn(name string, value any) Expression {
	return e.And(New(name, OperatorNotIn, value))
}

// AndInAll is a part of the Expression interface.
// It creates a new AndExpression using the InAll comparison
func (e OrExpression) AndInAll(field string, values ...any) Expression {
	return e.And(New(field, OperatorInAll, values))
}

// Match is a part of the Expression interface.
// It loops through all sub-expressions and returns TRUE if any of them match
func (e OrExpression) Match(fn MatcherFunc) bool {

	for _, expression := range e {
		if expression.Match(fn) {
			return true
		}
	}

	return false
}

// IsEmpty is a part of the Expression interface.
// It returns TRUE if an expression does not have any predicates
func (e OrExpression) IsEmpty() bool {
	return len(e) == 0
}

// NotEmpty is a part of the Expression interface.
// It returns TRUE if an expression has one or more predicates
func (e OrExpression) NotEmpty() bool {
	return len(e) > 0
}

// Fields is a part of the Expression interface.
// It returns a slice of field names that are used in this expression.
func (e OrExpression) Fields() []string {
	var result []string

	for _, expression := range e {
		result = append(result, expression.Fields()...)
	}

	return result
}
