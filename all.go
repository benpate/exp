package exp

// All returns an Expression that matches every record in a dataset.
func All() AndExpression {

	// An empty AndExpression is vacuously true, so it matches everything.  This
	// alias exists so that "give me all of them" reads that way at the call site.
	return And()
}
