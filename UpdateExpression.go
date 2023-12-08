package filter

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type fullUpdateOperator map[string]bson.D

func (fuo fullUpdateOperator) bson() bson.D {
	values := bson.D{}
	for operator, ops := range fuo {
		values = append(values, bson.E{Key: operator, Value: ops})
	}
	return values
}

type UpdateOperator struct {
	operator string
	values   map[Field]any
	field    Field
	value    any
}

func (uo UpdateOperator) bson() bson.D {
	values := bson.D{}
	for k, v := range uo.values {
		values = append(values, bson.E{Key: string(k), Value: v})
	}
	return bson.D{{Key: uo.operator, Value: values}}
}

func (uo UpdateOperator) bsonWithoutOperator() bson.E {
	return bson.E{Key: string(uo.field), Value: uo.value}
}

func set(field Field, value any) UpdateOperator {
	return UpdateOperator{operator: "$set", field: field, value: value}
}

func inc(field Field, value any) UpdateOperator {
	return UpdateOperator{operator: "$inc", field: field, value: value}
}

func min(field Field, value any) UpdateOperator {
	return UpdateOperator{operator: "$min", field: field, value: value}
}

func max(field Field, value any) UpdateOperator {
	return UpdateOperator{operator: "$max", field: field, value: value}
}

func mul(field Field, value any) UpdateOperator {
	return UpdateOperator{operator: "$mul", field: field, value: value}
}

func rename(field Field, value string) UpdateOperator {
	return UpdateOperator{operator: "$rename", field: field, value: value}
}

func unset(field Field) UpdateOperator {
	return UpdateOperator{operator: "$unset", field: field, value: ""}
}

func currentDate(field Field) UpdateOperator {
	return UpdateOperator{operator: "$currentDate", field: "lastModified", value: true}
}

type UpdateExpression struct {
	value any
}

// Set replaces the value of the field with the specified value.
func (f Field) Set(value any) UpdateExpression {
	return UpdateExpression{value: set(f, value)}
}

// Inc increments the field by a specified value.
func (f Field) Inc(value any) UpdateExpression {
	return UpdateExpression{value: inc(f, value)}
}

// Min updates the value of the field to a specified value if the specified value
// is less than the current value of the field.
//
// The '$min operator', which is used internally, can compare values of different
// types, using the BSON comparison order.
func (f Field) Min(value any) UpdateExpression {
	return UpdateExpression{value: min(f, value)}
}

// Max updates the value of the field to a specified value if the specified value is
// greater than the current value of the field.
//
// The '$max operator', which is used internally, can compare values of different
// types, using the BSON comparison order.
func (f Field) Max(value any) UpdateExpression {
	return UpdateExpression{value: max(f, value)}
}

// Mul multiplies the value of the field by a number.
//
// The field to update must contain a numeric value.
func (f Field) Mul(value any) UpdateExpression {
	return UpdateExpression{value: mul(f, value)}
}

// Rename updates the name of the field.
//
// The new field name must differ from the existing field name.
func (f Field) Rename(value Field) UpdateExpression {
	return UpdateExpression{value: rename(f, string(value))}
}

// Unset deletes the particular field.
func (f Field) Unset() UpdateExpression {
	return UpdateExpression{value: unset(f)}
}

// CurrentDate sets the value of a field to the current date, either as a Date or
// a timestamp.
func (f Field) CurrentDate() UpdateExpression {
	return UpdateExpression{value: currentDate(f)}
}

func (ue UpdateExpression) MarshalBSON() ([]byte, error) {
	data := ue.bsonD()
	return bson.Marshal(data)
}

func (ue UpdateExpression) And(ue2 ...UpdateExpression) UpdateExpression {
	var all []any
	all = append(all, ue.value)
	for _, upex := range ue2 {
		all = append(all, upex.value)
	}

	return UpdateExpression{value: expressionCollector{value: all}}
}

type expressionCollector struct {
	value []any
}

func (ue UpdateExpression) bsonD() bson.D {
	var returnValue bson.D
	switch ue.value.(type) {
	case expressionCollector:
		ce := ue.value.(expressionCollector)
		fuo := fullUpdateOperator{}
		for _, val := range ce.value {
			uo, ok := val.(UpdateOperator)
			if !ok {
				log.Println("not an update operator")
				continue
			}

			fuo[uo.operator] = append(fuo[uo.operator], uo.bsonWithoutOperator())
		}

		returnValue = fuo.bson()
	case UpdateOperator:
		uo := ue.value.(UpdateOperator)
		returnValue = uo.bson()
	default:
		// TODO: check default value
		// is there a reasonable default value?? maybe it's better to not update as default..
		returnValue = bson.D{{"$noop", ue.value}}
	}

	log.Println(returnValue)
	return returnValue
}
