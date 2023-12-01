package filter

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Field represents a single field in a BSON document.
type Field string

// ArrayField represents an array field in a BSON document.
type ArrayField string

// QueryOperator is used to represent a MongoDB Query Operator.
type QueryOperator struct {
	operator string
	value    any
}

// LogicalOperator is used to represent a logical MongoDB Operator.
type LogicalOperator struct {
	operator    string
	expressions []Expression
}

// Expression represents a MongoDB Expression for a specific field.
type Expression struct {
	field Field
	value any
}

// Equals represents a query operation for 'equals' comparison.
func (f Field) Equals(value any) Expression {
	return Expression{field: f, value: value}
}

// Gt represents a query operation for 'greater than' comparison.
func (f Field) Gt(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$gt", value: value}}
}

// GreaterThan represents a query operation for 'greater than' comparison.
func (f Field) GreaterThan(value any) Expression {
	return f.Gt(value)
}

// Lt represents a query operation for 'less than' comparison.
func (f Field) Lt(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$lt", value: value}}
}

// LessThan represents a query operation for 'less than' comparison.
func (f Field) LessThan(value any) Expression {
	return f.Lt(value)
}

// Lte represents a query operation for 'less than or equal' comparison.
func (f Field) Lte(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$lte", value: value}}
}

// LessThanOrEqual represents a query operation for 'less than or equals' comparison.
func (f Field) LessThanOrEqual(value any) Expression {
	return f.Lte(value)
}

// Ne represents a query operation for 'not equals' comparison.
func (f Field) Ne(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$ne", value: value}}
}

// NotEquals represents a query operation for 'not equals' comparison.
func (f Field) NotEquals(value any) Expression {
	return f.Ne(value)
}

// Gte represents a query operation for 'greater than or equals' comparison.
func (f Field) Gte(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$gte", value: value}}
}

// GreaterThanOrEquals represents a query operation for 'greater than or equals' comparison.
func (f Field) GreaterThanOrEquals(value any) Expression {
	return f.Gte(value)
}

// In represents a query operation for 'in' comparison. The operator selects
// the documents where the value of a field equals any value in the specified parameter(s).
func (f Field) In(value ...any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$in", value: value}}
}

// NotIn represents a query operation for 'not in' comparison. The operator selects
// the documents where:
//   - the specified field value is not in the specified array or
//   - the specified field does not exist.
func (f Field) NotIn(value ...any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$nin", value: value}}
}

// Exists represents a element query operation to check if a field exists. It Matches
// documents that have the specified field.
func (f Field) Exists() Expression {
	return Expression{field: f, value: QueryOperator{operator: "$exists", value: true}}
}

// NotExists represents a element query operation to check if a field does not exist.
// It Matches documents that do not have the specified field.
func (f Field) NotExists() Expression {
	return Expression{field: f, value: QueryOperator{operator: "$exists", value: false}}
}

func (f ArrayField) ArrayContainsAll(val ...any) Expression {
	return Expression{field: Field(f), value: QueryOperator{operator: "$all", value: val}}
}

func (f ArrayField) ArraySize(size int) Expression {
	return Expression{field: Field(f), value: QueryOperator{operator: "$size", value: size}}
}

// And represents a logical query operation for 'and' condition. It takes one or more
// Expression(s) and selects the documents that satisfy all the expressions.
func (e Expression) And(e2 ...Expression) Expression {
	var all []Expression
	all = append(all, e)
	all = append(all, e2...)
	return Expression{value: LogicalOperator{operator: "$and", expressions: all}}
}

// Or represents a logical query operation for 'or' condition. It takes one or more
// Expression(s) and selects the documents that satisfy at least one expression.
func (e Expression) Or(e2 ...Expression) Expression {
	var all []Expression
	all = append(all, e)
	all = append(all, e2...)
	return Expression{value: LogicalOperator{operator: "$or", expressions: all}}
}

// MarshalBSON serializes the Expression to BSON data.
func (e Expression) MarshalBSON() ([]byte, error) {
	data := e.bsonD()
	return bson.Marshal(data)
}

func (e Expression) bsonD() bson.D {

	var returnValue bson.D
	switch e.value.(type) {
	case LogicalOperator:
		qo := e.value.(LogicalOperator)
		returnValue = bson.D{{qo.operator, expressionsToBSON(qo.expressions)}}
	case QueryOperator:
		qo := e.value.(QueryOperator)
		returnValue = bson.D{{string(e.field), bson.D{{qo.operator, qo.value}}}}
	default:
		returnValue = bson.D{{string(e.field), e.value}}
	}

	return returnValue
}

func expressionsToBSON(expressions []Expression) []bson.D {
	values := []bson.D{}
	for _, expression := range expressions {
		var d bson.D
		switch expression.value.(type) {
		case QueryOperator:
			qo := expression.value.(QueryOperator)
			d = bson.D{primitive.E{Key: string(expression.field), Value: bson.D{{Key: qo.operator, Value: qo.value}}}}
		default:
			d = bson.D{primitive.E{Key: string(expression.field), Value: expression.value}}
		}
		values = append(values, d)
	}
	return values
}
