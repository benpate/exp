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
func (orExpression OrExpression) Or(exp Expression) Expression {

	switch value := exp.(type) {
	case EmptyExpression:
		return orExpression
	case OrExpression:
		return append(orExpression, value...)
	default:
		return append(orExpression, value)
	}
}

// And returns a fully populated AndExpression
func (orExpression OrExpression) And(exp Expression) Expression {

	if _, ok := exp.(EmptyExpression); ok {
		return orExpression
	}

	return And(orExpression, exp)
}

func (orExpression OrExpression) AndEqual(name string, value any) Expression {
	return orExpression.And(New(name, OperatorEqual, value))
}

func (orExpression OrExpression) AndNotEqual(name string, value any) Expression {
	return orExpression.And(New(name, OperatorNotEqual, value))
}

func (orExpression OrExpression) AndLessThan(name string, value any) Expression {
	return orExpression.And(New(name, OperatorLessThan, value))
}

func (orExpression OrExpression) AndLessOrEqual(name string, value any) Expression {
	return orExpression.And(New(name, OperatorLessOrEqual, value))
}

func (orExpression OrExpression) AndGreaterThan(name string, value any) Expression {
	return orExpression.And(New(name, OperatorGreaterThan, value))
}

func (orExpression OrExpression) AndGreaterOrEqual(name string, value any) Expression {
	return orExpression.And(New(name, OperatorGreaterOrEqual, value))
}

func (orExpression OrExpression) AndIn(name string, value any) Expression {
	return orExpression.And(New(name, OperatorIn, value))
}

func (orExpression OrExpression) AndNotIn(name string, value any) Expression {
	return orExpression.And(New(name, OperatorNotIn, value))
}

// Match implements the Expression interface.  It loops through all sub-expressions and returns TRUE if any of them match
func (orExpression OrExpression) Match(fn MatcherFunc) bool {

	for _, expression := range orExpression {
		if expression.Match(fn) {
			return true
		}
	}

	return false
}
