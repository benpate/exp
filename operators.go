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

// OperatorBeginsWith represents a "begins with" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorBeginsWith = "BEGINS"

// OperatorEndsWith represents a "ends with" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorEndsWith = "ENDS"

// OperatorContains represents a "contains" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorContains = "CONTAINS"

// OperatorContainedBy represents a "contained by" comparison, when used in Predicates and Criteria.  It is only valid for string values.
const OperatorContainedBy = "CONTAINED BY"

// Operator tries to convert non-standard values into standard operators
func Operator(value string) string {

	value = strings.ToUpper(value)

	switch value {

	case OperatorGreaterThan, "GT":
		return OperatorGreaterThan

	case OperatorGreaterOrEqual, "GE":
		return OperatorGreaterOrEqual

	case OperatorEqual, "EQ":
		return OperatorEqual

	case OperatorNotEqual, "NE":
		return OperatorNotEqual

	case OperatorLessOrEqual, "LE":
		return OperatorLessOrEqual

	case OperatorLessThan, "LT":
		return OperatorLessThan

	case OperatorBeginsWith:
		return OperatorBeginsWith

	case OperatorEndsWith:
		return OperatorEndsWith

	case OperatorContains:
		return OperatorContains

	case OperatorContainedBy:
		return OperatorContainedBy

	default:
		return OperatorEqual
	}
}
