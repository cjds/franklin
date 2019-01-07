// Author(s): Carl Saldanha

package model

import "reflect"

// copyStructFields performs an O(n) operation to copy source struct field values to target struct
// field values, given that the fields have the same name and data type.
func copyStructFields(src interface{}, target interface{}) {
	mapFieldNameToTypeToValue := make(map[string]map[reflect.Type]reflect.Value)

	// Populate map with src data
	srcStruct := reflect.ValueOf(src).Elem()
	for i := 0; i < srcStruct.NumField(); i++ {
		srcFieldValue := srcStruct.Field(i)
		srcFieldType := srcStruct.Type().Field(i)
		mapFieldNameToTypeToValue[srcFieldType.Name] = map[reflect.Type]reflect.Value{
			srcFieldType.Type: srcFieldValue,
		}
	}

	// Check map to copy src fields to target
	targetStruct := reflect.ValueOf(target).Elem()
	for i := 0; i < targetStruct.NumField(); i++ {
		targetFieldValue := targetStruct.Field(i)
		targetFieldType := targetStruct.Type().Field(i)
		mapFieldTypeToValue, ok1 := mapFieldNameToTypeToValue[targetFieldType.Name]
		if ok1 {
			srcValue, ok2 := mapFieldTypeToValue[targetFieldType.Type]
			if ok2 {
				targetFieldValue.Set(reflect.Value(srcValue))
			}
		}
	}
}

// copyStructFieldsSkipEmptySlice performs an O(n) operation to copy source struct field values to
// to target struct field values, except when the source field is a slice and it is empty.
func copyStructFieldsSkipEmptySlice(src interface{}, target interface{}) {
	mapFieldNameToTypeToValue := make(map[string]map[reflect.Type]reflect.Value)

	// Populate map with src data
	srcStruct := reflect.ValueOf(src).Elem()
	for i := 0; i < srcStruct.NumField(); i++ {
		if srcStruct.Field(i).Kind() == reflect.Slice && srcStruct.Field(i).Len() == 0 {
			continue
		}

		srcFieldValue := srcStruct.Field(i)
		srcFieldType := srcStruct.Type().Field(i)
		mapFieldNameToTypeToValue[srcFieldType.Name] = map[reflect.Type]reflect.Value{
			srcFieldType.Type: srcFieldValue,
		}
	}

	// Check map to copy src fields to target
	targetStruct := reflect.ValueOf(target).Elem()
	for i := 0; i < targetStruct.NumField(); i++ {
		targetFieldValue := targetStruct.Field(i)
		targetFieldType := targetStruct.Type().Field(i)
		mapFieldTypeToValue, ok1 := mapFieldNameToTypeToValue[targetFieldType.Name]
		if ok1 {
			srcValue, ok2 := mapFieldTypeToValue[targetFieldType.Type]
			if ok2 {
				targetFieldValue.Set(reflect.Value(srcValue))
			}
		}
	}
}
