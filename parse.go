package exp

import (
	"strings"
)

// Parse converts a "field operator value" string into an Expression.
// Malformed input that does not contain both a field and an operator
// returns an EmptyExpression instead of panicking.
func Parse(value string) Expression {

	// Extract the field
	field, tail, ok := strings.Cut(value, " ")
	if !ok {
		return Empty()
	}

	// Split the operator and value
	operator, value, ok := strings.Cut(tail, " ")
	if !ok {
		return Empty()
	}

	// Parse the operator into a recognized token
	operator, ok = OperatorOk(operator)
	if !ok {
		return Empty()
	}

	// Success
	return Predicate{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}
