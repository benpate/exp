package exp

// EmptyExpression is an implementation of the Expression interface that represents an empty expression.
type EmptyExpression struct{}

// Empty creates and returns a new EmptyExpression.
func Empty() EmptyExpression {
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
// It returns an Equal predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndEqual(name string, value any) Expression {
	return New(name, OperatorEqual, value)
}

// AndNotEqual is a part of the Expression interface.
// It returns a NotEqual predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndNotEqual(name string, value any) Expression {
	return New(name, OperatorNotEqual, value)
}

// AndLessThan is a part of the Expression interface.
// It returns a LessThan predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndLessThan(name string, value any) Expression {
	return New(name, OperatorLessThan, value)
}

// AndLessOrEqual is a part of the Expression interface.
// It returns a LessOrEqual predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndLessOrEqual(name string, value any) Expression {
	return New(name, OperatorLessOrEqual, value)
}

// AndGreaterThan is a part of the Expression interface.
// It returns a GreaterThan predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndGreaterThan(name string, value any) Expression {
	return New(name, OperatorGreaterThan, value)
}

// AndGreaterOrEqual is a part of the Expression interface.
// It returns a GreaterOrEqual predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndGreaterOrEqual(name string, value any) Expression {
	return New(name, OperatorGreaterOrEqual, value)
}

// AndIn is a part of the Expression interface.
// It returns an In predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndIn(name string, value any) Expression {
	return New(name, OperatorIn, value)
}

// AndNotIn is a part of the Expression interface.
// It returns a NotIn predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndNotIn(name string, value any) Expression {
	return New(name, OperatorNotIn, value)
}

// AndInAll is a part of the Expression interface.
// It returns an InAll predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) AndInAll(name string, value ...any) Expression {
	return New(name, OperatorInAll, value)
}

// OrEqual is a part of the Expression interface.
// It returns an Equal predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrEqual(name string, value any) Expression {
	return New(name, OperatorEqual, value)
}

// OrNotEqual is a part of the Expression interface.
// It returns a NotEqual predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrNotEqual(name string, value any) Expression {
	return New(name, OperatorNotEqual, value)
}

// OrLessThan is a part of the Expression interface.
// It returns a LessThan predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrLessThan(name string, value any) Expression {
	return New(name, OperatorLessThan, value)
}

// OrLessOrEqual is a part of the Expression interface.
// It returns a LessOrEqual predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrLessOrEqual(name string, value any) Expression {
	return New(name, OperatorLessOrEqual, value)
}

// OrGreaterThan is a part of the Expression interface.
// It returns a GreaterThan predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrGreaterThan(name string, value any) Expression {
	return New(name, OperatorGreaterThan, value)
}

// OrGreaterOrEqual is a part of the Expression interface.
// It returns a GreaterOrEqual predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrGreaterOrEqual(name string, value any) Expression {
	return New(name, OperatorGreaterOrEqual, value)
}

// OrIn is a part of the Expression interface.
// It returns an In predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrIn(name string, value any) Expression {
	return New(name, OperatorIn, value)
}

// OrNotIn is a part of the Expression interface.
// It returns a NotIn predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrNotIn(name string, value any) Expression {
	return New(name, OperatorNotIn, value)
}

// OrInAll is a part of the Expression interface.
// It returns an InAll predicate, since combining with an empty expression has no effect.
func (e EmptyExpression) OrInAll(name string, value ...any) Expression {
	return New(name, OperatorInAll, value)
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
