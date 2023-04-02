package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID   int
	Name string
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return new(User)
		},
	}

	user := pool.Get().(*User)
	user.ID = 1
	user.Name = "Alice"
	fmt.Println(*user)

	pool.Put(user)

	user = pool.Get().(*User)
	fmt.Println(*user)
}
