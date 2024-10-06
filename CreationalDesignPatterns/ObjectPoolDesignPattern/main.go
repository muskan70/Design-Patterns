package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	resourcePool := GetResourcePoolManagerInstance()

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res := resourcePool.GetResourceInstanceFromPool()
			if res != nil {
				fmt.Printf("ResourceId : %d acquired by GoRoutine : %d\n", res.ResourceId, i)
				time.Sleep(time.Second)
				fmt.Printf("ResourceId : %d released by GoRoutine : %d\n", res.ResourceId, i)
				resourcePool.ReleaseResourceInstance(res)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Resource Allocation work done!!")
}
