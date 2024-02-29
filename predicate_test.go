package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPredicates(t *testing.T) {

	// Simple Numeric Predicates
	require.Equal(t, Predicate{Field: "field", Operator: OperatorEqual, Value: 1}, New("field", "=", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorEqual, Value: 1}, Equal("field", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorNotEqual, Value: 1}, NotEqual("field", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorGreaterThan, Value: 1}, GreaterThan("field", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorGreaterOrEqual, Value: 1}, GreaterOrEqual("field", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorLessThan, Value: 1}, LessThan("field", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorLessOrEqual, Value: 1}, LessOrEqual("field", 1))

	// String-Based Predicates
	require.Equal(t, Predicate{Field: "field", Operator: OperatorEqual, Value: "John Connor"}, New("field", "=", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorEqual, Value: "John Connor"}, Equal("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorNotEqual, Value: "John Connor"}, NotEqual("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorGreaterThan, Value: "John Connor"}, GreaterThan("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorGreaterOrEqual, Value: "John Connor"}, GreaterOrEqual("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorLessThan, Value: "John Connor"}, LessThan("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorLessOrEqual, Value: "John Connor"}, LessOrEqual("field", "John Connor"))

	// Advanced String Predicates
	require.Equal(t, Predicate{Field: "field", Operator: OperatorBeginsWith, Value: "John Connor"}, BeginsWith("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorEndsWith, Value: "John Connor"}, EndsWith("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorContains, Value: "John Connor"}, Contains("field", "John Connor"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorContainedBy, Value: "John Connor"}, ContainedBy("field", "John Connor"))
}
