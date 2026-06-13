package exp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This tests our ability to "collapse" AndExpressions into a single expression, which should keep
// expression trees simpler, and make it easier to traverse/troubleshoot them.
func TestAndExpression(t *testing.T) {

	exp := And(
		And(
			New("field0", "=", 0),
		),
		And(
			New("field1", "=", 1),
			New("field2", "=", 2),
		),
		And(
			New("field3", "=", 3),
			New("field4", "=", 4),
			And(
				New("field5", "=", 5),
				New("field6", "=", 6),
			),
		),
	)

	require.Equal(t, "field0", exp[0].(Predicate).Field)
	require.Equal(t, "field1", exp[1].(Predicate).Field)
	require.Equal(t, "field2", exp[2].(Predicate).Field)
	require.Equal(t, "field3", exp[3].(Predicate).Field)
}

// TestAndNoAliasing confirms that deriving two expressions from a shared base
// does not let one clobber the other through a shared backing array.
func TestAndNoAliasing(t *testing.T) {

	// Three predicates yields len 3, cap 4 (via append doubling), so the base
	// has spare capacity that append would otherwise reuse.
	base := And(Equal("a", 1), Equal("b", 2), Equal("c", 3))

	x := base.And(Equal("X", 9)).(AndExpression)
	y := base.And(Equal("Y", 8)).(AndExpression)

	require.Equal(t, "X", x[3].(Predicate).Field)
	require.Equal(t, "Y", y[3].(Predicate).Field)
}

// TestOrNoAliasing confirms the same independence for OrExpression.
func TestOrNoAliasing(t *testing.T) {

	base := Or(Equal("a", 1), Equal("b", 2), Equal("c", 3))

	x := base.Or(Equal("X", 9)).(OrExpression)
	y := base.Or(Equal("Y", 8)).(OrExpression)

	require.Equal(t, "X", x[3].(Predicate).Field)
	require.Equal(t, "Y", y[3].(Predicate).Field)
}

func TestAndEmpty(t *testing.T) {

	{
		exp := AndExpression{}.And(Empty())
		require.Zero(t, len(exp.(AndExpression)))
	}

	{
		exp := AndExpression{}.Or(Empty())
		require.Zero(t, len(exp.(AndExpression)))
	}

	{
		andExp := And(Predicate{
			Field: "name",
			Value: "John Connor",
		})

		require.Equal(t, 1, len(andExp))

		orExp := andExp.Or(Predicate{
			Field: "name",
			Value: "Sarah Connor",
		})

		require.Equal(t, 2, len(orExp.(OrExpression)))
	}
}

func TestAndEqual(t *testing.T) {

	exp := Equal("field0", 0).AndEqual("field1", 1).AndEqual("field2", "2").AndEqual("field3", 3)

	andExpression := exp.(AndExpression)
	require.Equal(t, 4, len(andExpression))
	require.Equal(t, "field0", andExpression[0].(Predicate).Field)
	require.Equal(t, "field1", andExpression[1].(Predicate).Field)
	require.Equal(t, "field2", andExpression[2].(Predicate).Field)
	require.Equal(t, "field3", andExpression[3].(Predicate).Field)

}

func TestAndExpression4(t *testing.T) {

	exp := Equal("field0", 0).And(Equal("field1", 1)).And(Or(Equal("field2", 2), LessThan("field3", 3)))

	assert.Equal(t, "field0", exp.(AndExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(AndExpression)[0].(Predicate).Operator)
	assert.Equal(t, 0, exp.(AndExpression)[0].(Predicate).Value)

	assert.Equal(t, "field1", exp.(AndExpression)[1].(Predicate).Field)
	assert.Equal(t, "=", exp.(AndExpression)[1].(Predicate).Operator)
	assert.Equal(t, 1, exp.(AndExpression)[1].(Predicate).Value)

	assert.Equal(t, "field2", exp.(AndExpression)[2].(OrExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(AndExpression)[2].(OrExpression)[0].(Predicate).Operator)
	assert.Equal(t, 2, exp.(AndExpression)[2].(OrExpression)[0].(Predicate).Value)

	assert.Equal(t, "field3", exp.(AndExpression)[2].(OrExpression)[1].(Predicate).Field)
	assert.Equal(t, "<", exp.(AndExpression)[2].(OrExpression)[1].(Predicate).Operator)
	assert.Equal(t, 3, exp.(AndExpression)[2].(OrExpression)[1].(Predicate).Value)
}

// TestAndState confirms the empty/not-empty reporting of an AndExpression.
func TestAndState(t *testing.T) {
	t.Parallel()

	require.True(t, And().IsEmpty())
	require.False(t, And().NotEmpty())

	populated := And(Equal("a", 1))
	require.False(t, populated.IsEmpty())
	require.True(t, populated.NotEmpty())
}

// TestAndFields confirms that Fields gathers the field names of every
// sub-expression in order.
func TestAndFields(t *testing.T) {
	t.Parallel()

	exp := And(Equal("a", 1), Equal("b", 2), Or(Equal("c", 3), Equal("d", 4)))
	require.Equal(t, []string{"a", "b", "c", "d"}, exp.Fields())
}

// TestAndMatch confirms AND semantics: every sub-expression must match, and
// evaluation short-circuits on the first failure.
func TestAndMatch(t *testing.T) {
	t.Parallel()

	t.Run("empty matches everything", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.True(t, And().Match(fn))
		require.Zero(t, *calls)
	})

	t.Run("all true matches", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.True(t, And(Equal("a", true), Equal("b", true)).Match(fn))
		require.Equal(t, 2, *calls)
	})

	t.Run("trailing false fails", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.False(t, And(Equal("a", true), Equal("b", false)).Match(fn))
		require.Equal(t, 2, *calls)
	})

	t.Run("leading false short-circuits", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.False(t, And(Equal("a", false), Equal("b", true)).Match(fn))
		require.Equal(t, 1, *calls)
	})
}

// TestAndShortcuts confirms that each And* shortcut appends a predicate with the
// correct operator to the AndExpression.
func TestAndShortcuts(t *testing.T) {
	t.Parallel()

	base := And(Equal("base", 0))

	run := func(name string, got Expression, operator string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			and := requireAnd(t, got, 2)
			require.Equal(t, base[0], and[0]) // original predicate is preserved as the first entry
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
		require.Equal(t, base[0], and[0]) // original predicate is preserved as the first entry
		requireLast(t, and, "f", OperatorInAll, []any{1, 2})
	})
}

// TestAndOr confirms that OR-ing an AndExpression produces an OrExpression, and
// that an empty operand is ignored.
func TestAndOr(t *testing.T) {
	t.Parallel()

	base := And(Equal("a", 1))

	combined := requireOr(t, base.Or(Equal("b", 2)), 2)
	require.Equal(t, base, combined[0])

	require.Equal(t, base, base.Or(Empty()))
}
