package main

import (
	"errors"
	"fmt"
	"reflect"
)

/*
	假设需要将map的数据填充到Struct对应的属性，如何能够适应任意map和Struct？
 */
type Teacher struct {
	Name	string
	Class	string
}

func FillStructByMap() {
	m := map[string]interface{}{"Name": "pyl", "Age": 33, "Class": "Math"}
	s := Student{}
	t := Teacher{}
	fmt.Println("student before", s)
	_ = fillStructByMap(&s, m)
	fmt.Println("student after", s)
	fmt.Println("teacher before", t)
	_ = fillStructByMap(&t, m)
	fmt.Println("teacher after", t)
}

func fillStructByMap(s interface{}, m map[string]interface{}) error {
	isPtr := reflect.TypeOf(s).Kind() == reflect.Ptr
	isStruct := reflect.TypeOf(s).Elem().Kind() == reflect.Struct
	if !isPtr || !isStruct {
		return errors.New("s should be a pointer of struct")
	}
	if m == nil {
		return errors.New("m must not be nil")
	}
	var (
		field	reflect.StructField
		ok		bool
	)
	for k, v := range m {
		//判断map的key在Struct中是否有对应属性名存在
		if field, ok = reflect.ValueOf(s).Elem().Type().FieldByName(k); !ok {
			continue
		}
		//存在对应属性名，判断Struct属性类型与map的value类型是否一致
		if field.Type == reflect.TypeOf(v) {
			value := reflect.ValueOf(s)
			//获取指针指向的Struct
			elem := value.Elem()
			//设置Struct属性值
			elem.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}