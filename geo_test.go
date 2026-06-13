package exp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// testGeo is a minimal GeoJSONer used to exercise the geometric predicates.
type testGeo struct {
	shape map[string]any
}

func (geo testGeo) GeoJSON() map[string]any {
	return geo.shape
}

// TestGeoPredicates confirms that the geometric constructors store the shape's
// GeoJSON representation as the predicate value.
func TestGeoPredicates(t *testing.T) {
	t.Parallel()

	shape := testGeo{shape: map[string]any{"type": "Point", "coordinates": []any{1, 2}}}

	require.Equal(t,
		Predicate{Field: "location", Operator: OperatorGeoWithin, Value: shape.shape},
		GeoWithin("location", shape),
	)

	require.Equal(t,
		Predicate{Field: "location", Operator: OperatorGeoIntersects, Value: shape.shape},
		GeoIntersects("location", shape),
	)
}
