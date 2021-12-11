package exp

// Expression is an interface that is implemented by Predicates, AndExpressions, and OrExpressions.
// It enables any of these items to be embedded into the criteria of a data.Query
type Expression interface {
	And(Expression) Expression
	Or(Expression) Expression
	Match(MatcherFunc) bool
}
