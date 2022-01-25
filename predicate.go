package exp

// Predicate represents a single true/false comparison
type Predicate struct {
	Field    string
	Operator string
	Value    interface{}
}

// New returns a fully populated Predicate
func New(field string, operator string, value interface{}) Predicate {
	return Predicate{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}

// Equal creates a new Predicate using an "Equals" comparison
func Equal(field string, value interface{}) Predicate {
	return New(field, OperatorEqual, value)
}

// NotEqual creates a new Predicate using an "Not Equals" comparison
func NotEqual(field string, value interface{}) Predicate {
	return New(field, OperatorNotEqual, value)
}

// LessThan creates a new Predicate using an "Less Than" comparison
func LessThan(field string, value interface{}) Predicate {
	return New(field, OperatorLessThan, value)
}

// LessOrEqual creates a new Predicate using an "Less Or Equal" comparison
func LessOrEqual(field string, value interface{}) Predicate {
	return New(field, OperatorLessOrEqual, value)
}

// GreaterThan creates a new Predicate using an "Greater Than" comparison
func GreaterThan(field string, value interface{}) Predicate {
	return New(field, OperatorGreaterThan, value)
}

// GreaterOrEqual creates a new Predicate using an "Greater Or Equal" comparison
func GreaterOrEqual(field string, value interface{}) Predicate {
	return New(field, OperatorGreaterOrEqual, value)
}

// In creates a new Predicate using an "in" comparison
func In(field string, value interface{}) Predicate {
	return New(field, OperatorIn, value)
}

// Contains creates a new Predicate using an "Contains" comparison
func Contains(field string, value interface{}) Predicate {
	return New(field, OperatorContains, value)
}

// ContainedBy creates a new Predicate using an "ContainedBy" comparison
func ContainedBy(field string, value interface{}) Predicate {
	return New(field, OperatorContainedBy, value)
}

// BeginsWith creates a new Predicate using an "BeginsWith" comparison
func BeginsWith(field string, value interface{}) Predicate {
	return New(field, OperatorBeginsWith, value)
}

// EndsWith creates a new Predicate using an "EndsWith" comparison
func EndsWith(field string, value interface{}) Predicate {
	return New(field, OperatorEndsWith, value)
}

// And combines this predicate with another pre-existing expression into a new And expression
func (predicate Predicate) And(exp Expression) Expression {

	// Skip EmptyExpressions
	if _, ok := exp.(EmptyExpression); ok {
		return predicate
	}

	return AndExpression{predicate, exp}
}

// AndEqual combines this predicate with another one (created from the arguments) into an AndExpression
func (predicate Predicate) AndEqual(name string, value interface{}) Expression {
	return predicate.And(New(name, OperatorEqual, value))
}

// AndNotEqual combines this predicate with another one (created from the arguments) into an Expression
func (predicate Predicate) AndNotEqual(name string, value interface{}) Expression {
	return predicate.And(New(name, OperatorNotEqual, value))
}

// AndLessThan combines this predicate with another one (created from the arguments) into an Expression
func (predicate Predicate) AndLessThan(name string, value interface{}) Expression {
	return predicate.And(New(name, OperatorLessThan, value))
}

// AndLessOrEqual combines this predicate with another one (created from the arguments) into an Expression
func (predicate Predicate) AndLessOrEqual(name string, value interface{}) Expression {
	return predicate.And(New(name, OperatorLessOrEqual, value))
}

// AndGreaterThan combines this predicate with another one (created from the arguments) into an Expression
func (predicate Predicate) AndGreaterThan(name string, value interface{}) Expression {
	return predicate.And(New(name, OperatorGreaterThan, value))
}

// AndGreaterOrEqual combines this predicate with another one (created from the arguments) into an Expression
func (predicate Predicate) AndGreaterOrEqual(name string, value interface{}) Expression {
	return predicate.And(New(name, OperatorGreaterOrEqual, value))
}

// Or combines this predicate with another pre-existing expression into a new Or expression
func (predicate Predicate) Or(exp Expression) Expression {

	// Skip EmptyExpressions
	if _, ok := exp.(EmptyExpression); ok {
		return predicate
	}

	return OrExpression{predicate, exp}
}

// Match implements the Expression interface.  It uses a MatcherFunc to determine if this predicate matches an arbitrary dataset.
func (predicate Predicate) Match(fn MatcherFunc) bool {
	return fn(predicate)
}
