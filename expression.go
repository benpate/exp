package exp

// Expression is an interface that is implemented by Predicates, AndExpressions, and OrExpressions.
// It enables any of these items to be embedded into the criteria of a data.Query
type Expression interface {

	// Match evaluates the expression against a MatcherFunc, returning TRUE if the expression matches
	Match(MatcherFunc) bool

	// And returns a new expression that combines this expression with another as an AndExpression
	And(Expression) Expression

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

	// Or returns a new expression that combines this expression with another as an OrExpression
	Or(Expression) Expression

	// OrEqual is a shortcut that creates a new OrExpression using the Equal comparison
	OrEqual(name string, value any) Expression

	// OrNotEqual is a shortcut that creates a new OrExpression using the NotEqual comparison
	OrNotEqual(name string, value any) Expression

	// OrLessThan is a shortcut that creates a new OrExpression using the LessThan comparison
	OrLessThan(name string, value any) Expression

	// OrLessOrEqual is a shortcut that creates a new OrExpression using the LessOrEqual comparison
	OrLessOrEqual(name string, value any) Expression

	// OrGreaterThan is a shortcut that creates a new OrExpression using the GreaterThan comparison
	OrGreaterThan(name string, value any) Expression

	// OrGreaterOrEqual is a shortcut that creates a new OrExpression using the GreaterOrEqual comparison
	OrGreaterOrEqual(name string, value any) Expression

	// OrIn is a shortcut that creates a new OrExpression using the In comparison
	OrIn(name string, value any) Expression

	// OrNotIn is a shortcut that creates a new OrExpression using the NotIn comparison
	OrNotIn(name string, value any) Expression

	// OrInAll is a shortcut that creates a new OrExpression using the InAll comparison
	OrInAll(name string, value ...any) Expression

	// IsEmpty returns TRUE if an expression does not have any predicates
	IsEmpty() bool

	// NotEmpty returns TRUE if an expression has one or more predicates
	NotEmpty() bool

	// Fields returns the list of fields that are used in this expression
	Fields() []string
}
