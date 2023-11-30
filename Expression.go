package filter

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Field string
type ArrayField string

type QueryOperator struct {
	operator string
	value    any
}

type LogicalOperator struct {
	operator    string
	expressions []Expression
}

type Expression struct {
	field Field
	value any
}

func (f Field) Equals(value any) Expression {
	return Expression{field: f, value: value}
}

func (f Field) Gt(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$gt", value: value}}
}

func (f Field) Lt(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$lt", value: value}}
}

func (f Field) Lte(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$lte", value: value}}
}

func (f Field) Ne(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$ne", value: value}}
}

func (f Field) Gte(value any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$gte", value: value}}
}

func (f Field) In(value ...any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$in", value: value}}
}

func (f Field) NotIn(value ...any) Expression {
	return Expression{field: f, value: QueryOperator{operator: "$nin", value: value}}
}

func (f Field) Exists() Expression {
	return Expression{field: f, value: QueryOperator{operator: "$exists", value: true}}
}

func (f Field) NotExists() Expression {
	return Expression{field: f, value: QueryOperator{operator: "$exists", value: false}}
}

func (f ArrayField) ArrayContainsAll(val ...any) Expression {
	return Expression{field: Field(f), value: QueryOperator{operator: "$all", value: val}}
}

func (f ArrayField) ArraySize(size int) Expression {
	return Expression{field: Field(f), value: QueryOperator{operator: "$size", value: size}}
}

func (e Expression) And(e2 ...Expression) Expression {
	all := []Expression{}
	all = append(all, e)
	all = append(all, e2...)
	return Expression{value: LogicalOperator{operator: "$and", expressions: all}}
}

func (e Expression) Or(e2 ...Expression) Expression {
	all := []Expression{}
	all = append(all, e)
	all = append(all, e2...)
	return Expression{value: LogicalOperator{operator: "$or", expressions: all}}
}

func (e Expression) MarshalBSON() ([]byte, error) {
	data := e.bsonD()
	fmt.Println(data)
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
