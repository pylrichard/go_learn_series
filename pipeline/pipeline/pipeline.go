package pipeline

import "go/go_learn_series/pipeline/typedef"

type Pipeline struct {
	Name	string
	Filters *[]typedef.Filter
}

func New(name string, filters ...typedef.Filter) *Pipeline {
	return &Pipeline{
		Name: name,
		Filters: &filters,
	}
}

func (f *Pipeline) Process(data typedef.Request) (typedef.Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}

	return ret, err
}