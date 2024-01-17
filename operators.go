package exp

import "strings"

// OperatorGreaterThan represents an "greater than" comparison, when used in Predicates and Criteria
const OperatorGreaterThan = ">"

// OperatorGreaterOrEqual represents an "greater or equal" comparison, when used in Predicates and Criteria
const OperatorGreaterOrEqual = ">="

// OperatorEqual represents an "equals" comparison, when used in Predicates and Criteria
const OperatorEqual = "="

// OperatorNotEqual represents a "not equals" comparison, when used in Predicates and Criteria
const OperatorNotEqual = "!="

// OperatorLessOrEqual represents an "less or equal" comparison, when used in Predicates and Criteria
const OperatorLessOrEqual = "<="

// OperatorLessThan represents a "less than" comparison, when used in Predicates and Criteria
const OperatorLessThan = "<"

// OperatorIn represents a "in" comparison, when used in Predicates and Criteria.
const OperatorIn = "IN"

// OperatorNotIn represents a "not in" comparison, when used in Predicates and Criteria.
const OperatorNotIn = "NOT IN"

// OperatorBeginsWith represents a "begins with" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorBeginsWith = "BEGINS"

// OperatorEndsWith represents a "ends with" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorEndsWith = "ENDS"

// OperatorContains represents a "contains" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorContains = "CONTAINS"

// OperatorContainedBy represents a "contained by" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorContainedBy = "CONTAINED BY"

func Operator(value string) string {
	result, _ := OperatorOk(value)
	return result
}

// OperatorOk tries to convert non-standard values into standard operators.
// If a match is found, then it returns the standardized value and TRUE.
// If a match is not found, then the default EQUAL is returned along with a FALSE.
func OperatorOk(value string) (string, bool) {

	value = strings.ToUpper(value)

	switch value {

	case OperatorGreaterThan, "GT":
		return OperatorGreaterThan, true

	case OperatorGreaterOrEqual, "GE":
		return OperatorGreaterOrEqual, true

	case OperatorEqual, "EQ":
		return OperatorEqual, true

	case OperatorNotEqual, "NE":
		return OperatorNotEqual, true

	case OperatorLessOrEqual, "LE":
		return OperatorLessOrEqual, true

	case OperatorLessThan, "LT":
		return OperatorLessThan, true

	case OperatorIn:
		return OperatorIn, true

	case OperatorNotIn:
		return OperatorNotIn, true

	case OperatorBeginsWith:
		return OperatorBeginsWith, true

	case OperatorEndsWith:
		return OperatorEndsWith, true

	case OperatorContains:
		return OperatorContains, true

	case OperatorContainedBy:
		return OperatorContainedBy, true

	default:
		return OperatorEqual, false
	}
}
