package file

import (
	"fmt"
	"reflect"
	"strings"

	"4.1.H/domain/prescription"
)

func Prescription2Csv(p prescription.Prescription) ([]byte, error) {
	header := getStructFieldNames(p)
	raw := getStructFieldValues(p)
	csv := strings.Join(header, ",") + "\n" + strings.Join(raw, ",")
	return []byte(csv), nil
}

func getStructFieldValues(data interface{}) []string {
	v := reflect.ValueOf(data)
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		if value, ok := v.Field(i).Interface().(string); ok {
			values[i] = value
		} else {
			values[i] = fmt.Sprintf("%v", v.Field(i))
		}
	}
	return values
}

func getStructFieldNames(data interface{}) []string {
	t := reflect.TypeOf(data)
	names := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		names[i] = t.Field(i).Name
	}
	return names
}
