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
		expression := Parse("foo derp bar")
		predicate := expression.(Predicate)
		require.Equal(t, "foo", predicate.Field)
		require.Equal(t, "", predicate.Operator)
		require.Equal(t, "bar", predicate.Value)
	}
}
