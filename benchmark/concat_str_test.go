package main

import (
	"bytes"
	"testing"
)

/*
	go test -bench=.
 */
func BenchmarkConcatStrByAdd(b *testing.B) {
	elements := []string{ "1", "2", "3", "4" }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, element := range elements {
			ret += element
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStrByBytesBuffer(b *testing.B) {
	elements := []string{ "1", "2", "3", "4" }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, element := range elements {
			buf.WriteString(element)
		}
	}
	b.StopTimer()
}