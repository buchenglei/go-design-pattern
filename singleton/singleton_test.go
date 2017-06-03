package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingletonInstance(t *testing.T) {
	var wg sync.WaitGroup

	// 初始化一次取一个基准值
	instance := SingletonInstance("No.110")

	// 即使在并发的情况下也会保证只打印一个值
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			name := fmt.Sprintf("No.%d", i)
			instance = SingletonInstance(name)
			if instance.Name != "No.110" {
				t.Errorf("Expect No.110, but get %s", instance.Name)
			}
		}()
	}

	wg.Wait()
}
