// Package exp builds database-agnostic query expressions.
//
// Every database has its own query language, so this package provides an
// intermediate format that adapter libraries translate into whatever dialect
// they need.  It describes only the shape of a logical expression, and contains
// no implementation for any particular data source.
//
// Expressions are built from four types, all of which satisfy the Expression
// interface:
//
//	Predicate        a single comparison, such as `name = "Sarah"`
//	AndExpression    a list of sub-expressions, all of which must match
//	OrExpression     a list of sub-expressions, any one of which may match
//	EmptyExpression  no constraint at all, which matches everything
//
// Predicates chain into larger expressions using the And* and Or* methods:
//
//	criteria := exp.Equal("_id", 42).AndEqual("deleteDate", 0)
//
// To walk a finished expression, pass a MatcherFunc to its Match method.  The
// MatcherFunc answers "does this one Predicate match my data?" and the
// Expression handles the boolean logic around it.
package exp
