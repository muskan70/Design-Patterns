package main

import (
	"fmt"
	"sync"
)

type ResourcePoolManager struct {
	FreeResources   []*Resource
	InUseResources  []*Resource
	InitialPoolSize int
	MaxPoolSize     int
	poolLock        *sync.Mutex
}

var resourcePoolManagerInstance *ResourcePoolManager
var lock = &sync.Mutex{}

func GetResourcePoolManagerInstance() *ResourcePoolManager {
	if resourcePoolManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if resourcePoolManagerInstance == nil {
			resourcePoolManagerInstance = &ResourcePoolManager{
				InitialPoolSize: 3,
				MaxPoolSize:     6,
				poolLock:        &sync.Mutex{},
			}
			for i := 0; i < resourcePoolManagerInstance.InitialPoolSize; i++ {
				resourcePoolManagerInstance.FreeResources = append(resourcePoolManagerInstance.FreeResources, GetResourceInstance())
			}
		}
	}
	return resourcePoolManagerInstance
}

func (r *ResourcePoolManager) GetResourceInstanceFromPool() *Resource {
	r.poolLock.Lock()
	defer r.poolLock.Unlock()
	if len(r.FreeResources) == 0 && len(r.InUseResources) < r.MaxPoolSize {
		resourcePoolManagerInstance.FreeResources = append(resourcePoolManagerInstance.FreeResources, GetResourceInstance())
	} else if len(r.FreeResources) == 0 && len(r.InUseResources) >= r.MaxPoolSize {
		fmt.Println("All resources are already in use")
		return nil
	}
	res := r.FreeResources[0]
	r.FreeResources = r.FreeResources[1:]
	r.InUseResources = append(r.InUseResources, res)
	return res
}

func (r *ResourcePoolManager) ReleaseResourceInstance(res *Resource) {
	r.poolLock.Lock()
	defer r.poolLock.Unlock()
	for i := 0; i < len(r.InUseResources); i++ {
		if r.InUseResources[i] == res {
			r.InUseResources = append(r.InUseResources[:i], r.InUseResources[i+1:]...)
		}
	}
	r.FreeResources = append(r.FreeResources, res)
}
