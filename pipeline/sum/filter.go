package sum

import (
	"errors"
	"go/go_learn_series/pipeline/typedef"
)

var FilterError = errors.New("error data format in sum filter")

type Filter struct {}

func New() *Filter {
	return &Filter{}
}

func (f *Filter) Process(data typedef.Request) (typedef.Response, error) {
	elements, ok := data.([]int)
	if !ok {
		return nil, FilterError
	}
	ret := 0
	for _, element := range elements {
		ret += element
	}

	return ret, nil
}