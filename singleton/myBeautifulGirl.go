package singleton

import (
	"fmt"
	"sync"
)

//由于go中没有类的概念，因此无法使用Java中的类的局部变量的方式，只能使用全局变量,然后将GetInstance方法export出去供访问
var (
	instance *myBeauGirl
	once     sync.Once
)

func GetInstance(name string) *myBeauGirl {
	if instance == nil {
		once.Do(func() {
			instance = &myBeauGirl{name: name}
		})
	}
	return instance
}

type myBeauGirl struct {
	name string
}

func (m myBeauGirl) ShowMyHeart() {
	fmt.Printf("%s是我心中的唯一\n", m.name)
}

type love interface {
	ShowMyHeart()
}
