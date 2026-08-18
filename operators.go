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

// OperatorInAll represents an "in all" comparison, when used in Predicates and Criteria.
const OperatorInAll = "IN ALL"

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

// OperatorExists represents an "exists" comparison, which should return TRUE if the provided field exists.
const OperatorExists = "EXISTS"

// OperatorGeoWithin represents a geometric search within a given shape.
const OperatorGeoWithin = "GEO-WITHIN"

// OperatorGeoIntersects represents a geometric search that intersects with a given shape
const OperatorGeoIntersects = "GEO-INTERSECTS"

// Operator tries to convert non-standard values into standard operators.
func Operator(value string) string {
	result, _ := OperatorOk(value)
	return result
}

// OperatorOk converts a non-standard value into a standard operator constant,
// returning the standardized operator and TRUE when the value is recognized.
func OperatorOk(value string) (string, bool) {

	// Matching is case-insensitive.  The HTML entity spellings ("&gt;", "&le;")
	// are accepted so that operators survive a trip through a URL or an HTML
	// attribute without being mangled.
	value = strings.ToUpper(value)

	switch value {

	case OperatorGreaterThan, "GT", "&GT;":
		return OperatorGreaterThan, true

	case OperatorGreaterOrEqual, "GTE", "GE", "&GE;":
		return OperatorGreaterOrEqual, true

	case OperatorEqual, "EQ", "IS", "==", "&EQUALS;":
		return OperatorEqual, true

	case OperatorNotEqual, "NEQ", "NE", "&NE;":
		return OperatorNotEqual, true

	case OperatorLessOrEqual, "LTE", "LE", "&LE;":
		return OperatorLessOrEqual, true

	case OperatorLessThan, "LT", "&LT;":
		return OperatorLessThan, true

	case OperatorIn:
		return OperatorIn, true

	case OperatorNotIn:
		return OperatorNotIn, true

	case OperatorInAll:
		return OperatorInAll, true

	case OperatorBeginsWith:
		return OperatorBeginsWith, true

	case OperatorEndsWith:
		return OperatorEndsWith, true

	case OperatorContains:
		return OperatorContains, true

	case OperatorContainedBy:
		return OperatorContainedBy, true

	case OperatorExists:
		return OperatorExists, true

	case OperatorGeoWithin:
		return OperatorGeoWithin, true

	case OperatorGeoIntersects:
		return OperatorGeoIntersects, true

	default:

		// RULE: An unrecognized value falls back to EQUAL, so that a caller who
		// ignores the boolean still receives a usable operator.
		return OperatorEqual, false
	}
}
