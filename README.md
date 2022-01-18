# Expression Builder

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/benpate/exp)
[![Build Status](http://img.shields.io/travis/benpate/exp.svg?style=flat-square)](https://travis-ci.org/benpate/exp)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/exp.svg?style=flat-square)](https://codecov.io/gh/benpate/exp)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/exp?style=flat-square)](https://goreportcard.com/report/github.com/benpate/exp)
![Version](https://img.shields.io/github/v/release/benpate/exp?include_prereleases&style=flat-square&color=brightgreen)

## Intermediate Expression Formats for the Masses

Every database has its own query language, so this library provides in intermediate format that should be easy to convert into whatever specific language you need to use.  

The expression library only represents the structure of the logical expression, and does not include implementations for any data sources.  Those should be implemented in each individual data source adapter library.

```go

// build single predicate expressions
criteria := exp.Equal("_id", 42)

// use chaining for logical constructs
criteria := exp.Equal("_id", 42).AndEqual("deleteDate",  0)
criteria := exp.Equal("name", "John").OrEqual("name", "Sarah")

// Also supports complex and/or logic

criteria := exp.Or(
    exp.Equal("_id", 42).AndEqual("deleteDate",  0),
    exp.Equal("_id", 42).AndEqual("name", "Sarah"),
)

// Constants define standard expected operators
data.OperatorEqual          = "="
data.OperatorNotEqual       = "!="
data.OperatorLessThan       = "<"
data.OperatorLessOrEqual    = "<="
data.OperatorGreaterThan    = ">"
data.OperatorGreaterOrEqual = ">="
```

## Interfaces

This is accomplished with three very similar data types that all implement the same `Expression` interface.

**`Predicate`** represents a single predicate/comparison.  Using `.And()` and `.Or()` will return the corresponding `AndExpression` or `OrExpression` object

**`AndExpression`** represents multiple predicates, all chained together with AND logic.  Only supports the `.And()` method for additional predicates

**`OrExpression`** represents multiple predicates, all chained together with OR logic.  Only supports the `.Or()` method for additional predicates.

## Manually Walking the Logic Tree

Each of the three interfaces above implements a `.Match()` function that can be used by external programs to see if a dataset matches this exp.  You must pass in a `MatcherFunc` that accepts a predicate and returns `TRUE` if that predicate matches the dataaset.  `AndExpression` and `OrExpression` objects will call this function repeatedly for each element in their logic tree, and return a final boolean value for the entire logic structure.

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 📚
