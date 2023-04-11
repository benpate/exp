package exp

import (
	"strings"
)

// Parse converts a string into an Expression
func Parse(value string) Expression {

	// TODO: LOW: This expression parser is too simple and brittle.  It should be expanded to support more complex expressions, and should better support malformed expressions.

	result := Predicate{}

	space := strings.IndexByte(value, ' ')
	result.Field = value[:space]
	value = value[space+1:]

	space = strings.IndexByte(value, ' ')
	result.Operator = parseOperator(value[:space])
	result.Value = value[space+1:]

	return result
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
