package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperatorOk(t *testing.T) {

	// Each canonical operator must be recognized.
	for _, operator := range []string{
		OperatorGreaterThan, OperatorGreaterOrEqual,
		OperatorEqual, OperatorNotEqual,
		OperatorLessOrEqual, OperatorLessThan,
		OperatorIn, OperatorNotIn, OperatorInAll,
		OperatorBeginsWith, OperatorEndsWith,
		OperatorContains, OperatorContainedBy,
		OperatorExists, OperatorGeoWithin, OperatorGeoIntersects,
	} {
		result, ok := OperatorOk(operator)
		require.True(t, ok, "operator: %q", operator)
		require.Equal(t, operator, result, "operator: %q", operator)
	}

	// Aliases (including lower-case) normalize to their canonical form.
	for alias, expected := range map[string]string{
		"gt":     OperatorGreaterThan,
		"gte":    OperatorGreaterOrEqual,
		"ge":     OperatorGreaterOrEqual,
		"eq":     OperatorEqual,
		"neq":    OperatorNotEqual,
		"ne":     OperatorNotEqual,
		"lte":    OperatorLessOrEqual,
		"le":     OperatorLessOrEqual,
		"lt":     OperatorLessThan,
		"in all": OperatorInAll,
	} {
		result, ok := OperatorOk(alias)
		require.True(t, ok, "alias: %q", alias)
		require.Equal(t, expected, result, "alias: %q", alias)
	}

	// Unknown input falls back to Equal and reports not-ok.
	result, ok := OperatorOk("derp")
	require.False(t, ok)
	require.Equal(t, OperatorEqual, result)
}

// TestOperator confirms that the Operator wrapper discards the ok flag and
// returns the normalized operator (or the Equal default for unknown input).
func TestOperator(t *testing.T) {
	t.Parallel()

	require.Equal(t, OperatorGreaterThan, Operator("gt"))
	require.Equal(t, OperatorInAll, Operator("in all"))
	require.Equal(t, OperatorEqual, Operator("derp"))
}
