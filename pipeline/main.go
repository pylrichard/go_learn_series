package main

import (
	"fmt"

	"go/go_learn_series/pipeline/convert"
	"go/go_learn_series/pipeline/pipeline"
	"go/go_learn_series/pipeline/split"
	"go/go_learn_series/pipeline/sum"
)

func main() {
	splitFilter := split.New(",")
	convertFilter := convert.New()
	sumFilter := sum.New()
	pipeFilter := pipeline.New("pipeline",
							splitFilter, convertFilter, sumFilter)
	ret, err := pipeFilter.Process("1,2,3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}