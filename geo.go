package exp

type GeoJSONer interface {
	GeoJSON() map[string]any
}
