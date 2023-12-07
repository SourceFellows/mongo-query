package filter

import (
	"fmt"
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

func (qo QueryOperator) bson() bson.D {
	return bson.D{{qo.operator, qo.value}}
}

// Equals builds a simple QueryOperator for MongoDB operator "$eq".
func Equals(value any) QueryOperator {
	return QueryOperator{operator: "$eq", value: value}
}

// Gt builds a simple QueryOperator for MongoDB operator "$gt".
func Gt(value any) QueryOperator {
	return QueryOperator{operator: "$gt", value: value}
}

// Lt builds a simple QueryOperator for MongoDB operator "$lt".
// "Less than"
func Lt(value any) QueryOperator {
	return QueryOperator{operator: "$lt", value: value}
}

// Lte builds a simple QueryOperator for MongoDB operator "$lte".
// "Less than or equals"
func Lte(value any) QueryOperator {
	return QueryOperator{operator: "$lte", value: value}
}

// Ne builds a simple QueryOperator for MongoDB operator "$ne".
// "not equals"
func Ne(value any) QueryOperator {
	return QueryOperator{operator: "$ne", value: value}
}

// Gte builds a simple QueryOperator for MongoDB operator "$gte".
// "greater or equals"
func Gte(value any) QueryOperator {
	return QueryOperator{operator: "$gte", value: value}
}

// In builds a simple QueryOperator for MongoDB operator "$in".
// "in"
func In(value ...any) QueryOperator {
	return QueryOperator{operator: "$in", value: value}
}

// NotIn builds a simple QueryOperator for MongoDB operator "$nin".
// "in"
func NotIn(value ...any) QueryOperator {
	return QueryOperator{operator: "$nin", value: value}
}

// Exists builds a simple QueryOperator for MongoDB operator "$exists".
// "exists"
func Exists() QueryOperator {
	return QueryOperator{operator: "$exists", value: true}
}

// NotExists builds a simple QueryOperator for MongoDB operator "$exists".
// "not exists"
func NotExists() QueryOperator {
	return QueryOperator{operator: "$exists", value: false}
}

// All builds a simple QueryOperator for MongoDB operator "$all".
// "all"
func All(val ...any) QueryOperator {
	return QueryOperator{operator: "$all", value: val}
}

// Size builds a simple QueryOperator for MongoDB operator "$size".
// "size of array"
func Size(size int) QueryOperator {
	return QueryOperator{operator: "$size", value: size}
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
	return Expression{field: f, value: Gt(value)}
}

// GreaterThan represents a query operation for 'greater than' comparison.
func (f Field) GreaterThan(value any) Expression {
	return f.Gt(value)
}

// Lt represents a query operation for 'less than' comparison.
func (f Field) Lt(value any) Expression {
	return Expression{field: f, value: Lt(value)}
}

// LessThan represents a query operation for 'less than' comparison.
func (f Field) LessThan(value any) Expression {
	return f.Lt(value)
}

// Lte represents a query operation for 'less than or equal' comparison.
func (f Field) Lte(value any) Expression {
	return Expression{field: f, value: Lte(value)}
}

// LessThanOrEqual represents a query operation for 'less than or equals' comparison.
func (f Field) LessThanOrEqual(value any) Expression {
	return f.Lte(value)
}

// Ne represents a query operation for 'not equals' comparison.
func (f Field) Ne(value any) Expression {
	return Expression{field: f, value: Ne(value)}
}

// NotEquals represents a query operation for 'not equals' comparison.
func (f Field) NotEquals(value any) Expression {
	return f.Ne(value)
}

// Gte represents a query operation for 'greater than or equals' comparison.
func (f Field) Gte(value any) Expression {
	return Expression{field: f, value: Gte(value)}
}

// GreaterThanOrEquals represents a query operation for 'greater than or equals' comparison.
func (f Field) GreaterThanOrEquals(value any) Expression {
	return f.Gte(value)
}

// In represents a query operation for 'in' comparison. The operator selects
// the documents where the value of a field equals any value in the specified parameter(s).
func (f Field) In(value ...any) Expression {
	return Expression{field: f, value: In(value...)}
}

// NotIn represents a query operation for 'not in' comparison. The operator selects
// the documents where:
//   - the specified field value is not in the specified array or
//   - the specified field does not exist.
func (f Field) NotIn(value ...any) Expression {
	return Expression{field: f, value: NotIn(value...)}
}

// Exists represents an element query operation to check if a field exists. It Matches
// documents that have the specified field.
func (f Field) Exists() Expression {
	return Expression{field: f, value: Exists()}
}

// NotExists represents an element query operation to check if a field does not exist.
// It Matches documents that do not have the specified field.
func (f Field) NotExists() Expression {
	return Expression{field: f, value: NotExists()}
}

// ArrayContainsAll matches all documents where the given values are in the array.
func (f ArrayField) ArrayContainsAll(val ...any) Expression {
	return Expression{field: Field(f), value: All(val...)}
}

// ArrayContainsExact matches all documents where ONLY the given values are in the array.
func (f ArrayField) ArrayContainsExact(val ...any) Expression {
	return Expression{field: Field(f), value: val}
}

func (f ArrayField) ArrayContainsElement(queries ...QueryOperator) Expression {
	return Expression{field: Field(f), value: queries}
}

// ArrayContainsElementMatchesExpression matches all documents which meet the EXACT expression.
func (f ArrayField) ArrayContainsElementMatchesExpression(expressions ...Expression) Expression {
	return Expression{field: Field(f), value: expressions}
}

func (f ArrayField) ArraySize(size int) Expression {
	return Expression{field: Field(f), value: Size(size)}
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

func (e Expression) String() string {

	data := e.bsonD()
	bytes, _ := bson.Marshal(data)
	mapdata := make(map[string]any)
	_ = bson.Unmarshal(bytes, &mapdata)

	return mapToString(mapdata)
}

func mapToString(mapdata map[string]any) string {
	returnVal := "bson.D{"
	for k, v := range mapdata {
		switch v.(type) {
		case map[string]any:
			returnVal += fmt.Sprintf("{\"%s\", %v}", k, mapToString(v.(map[string]any)))
		case primitive.A:
			value := v.(primitive.A)
			returnVal += fmt.Sprintf("{\"%s\", ", k)
			separator := ""
			returnVal += "[]bson.D{"
			for _, v = range value {
				returnVal += separator
				separator = ", "
				if v1, ok := v.(map[string]any); ok {
					returnVal += fmt.Sprintf("%s", mapToString(v1))
				} else {
					returnVal += fmt.Sprintf("{\"%s\", %v}", k, v)
				}
			}
			returnVal += "}}"
		case int, int32, int64, float32, float64:
			returnVal += fmt.Sprintf("{\"%s\", %v}", k, v)
		default:
			returnVal += fmt.Sprintf("{\"%s\", \"%v\"}", k, v)
		}
	}
	returnVal += "}"

	return returnVal
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
	case []QueryOperator:
		qos := e.value.([]QueryOperator)
		returnValue = bson.D{{string(e.field), queryOperatorsToBSON(qos)}}
	case QueryOperator:
		qo := e.value.(QueryOperator)
		returnValue = bson.D{{string(e.field), qo.bson()}}
	default:
		returnValue = bson.D{{string(e.field), e.value}}
	}

	return returnValue
}

func queryOperatorsToBSON(operators []QueryOperator) bson.D {
	value := bson.D{}
	for _, operator := range operators {
		value = append(value, bson.E{operator.operator, operator.value})
	}
	return value
}

func expressionsToBSON(expressions []Expression) []bson.D {
	values := []bson.D{}
	for _, expression := range expressions {
		var d bson.D
		switch expression.value.(type) {
		case QueryOperator:
			qo := expression.value.(QueryOperator)
			d = bson.D{primitive.E{Key: string(expression.field), Value: bson.D{{Key: qo.operator, Value: qo.value}}}}
		case []QueryOperator:
			qo := expression.value.([]QueryOperator)
			d = bson.D{primitive.E{Key: string(expression.field), Value: queryOperatorsToBSON(qo)}}
		default:
			d = bson.D{primitive.E{Key: string(expression.field), Value: expression.value}}
		}
		values = append(values, d)
	}
	return values
}
