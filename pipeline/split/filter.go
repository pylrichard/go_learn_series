package split

import (
	"errors"
	"strings"

	"go/go_learn_series/pipeline/typedef"
)

type Filter struct {
	delimiter string
}

var FilterError = errors.New("error data format in split filter")

func New(delimiter string) *Filter {
	return &Filter{delimiter: delimiter}
}

func (f *Filter) Process(data typedef.Request) (typedef.Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, FilterError
	}
	parts := strings.Split(str, f.delimiter)

	return parts, nil
}