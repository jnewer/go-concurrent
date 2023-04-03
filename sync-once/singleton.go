package singleton

import "sync"

type Singleton struct {
	state int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{state: 0}
	})

	return instance
}

func main() {

}
