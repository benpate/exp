package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNilExpressionNeverEntersTree is the regression test for a nil Expression
// being appended into a logic tree.  A nil element does not fail where it was
// introduced; it panics later, inside Match() or Fields(), so every combining
// method must absorb it at the point of entry.
func TestNilExpressionNeverEntersTree(t *testing.T) {
	t.Parallel()

	// Every way to combine a nil Expression, from every concrete type.
	combined := map[string]Expression{
		"And(nil)":            And(nil),
		"And(nil, nil)":       And(nil, nil),
		"Or(nil)":             Or(nil),
		"Or(nil, nil)":        Or(nil, nil),
		"AndExpression.And":   And(Equal("a", true)).And(nil),
		"AndExpression.Or":    And(Equal("a", true)).Or(nil),
		"OrExpression.Or":     Or(Equal("a", true)).Or(nil),
		"OrExpression.And":    Or(Equal("a", true)).And(nil),
		"Predicate.And":       Equal("a", true).And(nil),
		"Predicate.Or":        Equal("a", true).Or(nil),
		"EmptyExpression.And": Empty().And(nil),
		"EmptyExpression.Or":  Empty().Or(nil),
	}

	for name, expression := range combined {

		require.NotNil(t, expression, "%s returned a nil Expression", name)

		// A nil hiding in the tree only surfaces when the tree is walked, so walk it.
		require.NotPanics(t, func() {
			expression.Match(func(Predicate) bool { return false })
			expression.Match(func(Predicate) bool { return true })
			expression.Fields()
			expression.IsEmpty()
			expression.NotEmpty()
		}, "%s panicked while being evaluated", name)
	}
}

// TestNilExpressionIsIdentity confirms that absorbing a nil leaves the original
// expression untouched, rather than silently widening or narrowing it.
func TestNilExpressionIsIdentity(t *testing.T) {
	t.Parallel()

	predicate := Equal("a", true)
	and := And(predicate)
	or := Or(predicate)

	require.Equal(t, predicate, predicate.And(nil))
	require.Equal(t, predicate, predicate.Or(nil))
	require.Equal(t, and, and.And(nil))
	require.Equal(t, and, and.Or(nil))
	require.Equal(t, or, or.Or(nil))
	require.Equal(t, or, or.And(nil))
	require.Equal(t, Empty(), Empty().And(nil))
	require.Equal(t, Empty(), Empty().Or(nil))

	// A nil contributes no fields and no length.
	require.Equal(t, []string{"a"}, and.And(nil).Fields())
	require.Len(t, And(nil, predicate, nil), 1)
	require.Len(t, Or(nil, predicate, nil), 1)
}
