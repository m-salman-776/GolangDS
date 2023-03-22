package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func JSONEncode(v interface{}) ([]byte, error) {
	refObjectValue := reflect.ValueOf(v)
	refObjectType := reflect.TypeOf(v)
	buff := bytes.Buffer{}
	buff.WriteString("{")
	if refObjectValue.Kind() != reflect.Struct {
		return buff.Bytes(), fmt.Errorf(" Value of kind %s is not supported", refObjectValue.Kind())
	}
	keyValPair := make([]string, 0, refObjectValue.NumField()+10)
	for i := 0; i < refObjectValue.NumField(); i++ {
		structFieldRefObjVal := refObjectValue.Field(i)
		structFieldRefObjType := refObjectType.Field(i)
		tag := structFieldRefObjType.Tag.Get("json")
		switch structFieldRefObjVal.Kind() {
		case reflect.String:
			value := structFieldRefObjVal.Interface().(string)
			keyValPair = append(keyValPair, `"`+string(tag)+`":"`+value+`"`)
			//tag := structFieldRefObjType.Tag
		case reflect.Int64:
			value := structFieldRefObjVal.Interface().(int64)
			//tag := string(structFieldRefObjType.Tag)
			keyValPair = append(keyValPair, `"`+tag+`":"`+fmt.Sprintf("%d", value)+`"`)
		default:
			return buff.Bytes(), fmt.Errorf(" Filed with the name %s and the kind %s is not supported", structFieldRefObjType.Name, structFieldRefObjVal.Kind())
		}
	}
	buff.WriteString(strings.Join(keyValPair, ","))
	buff.Write([]byte{'}'})
	return buff.Bytes(), nil
}
