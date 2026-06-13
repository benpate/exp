package exp

import (
	"strings"
)

// Parse converts a "field operator value" string into an Expression.
// Malformed input that does not contain both a field and an operator
// returns an EmptyExpression instead of panicking.
func Parse(value string) Expression {

	field, rest, ok := strings.Cut(value, " ")
	if !ok {
		return Empty()
	}

	operator, val, ok := strings.Cut(rest, " ")
	if !ok {
		return Empty()
	}

	return Predicate{
		Field:    field,
		Operator: parseOperator(operator),
		Value:    val,
	}
}

func parseOperator(value string) string {
	switch value {

	case "eq", "is", "&equals;", "=", "==":
		return OperatorEqual

	case "ne", "&ne;", "!=":
		return OperatorNotEqual

	case "gt", "&gt;", ">":
		return OperatorGreaterThan

	case "lt", "&lt;", "<":
		return OperatorLessThan

	case "ge", "&ge;", ">=":
		return OperatorGreaterOrEqual

	case "le", "&le;", "<=":
		return OperatorLessOrEqual

	default:
		return ""
	}
}
