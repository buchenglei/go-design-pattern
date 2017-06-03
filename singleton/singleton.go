// 这个是使用标准库实现的方法
package singleton

import (
	"sync"
)

var once sync.Once
var singleton *Singleton

type Singleton struct {
	Name string
}

// SingletonInstance 获取一个单例对象
func SingletonInstance(name string) *Singleton {
	once.Do(func() {
		// 这里的初始化操作只会被执行一次
		singleton = &Singleton{
			Name: name,
		}
	})

	return singleton
}
