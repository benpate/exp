package exp

// Expression is an interface that is implemented by Predicates, AndExpressions, and OrExpressions.
// It enables any of these items to be embedded into the criteria of a data.Query
type Expression interface {

	// And returns a new expression that combines this expression with another as an AndExpression
	And(Expression) Expression

	// Or returns a new expression that combines this expression with another as an OrExpression
	Or(Expression) Expression

	Match(MatcherFunc) bool

	// AndEqual is a shortcut that creates a new AndExpression using the Equal comparison
	AndEqual(name string, value any) Expression

	// AndNotEqual is a shortcut that creates a new AndExpression using the NotEqual comparison
	AndNotEqual(name string, value any) Expression

	// AndLessThan is a shortcut that creates a new AndExpression using the LessThan comparison
	AndLessThan(name string, value any) Expression

	// AndLessOrEqual is a shortcut that creates a new AndExpression using the LessOrEqual comparison
	AndLessOrEqual(name string, value any) Expression

	// AndGreaterThan is a shortcut that creates a new AndExpression using the GreaterThan comparison
	AndGreaterThan(name string, value any) Expression

	// AndGreaterOrEqual is a shortcut that creates a new AndExpression using the GreaterOrEqual comparison
	AndGreaterOrEqual(name string, value any) Expression

	// AndIn is a shortcut that creates a new AndExpression using the In comparison
	AndIn(name string, value any) Expression

	// AndNotIn is a shortcut that creates a new AndExpression using the NotIn comparison
	AndNotIn(name string, value any) Expression

	// AndInAll is a shortcut that creates a new AndExpression using the InAll comparison
	AndInAll(name string, value ...any) Expression

	// IsEmpty returns TRUE if an expression does not have any predicates
	IsEmpty() bool

	// NotEmpty returns TRUE if an expression has one or more predicates
	NotEmpty() bool

	// Fields returns the list of fields that are used in this expression
	Fields() []string
}
