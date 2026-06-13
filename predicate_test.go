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

// TestPredicateConstructors covers the remaining constructors that build set
// and existence predicates.
func TestPredicateConstructors(t *testing.T) {
	t.Parallel()

	require.Equal(t, Predicate{Field: "field", Operator: OperatorIn, Value: 1}, In("field", 1))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorNotIn, Value: 1}, NotIn("field", 1))

	// InAll is variadic and collects its arguments into a slice.
	require.Equal(t, Predicate{Field: "field", Operator: OperatorInAll, Value: []any{1, 2}}, InAll("field", 1, 2))

	// Exists / NotExists encode the desired result in the Value.
	require.Equal(t, Predicate{Field: "field", Operator: OperatorExists, Value: true}, Exists("field"))
	require.Equal(t, Predicate{Field: "field", Operator: OperatorExists, Value: false}, NotExists("field"))
}

// TestPredicateState confirms that a Predicate is never empty and reports its
// own field.
func TestPredicateState(t *testing.T) {
	t.Parallel()

	predicate := Equal("name", "value")

	require.False(t, predicate.IsEmpty())
	require.True(t, predicate.NotEmpty())
	require.Equal(t, []string{"name"}, predicate.Fields())
}

// TestPredicateMatch confirms that Match simply delegates to the MatcherFunc.
func TestPredicateMatch(t *testing.T) {
	t.Parallel()

	fn, calls := boolMatcher()

	require.True(t, Equal("f", true).Match(fn))
	require.False(t, Equal("f", false).Match(fn))
	require.Equal(t, 2, *calls)
}

// TestPredicateAndShortcuts confirms that each And* shortcut appends a second
// predicate with the correct operator, yielding a two-element AndExpression.
func TestPredicateAndShortcuts(t *testing.T) {
	t.Parallel()

	base := Equal("base", 0)

	run := func(name string, got Expression, operator string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			and := requireAnd(t, got, 2)
			require.Equal(t, base, and[0]) // base predicate is preserved as the first entry
			requireLast(t, and, "f", operator, 1)
		})
	}

	run("AndEqual", base.AndEqual("f", 1), OperatorEqual)
	run("AndNotEqual", base.AndNotEqual("f", 1), OperatorNotEqual)
	run("AndLessThan", base.AndLessThan("f", 1), OperatorLessThan)
	run("AndLessOrEqual", base.AndLessOrEqual("f", 1), OperatorLessOrEqual)
	run("AndGreaterThan", base.AndGreaterThan("f", 1), OperatorGreaterThan)
	run("AndGreaterOrEqual", base.AndGreaterOrEqual("f", 1), OperatorGreaterOrEqual)
	run("AndIn", base.AndIn("f", 1), OperatorIn)
	run("AndNotIn", base.AndNotIn("f", 1), OperatorNotIn)

	t.Run("AndInAll", func(t *testing.T) {
		t.Parallel()
		and := requireAnd(t, base.AndInAll("f", 1, 2), 2)
		require.Equal(t, base, and[0]) // base predicate is preserved as the first entry
		requireLast(t, and, "f", OperatorInAll, []any{1, 2})
	})
}

// TestPredicateOrShortcuts confirms that each Or* shortcut appends a second
// predicate with the correct operator, yielding a two-element OrExpression.
func TestPredicateOrShortcuts(t *testing.T) {
	t.Parallel()

	base := Equal("base", 0)

	run := func(name string, got Expression, operator string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			or := requireOr(t, got, 2)
			require.Equal(t, base, or[0]) // base predicate is preserved as the first entry
			requireLast(t, or, "f", operator, 1)
		})
	}

	run("OrEqual", base.OrEqual("f", 1), OperatorEqual)
	run("OrNotEqual", base.OrNotEqual("f", 1), OperatorNotEqual)
	run("OrLessThan", base.OrLessThan("f", 1), OperatorLessThan)
	run("OrLessOrEqual", base.OrLessOrEqual("f", 1), OperatorLessOrEqual)
	run("OrGreaterThan", base.OrGreaterThan("f", 1), OperatorGreaterThan)
	run("OrGreaterOrEqual", base.OrGreaterOrEqual("f", 1), OperatorGreaterOrEqual)
	run("OrIn", base.OrIn("f", 1), OperatorIn)
	run("OrNotIn", base.OrNotIn("f", 1), OperatorNotIn)

	t.Run("OrInAll", func(t *testing.T) {
		t.Parallel()
		or := requireOr(t, base.OrInAll("f", 1, 2), 2)
		require.Equal(t, base, or[0]) // base predicate is preserved as the first entry
		requireLast(t, or, "f", OperatorInAll, []any{1, 2})
	})
}

// TestPredicateAndOrEmpty confirms that combining a predicate with an
// EmptyExpression returns the predicate unchanged.
func TestPredicateAndOrEmpty(t *testing.T) {
	t.Parallel()

	predicate := Equal("f", 1)

	require.Equal(t, predicate, predicate.And(Empty()))
	require.Equal(t, predicate, predicate.Or(Empty()))
}
