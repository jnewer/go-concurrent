package main

import (
	"fmt"
	"sync"
)

type Object struct {
	data string
}

func NewObject(data string) *Object {
	return &Object{data}
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return NewObject("default")
		},
	}

	obj := pool.Get().(*Object)
	fmt.Println(obj.data)

	obj.data = "foo"
	pool.Put(obj)

	obj = pool.Get().(*Object)
	fmt.Println(obj.data)
}
