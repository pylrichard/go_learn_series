package convert

import (
	"errors"
	"strconv"

	"go/go_learn_series/pipeline/typedef"
)

type Filter struct {}

var FilterError = errors.New("error data format in convert filter")

func New() *Filter {
	return &Filter{}
}

func (f *Filter) Process(data typedef.Request) (typedef.Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, FilterError
	}
	var ret []int
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}

	return ret, nil
}