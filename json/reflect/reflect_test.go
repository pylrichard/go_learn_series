package reflect

import (
	"encoding/json"
	"fmt"
	"testing"

	"go/go_learn_series/json/typedef"
)

func BenchmarkTestReflectJson(b *testing.B) {
	b.ResetTimer()
	t := new(typedef.Teacher)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(typedef.JsonStr), &t)
		if err != nil {
			fmt.Println("err json: ", err)
		}
		_, err = json.Marshal(t)
		if err != nil {
			fmt.Println("err json: ", err)
		}
	}
}