package filter

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type fullUpdateOperator map[string]bson.D

func (fuo fullUpdateOperator) bson() bson.D {
	values := bson.D{}
	for operator, ops := range fuo {
		values = append(values, bson.E{operator, ops})
	}
	return values
}

type UpdateOperator struct {
	operator string
	values   map[string]any
	field    string
	value    any
}

func (uo UpdateOperator) bson() bson.D {
	values := bson.D{}
	for k, v := range uo.values {
		values = append(values, bson.E{k, v})
	}
	return bson.D{{Key: uo.operator, Value: values}}
}

func (uo UpdateOperator) bsonWithoutOperator() bson.E {
	return bson.E{uo.field, uo.value}
}

func Set(field string, value any) UpdateOperator {
	return UpdateOperator{operator: "$set", field: field, value: value}
}

func CurrentDate(field string) UpdateOperator {
	return UpdateOperator{operator: "$currentDate", field: field, value: true}
}

type UpdateExpression struct {
	value any
}

func (f Field) Set(value any) UpdateExpression {
	return UpdateExpression{value: Set(string(f), value)}
}

func (f Field) CurrentDate() UpdateExpression {
	return UpdateExpression{value: CurrentDate(string(f))}
}

func (ue UpdateExpression) MarshalBSON() ([]byte, error) {
	data := ue.bsonD()
	return bson.Marshal(data)
}

func (ue UpdateExpression) Set(ue2 ...UpdateExpression) UpdateExpression {
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
