package main

/*
	小功能接口定义
 */
type Reader interface {
	Read(buf []byte) (int, error)
}

type Writer interface {
	Write(buf []byte) (int, error)
}

//大接口使用小接口嵌套组装
type IO interface {
	Reader
	Writer
}