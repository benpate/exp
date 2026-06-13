package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// requirePredicate asserts that an Expression is a Predicate and returns it.
func requirePredicate(t *testing.T, exp Expression) Predicate {
	t.Helper()
	predicate, ok := exp.(Predicate)
	require.True(t, ok, "expected Predicate, got %T", exp)
	return predicate
}

// requireAnd asserts that an Expression is an AndExpression of the given length.
func requireAnd(t *testing.T, exp Expression, length int) AndExpression {
	t.Helper()
	and, ok := exp.(AndExpression)
	require.True(t, ok, "expected AndExpression, got %T", exp)
	require.Len(t, and, length)
	return and
}

// requireOr asserts that an Expression is an OrExpression of the given length.
func requireOr(t *testing.T, exp Expression, length int) OrExpression {
	t.Helper()
	or, ok := exp.(OrExpression)
	require.True(t, ok, "expected OrExpression, got %T", exp)
	require.Len(t, or, length)
	return or
}

// requireLast asserts that the last entry of an expression slice is a Predicate
// with the expected field, operator, and value.
func requireLast(t *testing.T, exp []Expression, field string, operator string, value any) {
	t.Helper()
	require.NotEmpty(t, exp)
	predicate := requirePredicate(t, exp[len(exp)-1])
	require.Equal(t, field, predicate.Field)
	require.Equal(t, operator, predicate.Operator)
	require.Equal(t, value, predicate.Value)
}

// boolMatcher returns a MatcherFunc that reports each predicate's boolean Value,
// along with a pointer to a counter recording how many times it was invoked.
// The counter lets tests confirm short-circuit behavior.
func boolMatcher() (MatcherFunc, *int) {
	calls := 0
	fn := func(predicate Predicate) bool {
		calls++
		match, _ := predicate.Value.(bool)
		return match
	}
	return fn, &calls
}
