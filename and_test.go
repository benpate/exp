package exp

import (
	"testing"

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

func TestAndEqual(t *testing.T) {

	exp := Equal("field0", 0).AndEqual("field1", 1).AndEqual("field2", "2").AndEqual("field3", 3)

	andExpression := exp.(AndExpression)
	require.Equal(t, 4, len(andExpression))
	require.Equal(t, "field0", andExpression[0].(Predicate).Field)
	require.Equal(t, "field1", andExpression[1].(Predicate).Field)
	require.Equal(t, "field2", andExpression[2].(Predicate).Field)
	require.Equal(t, "field3", andExpression[3].(Predicate).Field)

}
