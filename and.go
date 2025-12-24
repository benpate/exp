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

// And is a part of the Expression interface.
// It combines another expression into a new AndExpression
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

// Or is a part of the Expression interface.
// It combines this AndExpression with another expression into a new OrExpression
func (e AndExpression) Or(exp Expression) Expression {

	if _, ok := exp.(EmptyExpression); ok {
		return e
	}
	return Or(e, exp)
}

// AndEqual is a part of the Expression interface.
// It creates a new AndExpression using the Equal comparison
func (e AndExpression) AndEqual(name string, value any) Expression {
	return e.And(New(name, OperatorEqual, value))
}

// AndNotEqual is a part of the Expression interface.
// It creates a new AndExpression using the NotEqual comparison
func (e AndExpression) AndNotEqual(name string, value any) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

// AndLessThan is a part of the Expression interface.
// It creates a new AndExpression using the LessThan comparison
func (e AndExpression) AndLessThan(name string, value any) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

// AndLessOrEqual is a part of the Expression interface.
// It creates a new AndExpression using the LessOrEqual comparison
func (e AndExpression) AndLessOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

// AndGreaterThan is a part of the Expression interface.
// It creates a new AndExpression using the GreaterThan comparison
func (e AndExpression) AndGreaterThan(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

// AndGreaterOrEqual is a part of the Expression interface.
// It creates a new AndExpression using the GreaterOrEqual comparison
func (e AndExpression) AndGreaterOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

// AndInAll is a part of the Expression interface.
// It creates a new AndExpression using the InAll comparison
func (e AndExpression) AndInAll(field string, values ...any) Expression {
	return e.And(New(field, OperatorInAll, values))
}

// AndIn is a part of the Expression interface.
// It creates a new AndExpression using the In comparison
func (e AndExpression) AndIn(name string, value any) Expression {
	return e.And(New(name, OperatorIn, value))
}

// AndNotIn is a part of the Expression interface.
// It creates a new AndExpression using the NotIn comparison
func (e AndExpression) AndNotIn(name string, value any) Expression {
	return e.And(New(name, OperatorNotIn, value))
}

// Match is a part of the Expression interface.
// It loops through all sub-expressions and returns TRUE if all of them match
func (e AndExpression) Match(fn MatcherFunc) bool {

	for _, expression := range e {

		if !expression.Match(fn) {
			return false
		}
	}

	return true
}

// IsEmpty is a part of the Expression interface.
// It returns TRUE if an expression does not have any predicates
func (e AndExpression) IsEmpty() bool {
	return len(e) == 0
}

// NotEmpty is a part of the Expression interface.
// It returns TRUE if an expression has one or more predicates
func (e AndExpression) NotEmpty() bool {
	return len(e) > 0
}

// Fields is a part of the Expression interface.
// It returns a slice of field names that are used in this expression.
func (e AndExpression) Fields() []string {
	fields := make([]string, 0)

	for _, expression := range e {
		fields = append(fields, expression.Fields()...)
	}

	return fields
}
