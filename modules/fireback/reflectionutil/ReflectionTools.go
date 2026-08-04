package reflectionutil

import (
	"log"
	"reflect"
)

// Converts a list of string into hireachical module 2 structure

func SetValue(obj interface{}, field string, value interface{}) {
	ref := reflect.ValueOf(obj)

	// If it's a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// Double-check we now have a struct
	if ref.Kind() != reflect.Struct {
		log.Fatal("Unexpected type")
	}

	prop := ref.FieldByName(field)

	// Handle setting *string type separately
	if prop.Kind() == reflect.Ptr && prop.Elem().Kind() == reflect.String {
		// Create a new pointer to string and set its value
		newValue := reflect.New(reflect.TypeOf(""))
		newValue.Elem().Set(reflect.ValueOf(value))
		prop.Set(newValue)
	} else {
		if prop.IsValid() {
			prop.Set(reflect.ValueOf(value))
		}
	}
}
