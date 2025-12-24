package exp

// EmptyExpression is an implementation of the Expression interface that represents an empty expression.
type EmptyExpression struct{}

// Empty creates and returns a new EmptyExpression.
func Empty() Expression {
	return EmptyExpression{}
}

// And is a part of the Expression interface.
// It returns the other expression since combining with an empty expression has no effect.
func (e EmptyExpression) And(exp Expression) Expression {
	return exp
}

// Or is a part of the Expression interface.
// It returns the other expression since combining with an empty expression has no effect.
func (e EmptyExpression) Or(exp Expression) Expression {
	return exp
}

// Match is a part of the Expression interface.
// It always returns TRUE for an EmptyExpression, since it matches everything.
func (e EmptyExpression) Match(_ MatcherFunc) bool {
	return true
}

// AndEqual is a part of the Expression interface.
// It creates a new AndExpression using the Equal comparison
func (e EmptyExpression) AndEqual(name string, value any) Expression {
	return e.And(New(name, OperatorEqual, value))
}

// AndNotEqual is a part of the Expression interface.
// It creates a new AndExpression using the NotEqual comparison
func (e EmptyExpression) AndNotEqual(name string, value any) Expression {
	return e.And(New(name, OperatorNotEqual, value))
}

// AndLessThan is a part of the Expression interface.
// It creates a new AndExpression using the LessThan comparison
func (e EmptyExpression) AndLessThan(name string, value any) Expression {
	return e.And(New(name, OperatorLessThan, value))
}

// AndLessOrEqual is a part of the Expression interface.
// It creates a new AndExpression using the LessOrEqual comparison
func (e EmptyExpression) AndLessOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorLessOrEqual, value))
}

// AndGreaterThan is a part of the Expression interface.
// It creates a new AndExpression using the GreaterThan comparison
func (e EmptyExpression) AndGreaterThan(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterThan, value))
}

// AndGreaterOrEqual is a part of the Expression interface.
// It creates a new AndExpression using the GreaterOrEqual comparison
func (e EmptyExpression) AndGreaterOrEqual(name string, value any) Expression {
	return e.And(New(name, OperatorGreaterOrEqual, value))
}

// AndIn is a part of the Expression interface.
// It creates a new AndExpression using the In comparison
func (e EmptyExpression) AndIn(name string, value any) Expression {
	return e.And(New(name, OperatorIn, value))
}

// AndNotIn is a part of the Expression interface.
// It creates a new AndExpression using the NotIn comparison
func (e EmptyExpression) AndNotIn(name string, value any) Expression {
	return e.And(New(name, OperatorNotIn, value))
}

// AndInAll is a part of the Expression interface.
// It creates a new AndExpression using the InAll comparison
func (e EmptyExpression) AndInAll(name string, value ...any) Expression {
	return e.And(New(name, OperatorInAll, value))
}

// IsEmpty is a part of the Expression interface.
// It always returns TRUE for an EmptyExpression, since it is always empty.
func (e EmptyExpression) IsEmpty() bool {
	return true
}

// NotEmpty is a part of the Expression interface.
// It always returns FALSE for an EmptyExpression, since it is always empty.
func (e EmptyExpression) NotEmpty() bool {
	return false
}

// Fields is a part of the Expression interface.
// It returns an empty slice since there are no fields in an EmptyExpression.
func (e EmptyExpression) Fields() []string {
	return make([]string, 0)
}
