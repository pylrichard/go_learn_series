package typedef

//Request 为了接收任意类型的传入参数，定义Request为interface
type Request interface {}

type Response interface {}

type Filter interface {
	Process(data Request) (Response, error)
}