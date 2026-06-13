package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseExpression(t *testing.T) {

	{
		expression := Parse("foo eq bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorEqual, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo == bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorEqual, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo ne bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorNotEqual, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo != bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorNotEqual, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo gt bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorGreaterThan, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo > bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorGreaterThan, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo &gt; bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorGreaterThan, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo ge bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorGreaterOrEqual, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo lt bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorLessThan, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		expression := Parse("foo le bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, OperatorLessOrEqual, predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}

	{
		// An unrecognized operator yields an EmptyExpression.
		expression := Parse("foo derp bar")
		require.IsType(t, EmptyExpression{}, expression)
	}
}

// TestParseMalformed confirms that malformed input returns an EmptyExpression
// instead of panicking.
func TestParseMalformed(t *testing.T) {

	for _, input := range []string{"", "foo", "foo eq", "foobar"} {
		require.NotPanics(t, func() { Parse(input) }, "input: %q", input)
		require.IsType(t, EmptyExpression{}, Parse(input), "input: %q", input)
	}
}

// FuzzParse confirms that Parse never panics and always returns a usable
// Expression for arbitrary input.
func FuzzParse(f *testing.F) {

	for _, seed := range []string{
		"",
		" ",
		"  ",
		"foo",
		"foo eq",
		"foo eq bar",
		"foo == bar baz",
		"foo &gt; bar",
		"foo derp bar",
	} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {

		expression := Parse(input)
		require.NotNil(t, expression)

		// Whatever comes back must satisfy the Expression contract without panicking.
		require.NotPanics(t, func() {
			expression.IsEmpty()
			expression.NotEmpty()
			expression.Fields()
			expression.Match(func(Predicate) bool { return true })
		})

		// A successfully parsed predicate must round-trip its field and value out of
		// the original input; anything else must collapse to an EmptyExpression.
		switch value := expression.(type) {
		case Predicate:
			require.Contains(t, input, value.Field)
			require.Contains(t, input, value.Value.(string))
		case EmptyExpression:
			// acceptable result for malformed input
		default:
			t.Fatalf("unexpected expression type %T for input %q", value, input)
		}
	})
}
