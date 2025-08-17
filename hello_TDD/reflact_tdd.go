package main

import "reflect"


// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/reflection#zhong-gou-3
func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walkValue(val.MapIndex(k))
		}
	}

	//numberOfValues := 0
	//var getField func(int) reflect.Value
	//
	//switch val.Kind() {
	//case reflect.Struct:
	//	numberOfValues = val.NumField()
	//	getField = val.Field
	//	//for i := 0; i < val.NumField(); i++ {
	//	//	walk(val.Field(i).Interface(), fn)
	//	//}
	//case reflect.Slice, reflect.Array:
	//	numberOfValues = val.Len()
	//	getField = val.Index
	//	//for i := 0; i < val.Len(); i++ {
	//	//	walk(val.Index(i).Interface(), fn)
	//	//}
	//case reflect.String:
	//	fn(val.String())
	//case reflect.Map:
	//	for _, k := range val.MapKeys() {
	//		walk(val.MapIndex(k).Interface(), fn)
	//	}
	//}
	//
	//for i := 0; i < numberOfValues; i++ {
	//	walk(getField(i).Interface(), fn)
	//}

	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//
	//	switch field.Kind() {
	//	case reflect.String:
	//		fn(field.String())
	//	case reflect.Struct:
	//		walk(field.Interface(), fn)
	//	}
	//}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// 指针类型的 Value 不能使用 NumField 方法，在执行此方法前需要调用 Elem() 提取底层值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
