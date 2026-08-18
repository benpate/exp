package exp

import (
	"strings"
)

// Parse converts a "field operator value" string into an Expression.
func Parse(value string) Expression {

	// RULE: Input that is missing any of the three parts returns an
	// EmptyExpression rather than panicking, since this string usually arrives
	// from a URL query and cannot be trusted to be well formed.

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
