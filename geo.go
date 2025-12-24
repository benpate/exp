package exp

// GeoJSONer is an interface for types that can represent themselves as GeoJSON
type GeoJSONer interface {

	// GeoJSON returns a GeoJSON representation of the object
	GeoJSON() map[string]any
}
