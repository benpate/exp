package exp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This tests our ability to "collapse" OrExpressions into a single expression, which should keep
// expression trees simpler, and make it easier to traverse/troubleshoot them.
func TestOrExpression(t *testing.T) {

	exp := Or(
		Or(
			New("field0", "=", 0),
		),
		Or(
			New("field1", "=", 1),
			New("field2", "=", 2),
		),
		Or(
			New("field3", "=", 3),
			New("field4", "=", 4),
			Or(
				New("field5", "=", 5),
				New("field6", "=", 6),
			),
		),
	)

	assert.Equal(t, "field0", exp[0].(Predicate).Field)
	assert.Equal(t, "field1", exp[1].(Predicate).Field)
	assert.Equal(t, "field2", exp[2].(Predicate).Field)
	assert.Equal(t, "field3", exp[3].(Predicate).Field)
}

func TestOrExpression2(t *testing.T) {

	exp := Equal("field0", 0).Or(Equal("field1", 1)).Or(Equal("field2", 2)).Or(LessThan("field3", 3))

	assert.Equal(t, "field0", exp.(OrExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[0].(Predicate).Operator)
	assert.Equal(t, 0, exp.(OrExpression)[0].(Predicate).Value)

	assert.Equal(t, "field1", exp.(OrExpression)[1].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[1].(Predicate).Operator)
	assert.Equal(t, 1, exp.(OrExpression)[1].(Predicate).Value)

	assert.Equal(t, "field2", exp.(OrExpression)[2].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[2].(Predicate).Operator)
	assert.Equal(t, 2, exp.(OrExpression)[2].(Predicate).Value)

	assert.Equal(t, "field3", exp.(OrExpression)[3].(Predicate).Field)
	assert.Equal(t, "<", exp.(OrExpression)[3].(Predicate).Operator)
	assert.Equal(t, 3, exp.(OrExpression)[3].(Predicate).Value)
}

func TestOrExpression3(t *testing.T) {

	exp := Or(Equal("field0", 0), Equal("field1", 1))

	assert.Equal(t, "field0", exp[0].(Predicate).Field)
	assert.Equal(t, "=", exp[0].(Predicate).Operator)
	assert.Equal(t, 0, exp[0].(Predicate).Value)

	assert.Equal(t, "field1", exp[1].(Predicate).Field)
	assert.Equal(t, "=", exp[1].(Predicate).Operator)
	assert.Equal(t, 1, exp[1].(Predicate).Value)

	t.Log(exp)
}

func TestOrExpression4(t *testing.T) {

	exp := Equal("field0", 0).Or(Equal("field1", 1)).Or(And(Equal("field2", 2), LessThan("field3", 3)))

	assert.Equal(t, "field0", exp.(OrExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[0].(Predicate).Operator)
	assert.Equal(t, 0, exp.(OrExpression)[0].(Predicate).Value)

	assert.Equal(t, "field1", exp.(OrExpression)[1].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[1].(Predicate).Operator)
	assert.Equal(t, 1, exp.(OrExpression)[1].(Predicate).Value)

	assert.Equal(t, "field2", exp.(OrExpression)[2].(AndExpression)[0].(Predicate).Field)
	assert.Equal(t, "=", exp.(OrExpression)[2].(AndExpression)[0].(Predicate).Operator)
	assert.Equal(t, 2, exp.(OrExpression)[2].(AndExpression)[0].(Predicate).Value)

	assert.Equal(t, "field3", exp.(OrExpression)[2].(AndExpression)[1].(Predicate).Field)
	assert.Equal(t, "<", exp.(OrExpression)[2].(AndExpression)[1].(Predicate).Operator)
	assert.Equal(t, 3, exp.(OrExpression)[2].(AndExpression)[1].(Predicate).Value)
}
