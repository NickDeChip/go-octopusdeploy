package resources

import (
	"github.com/jinzhu/copier"
)

func ToWorkerPool(workerPoolResource *WorkerPoolResource) (IWorkerPool, error) {
	// if isNil(workerPoolResource) {
	// 	return nil, createInvalidParameterError("ToWorkerPool", "workerPoolResource")
	// }

	var workerPool IWorkerPool
	var err error
	switch workerPoolResource.GetWorkerPoolType() {
	case WorkerPoolTypeDynamic:
		workerPool = NewDynamicWorkerPool(workerPoolResource.GetName(), workerPoolResource.GetWorkerType())
	case WorkerPoolTypeStatic:
		workerPool = NewStaticWorkerPool(workerPoolResource.GetName())
	}

	err = copier.Copy(workerPool, workerPoolResource)
	if err != nil {
		return nil, err
	}

	return workerPool, nil
}

func ToWorkerPoolResource(workerPool IWorkerPool) (*WorkerPoolResource, error) {
	// if isNil(workerPool) {
	// 	return nil, createInvalidParameterError("ToWorkerPoolResource", "workerPool")
	// }

	workerPoolResource := NewWorkerPoolResource(workerPool.GetName(), workerPool.GetWorkerPoolType())

	err := copier.Copy(&workerPoolResource, workerPool)
	if err != nil {
		return nil, err
	}

	return workerPoolResource, nil
}

func ToWorkerPoolArray(workerPoolResources []*WorkerPoolResource) []IWorkerPool {
	items := []IWorkerPool{}
	for _, workerPoolResource := range workerPoolResources {
		workerPool, err := ToWorkerPool(workerPoolResource)
		if err != nil {
			return nil
		}
		items = append(items, workerPool)
	}
	return items
}
