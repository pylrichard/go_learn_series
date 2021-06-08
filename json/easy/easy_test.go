package easy

import (
	"fmt"
	"testing"

	"go/go_learn_series/json/typedef"
)


func BenchmarkTestEasyJson(b *testing.B) {
	b.ResetTimer()
	t := new(typedef.Teacher)
	for i := 0; i < b.N; i++ {
		err := t.UnmarshalJSON([]byte(typedef.JsonStr))
		if err != nil {
			fmt.Println("err easy json: ", err)
		}
		_, err = t.MarshalJSON()
		if err != nil {
			fmt.Println("err easy json: ", err)
		}
	}
}