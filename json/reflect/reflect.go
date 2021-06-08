package reflect

import (
	"encoding/json"
	"fmt"

	"go/go_learn_series/json/typedef"
)

/*
	json解析内部是使用反射机制reflect实现
 */
func Json() {
	t := new(typedef.Teacher)
	err := json.Unmarshal([]byte(typedef.JsonStr), &t)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("teacher: ", t)
	}
	//返回[]byte类型，需要将其转成字符串
	jsonStr, err := json.Marshal(t)
	if err != nil {
		fmt.Println("err: ®", err)
	} else {
		fmt.Println("json string: ", string(jsonStr))
	}
}