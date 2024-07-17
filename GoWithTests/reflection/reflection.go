package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	value := GetValue(x)
	countValues := 0

	walkThrought := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkThrought(value.Field(i))
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			walkThrought(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkThrought(value.MapIndex(key))
		}
	}
	for i := 0; i < countValues; i++ {
		Walk(value.Index(i).Interface(), fn)
	}
	/*
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)

			switch field.Kind() {
			case reflect.String:
				fn(field.String())
			case reflect.Struct:
				Walk(field.Interface(), fn)
			case reflect.Slice:
				for i := 0; i < value.Len(); i++ {
					Walk(value.Index(i).Interface(), fn)
				}
			}
		}*/
}

func GetValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
