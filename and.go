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

func (andExpression AndExpression) And(exp Expression) Expression {

	switch value := exp.(type) {
	case EmptyExpression:
		return andExpression
	case AndExpression:
		return append(andExpression, value...)
	case Predicate:
		return append(andExpression, value)
	default:
		return Or(andExpression, value)
	}
}

func (andExpression AndExpression) Or(exp Expression) Expression {

	if _, ok := exp.(EmptyExpression); ok {
		return andExpression
	}
	return Or(andExpression, exp)
}

func (andExpression AndExpression) AndEqual(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorEqual, value))
}

func (andExpression AndExpression) AndNotEqual(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorNotEqual, value))
}

func (andExpression AndExpression) AndLessThan(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorLessThan, value))
}

func (andExpression AndExpression) AndLessOrEqual(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorLessOrEqual, value))
}

func (andExpression AndExpression) AndGreaterThan(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorGreaterThan, value))
}

func (andExpression AndExpression) AndGreaterOrEqual(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorGreaterOrEqual, value))
}

func (andExpression AndExpression) AndIn(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorIn, value))
}

func (andExpression AndExpression) AndNotIn(name string, value interface{}) Expression {
	return andExpression.And(New(name, OperatorNotIn, value))
}

// Match implements the Expression interface.  It loops through all sub-expressions and returns TRUE if all of them match
func (andExpression AndExpression) Match(fn MatcherFunc) bool {

	for _, expression := range andExpression {

		if !expression.Match(fn) {
			return false
		}
	}

	return true
}
