// contains struct comparison tools used for testing
package strava

import (
	"fmt"
	"reflect"
	"testing"
)

func structReflect(object interface{}) (reflect.Type, reflect.Value) {
	v := reflect.ValueOf(object)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		// this long thing allows for pointers to structs and structs to work
		t = reflect.Indirect(reflect.ValueOf(object).Elem()).Type()
	}

	return t, v
}

func structAttributeCount(object interface{}) int {
	// this is dumb but just for tests so...
	return len(structFieldInterfaces(object))
}

// returns a flattened array of interfaces representing the struct values
func structFieldInterfaces(object interface{}) []interface{} {
	t, v := structReflect(object)

	// de pointerize the object
	if v.Kind() == reflect.Ptr {
		v := reflect.Indirect(reflect.ValueOf(object).Elem())
		return structTypeValueInterfaces(v.Type(), v)
	}

	return structTypeValueInterfaces(t, v)
}

func structTypeValueInterfaces(t reflect.Type, v reflect.Value) []interface{} {
	if v.Kind() == reflect.Ptr {
		v := reflect.Indirect(v.Elem())
		return structTypeValueInterfaces(v.Type(), v)
	}

	count := t.NumField()
	values := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		// unexported field
		if !v.Field(i).CanSet() {
			continue
		}

		// recuse down into structs
		if v.Field(i).Kind() == reflect.Struct {
			values = append(values, structTypeValueInterfaces(v.Field(i).Type(), v.Field(i))...)
			continue
		}

		if v.Field(i).Kind() == reflect.Slice {
			for j := 0; j < v.Field(i).Len(); j++ {
				val := v.Field(i).Index(j)
				values = append(values, structTypeValueInterfaces(val.Type(), val)...)
			}

			if v.Field(i).Len() == 0 {
				values = append(values, nil)
			}
			continue
		}

		values = append(values, v.Field(i).Interface())
	}

	return values
}

// structCompare recurvively compares two structs
func structCompare(t *testing.T, o1 interface{}, o2 interface{}) []string {
	o1Type, _ := structReflect(o1)
	o2Type, _ := structReflect(o2)

	probs := make([]string, 0)

	if o1Type != o2Type {
		probs = append(probs, fmt.Sprintf("Type missmatch: %v != %v", o1Type, o2Type))
		return probs
	}

	o1Interfaces := structFieldInterfaces(o1)
	o2Interfaces := structFieldInterfaces(o2)

	if len(o1Interfaces) == 0 || len(o2Interfaces) == 0 {
		probs = append(probs, fmt.Sprintf("No attributes, are you sending a pointer: %v, %v", len(o1Interfaces), len(o2Interfaces)))
		return probs
	}

	if len(o1Interfaces) != len(o2Interfaces) {
		probs = append(probs, fmt.Sprintf("Attributes missmatch, probably slice sizes: %v != %v", len(o1Interfaces), len(o2Interfaces)))
		return probs
	}

	for i := range o1Interfaces {
		if !reflect.DeepEqual(o1Interfaces[i], o2Interfaces[i]) {
			probs = append(probs, fmt.Sprintf("Value missmatch #%d: %v != %v", i, o1Interfaces[i], o2Interfaces[i]))
		}
	}

	return probs
}
