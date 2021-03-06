package awsdyn

import (
	"github.com/aws/aws-lambda-go/events"
)

// FromDynamoDBMap generates regular map out of events.DynamoDBAttributeValue map.
func FromDynamoDBMap(record map[string]events.DynamoDBAttributeValue) map[string]interface{} {
	resultMap := make(map[string]interface{})

	for key, rec := range record {
		resultMap[key] = getDynamoDBAttributeValue(rec)
	}

	return resultMap
}

// getDynamoDBAttributeValue gets value from DynamoDBAttributeValue.
func getDynamoDBAttributeValue(record events.DynamoDBAttributeValue) interface{} {
	var val interface{}

	switch record.DataType() {
	case events.DataTypeBinary:
		val = record.Binary()
	case events.DataTypeBinarySet:
		val = record.BinarySet()
	case events.DataTypeBoolean:
		val = record.Boolean()
	case events.DataTypeList:
		list := record.List()
		s := make([]interface{}, 0, len(record.List()))
		for _, el := range list {
			s = append(s, getDynamoDBAttributeValue(el))
		}
		val = s
	case events.DataTypeMap:
		mapData := record.Map()
		// IMPORTANT: For DynamoDB only string can be a key in map.
		m := make(map[string]interface{}, len(mapData))
		for k, el := range mapData {
			m[k] = getDynamoDBAttributeValue(el)
		}
		val = m
	case events.DataTypeNull:
		val = nil
	case events.DataTypeNumber:
		val = record.Number()
	case events.DataTypeNumberSet:
		val = record.NumberSet()
	case events.DataTypeString:
		val = record.String()
	case events.DataTypeStringSet:
		val = record.StringSet()
	}

	return val
}
