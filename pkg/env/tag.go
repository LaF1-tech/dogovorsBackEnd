package env

import (
	"errors"
	"reflect"
	"strconv"
)

func assignStructFromMap(tag string, object interface{}, data map[string]string) error {
	rootObjectValue := reflect.ValueOf(object).Elem()
	typedObjectValue := rootObjectValue
	if typedObjectValue.Kind() == reflect.Interface {
		typedObjectValue = typedObjectValue.Elem()
	}

	var value reflect.Value

	switch typedObjectValue.Kind() {
	case reflect.Slice, reflect.Array:
		for j := 0; true; j++ {
			element := reflect.New(typedObjectValue.Type().Elem()).Elem().Interface()

			arrayElementTag := tag
			if arrayElementTag != "" {
				arrayElementTag += "."
			}
			arrayElementTag += strconv.Itoa(j)

			if err := assignStructFromMap(arrayElementTag, &element, data); err != nil {
				break
			}

			newElement := reflect.Append(typedObjectValue, reflect.ValueOf(element))
			rootObjectValue.Set(newElement)

			typedObjectValue = rootObjectValue
			if typedObjectValue.Kind() == reflect.Interface {
				typedObjectValue = typedObjectValue.Elem()
			}
		}

		return nil
	case reflect.Struct:
		elem := reflect.New(typedObjectValue.Type()).Elem()

		for i := 0; i < typedObjectValue.NumField(); i++ {
			structElement := reflect.New(typedObjectValue.Field(i).Type()).Elem().Interface()

			structElementTag := tag
			if structElementTag != "" {
				structElementTag += "_"
			}
			structElementTag += typedObjectValue.Type().Field(i).Tag.Get("env")

			if err := assignStructFromMap(structElementTag, &structElement, data); err != nil {
				return err
			}

			elem.Field(i).Set(reflect.ValueOf(structElement))
		}
		value = reflect.ValueOf(elem.Interface())
	case reflect.String:
		val, ok := data[tag]
		if !ok {
			return errors.New("no env: " + tag)
		}

		value = reflect.ValueOf(val)
	case reflect.Int, reflect.Int32, reflect.Int64:
		num, err := strconv.Atoi(data[tag])
		if err != nil {
			return err
		}
		value = reflect.ValueOf(num)
	case reflect.Bool:
		value = reflect.ValueOf(data[tag] == "true" || data[tag] == "1")
	default:

	}

	rootObjectValue.Set(value)

	return nil
}
