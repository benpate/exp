package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestEmptyState confirms the always-empty reporting of an EmptyExpression.
func TestEmptyState(t *testing.T) {
	t.Parallel()

	empty := Empty()

	require.IsType(t, EmptyExpression{}, empty)
	require.True(t, empty.IsEmpty())
	require.False(t, empty.NotEmpty())
	require.Equal(t, []string{}, empty.Fields())
}

// TestEmptyMatch confirms that an EmptyExpression matches everything, without
// ever invoking the MatcherFunc.
func TestEmptyMatch(t *testing.T) {
	t.Parallel()

	fn, calls := boolMatcher()

	require.True(t, Empty().Match(fn))
	require.Zero(t, *calls)
}

// TestEmptyAndOr confirms that combining an EmptyExpression with another
// expression simply returns the other expression.
func TestEmptyAndOr(t *testing.T) {
	t.Parallel()

	other := Equal("f", 1)

	require.Equal(t, other, Empty().And(other))
	require.Equal(t, other, Empty().Or(other))
}

// TestEmptyShortcuts confirms that each And* shortcut promotes an
// EmptyExpression into a single-predicate expression with the right operator.
func TestEmptyShortcuts(t *testing.T) {
	t.Parallel()

	run := func(name string, got Expression, operator string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			predicate := requirePredicate(t, got)
			require.Equal(t, "f", predicate.Field)
			require.Equal(t, operator, predicate.Operator)
			require.Equal(t, 1, predicate.Value)
		})
	}

	empty := EmptyExpression{}

	run("AndEqual", empty.AndEqual("f", 1), OperatorEqual)
	run("AndNotEqual", empty.AndNotEqual("f", 1), OperatorNotEqual)
	run("AndLessThan", empty.AndLessThan("f", 1), OperatorLessThan)
	run("AndLessOrEqual", empty.AndLessOrEqual("f", 1), OperatorLessOrEqual)
	run("AndGreaterThan", empty.AndGreaterThan("f", 1), OperatorGreaterThan)
	run("AndGreaterOrEqual", empty.AndGreaterOrEqual("f", 1), OperatorGreaterOrEqual)
	run("AndIn", empty.AndIn("f", 1), OperatorIn)
	run("AndNotIn", empty.AndNotIn("f", 1), OperatorNotIn)

	t.Run("AndInAll", func(t *testing.T) {
		t.Parallel()
		// An EmptyExpression collapses to the single new predicate; its variadic
		// arguments are stored as a slice value.
		predicate := requirePredicate(t, empty.AndInAll("f", 1, 2))
		require.Equal(t, "f", predicate.Field)
		require.Equal(t, OperatorInAll, predicate.Operator)
		require.Equal(t, []any{1, 2}, predicate.Value)
	})
}
