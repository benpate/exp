package exp

// MatcherFunc reports whether a single Predicate matches the caller's data.
// It is passed to the Match method of every Expression.
type MatcherFunc func(Predicate) bool
