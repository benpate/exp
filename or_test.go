package exp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This tests our ability to "collapse" OrExpressions into a single expression, which should keep
// expression trees simpler, and make it easier to traverse/troubleshoot them.
func TestOrExpression(t *testing.T) {

	exp := Or(
		Or(
			New("field0", "=", 0),
		),
		Or(
			New("field1", "=", 1),
			New("field2", "=", 2),
		),
		Or(
			New("field3", "=", 3),
			New("field4", "=", 4),
			Or(
				New("field5", "=", 5),
				New("field6", "=", 6),
			),
		),
	)

	assert.Equal(t, "field0", exp[0].(Predicate).Field)
	assert.Equal(t, "field1", exp[1].(Predicate).Field)
	assert.Equal(t, "field2", exp[2].(Predicate).Field)
	assert.Equal(t, "field3", exp[3].(Predicate).Field)
}

func TestOrExpression2(t *testing.T) {

	exp := Equal("field0", 0).Or(Equal("field1", 1)).Or(Equal("field2", 2)).Or(LessThan("field3", 3))

	assert.Equal(t, "field0", exp.(OrExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[0].(Predicate).Operator)
	assert.Equal(t, 0, exp.(OrExpression)[0].(Predicate).Value)

	assert.Equal(t, "field1", exp.(OrExpression)[1].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[1].(Predicate).Operator)
	assert.Equal(t, 1, exp.(OrExpression)[1].(Predicate).Value)

	assert.Equal(t, "field2", exp.(OrExpression)[2].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[2].(Predicate).Operator)
	assert.Equal(t, 2, exp.(OrExpression)[2].(Predicate).Value)

	assert.Equal(t, "field3", exp.(OrExpression)[3].(Predicate).Field)
	assert.Equal(t, "<", exp.(OrExpression)[3].(Predicate).Operator)
	assert.Equal(t, 3, exp.(OrExpression)[3].(Predicate).Value)
}

func TestOrExpression3(t *testing.T) {

	exp := Or(Equal("field0", 0), Equal("field1", 1))

	assert.IsType(t, OrExpression{}, exp)
	assert.Len(t, exp, 2)

	assert.Equal(t, "field0", exp[0].(Predicate).Field)
	assert.Equal(t, "=", exp[0].(Predicate).Operator)
	assert.Equal(t, 0, exp[0].(Predicate).Value)

	assert.Equal(t, "field1", exp[1].(Predicate).Field)
	assert.Equal(t, "=", exp[1].(Predicate).Operator)
	assert.Equal(t, 1, exp[1].(Predicate).Value)
}

func TestOrExpression4(t *testing.T) {

	exp := Equal("field0", 0).Or(Equal("field1", 1)).Or(And(Equal("field2", 2), LessThan("field3", 3)))

	assert.IsType(t, OrExpression{}, exp)
	assert.Len(t, exp, 3)

	assert.Equal(t, "field0", exp.(OrExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[0].(Predicate).Operator)
	assert.Equal(t, 0, exp.(OrExpression)[0].(Predicate).Value)

	assert.Equal(t, "field1", exp.(OrExpression)[1].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[1].(Predicate).Operator)
	assert.Equal(t, 1, exp.(OrExpression)[1].(Predicate).Value)

	assert.Equal(t, "field2", exp.(OrExpression)[2].(AndExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[2].(AndExpression)[0].(Predicate).Operator)
	assert.Equal(t, 2, exp.(OrExpression)[2].(AndExpression)[0].(Predicate).Value)

	assert.Equal(t, "field3", exp.(OrExpression)[2].(AndExpression)[1].(Predicate).Field)
	assert.Equal(t, "<", exp.(OrExpression)[2].(AndExpression)[1].(Predicate).Operator)
	assert.Equal(t, 3, exp.(OrExpression)[2].(AndExpression)[1].(Predicate).Value)
}

// TestOrState confirms the empty/not-empty reporting of an OrExpression.
func TestOrState(t *testing.T) {
	t.Parallel()

	require.True(t, Or().IsEmpty())
	require.False(t, Or().NotEmpty())

	populated := Or(Equal("a", 1))
	require.False(t, populated.IsEmpty())
	require.True(t, populated.NotEmpty())
}

// TestOrFields confirms that Fields gathers the field names of every
// sub-expression in order.
func TestOrFields(t *testing.T) {
	t.Parallel()

	exp := Or(Equal("a", 1), Equal("b", 2), And(Equal("c", 3), Equal("d", 4)))
	require.Equal(t, []string{"a", "b", "c", "d"}, exp.Fields())

	// An empty OrExpression returns an empty (non-nil) slice.
	require.Equal(t, []string{}, Or().Fields())
	require.NotNil(t, Or().Fields())
}

// TestOrMatch confirms OR semantics: any sub-expression matching is enough, and
// evaluation short-circuits on the first success.
func TestOrMatch(t *testing.T) {
	t.Parallel()

	t.Run("empty matches nothing", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.False(t, Or().Match(fn))
		require.Zero(t, *calls)
	})

	t.Run("all false fails", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.False(t, Or(Equal("a", false), Equal("b", false)).Match(fn))
		require.Equal(t, 2, *calls)
	})

	t.Run("trailing true matches", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.True(t, Or(Equal("a", false), Equal("b", true)).Match(fn))
		require.Equal(t, 2, *calls)
	})

	t.Run("leading true short-circuits", func(t *testing.T) {
		t.Parallel()
		fn, calls := boolMatcher()
		require.True(t, Or(Equal("a", true), Equal("b", false)).Match(fn))
		require.Equal(t, 1, *calls)
	})
}

// TestOrShortcuts confirms that each And* shortcut converts an OrExpression into
// an AndExpression combining the original with the new predicate.
func TestOrShortcuts(t *testing.T) {
	t.Parallel()

	base := Or(Equal("base", 0))

	run := func(name string, got Expression, operator string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			and := requireAnd(t, got, 2)
			require.Equal(t, base, and[0]) // original OrExpression is preserved as the first entry
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
		require.Equal(t, base, and[0]) // original OrExpression is preserved as the first entry
		requireLast(t, and, "f", OperatorInAll, []any{1, 2})
	})
}

// TestOrAnd confirms that AND-ing an OrExpression produces an AndExpression, and
// that an empty operand is ignored.
func TestOrAnd(t *testing.T) {
	t.Parallel()

	base := Or(Equal("a", 1))

	combined := requireAnd(t, base.And(Equal("b", 2)), 2)
	require.Equal(t, base, combined[0])

	require.Equal(t, base, base.And(Empty()))
}

// TestOrOrEmpty confirms that OR-ing an OrExpression with an empty operand
// returns the original expression unchanged.
func TestOrOrEmpty(t *testing.T) {
	t.Parallel()

	base := Or(Equal("a", 1))

	require.Equal(t, base, base.Or(Empty()))
}

// TestOrOrShortcuts confirms that each Or* shortcut flattens the new predicate
// into the existing OrExpression.
func TestOrOrShortcuts(t *testing.T) {
	t.Parallel()

	base := Or(Equal("base", 0))

	run := func(name string, got Expression, operator string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			or := requireOr(t, got, 2)
			require.Equal(t, base[0], or[0]) // original predicate is preserved as the first entry
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
		require.Equal(t, base[0], or[0]) // original predicate is preserved as the first entry
		requireLast(t, or, "f", OperatorInAll, []any{1, 2})
	})
}
