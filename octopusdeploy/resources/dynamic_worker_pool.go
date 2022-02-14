package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type DynamicWorkerPool struct {
	WorkerType WorkerType `json:"WorkerType"`

	WorkerPool
}

type DynamicWorkerPools struct {
	Items []*DynamicWorkerPool `json:"Items"`
	PagedResults
}

// NewDynamicWorkerPool creates and initializes a dynamic worker pool.
func NewDynamicWorkerPool(name string, workerType WorkerType) *DynamicWorkerPool {
	return &DynamicWorkerPool{
		WorkerType: workerType,
		WorkerPool: *octopusdeploy.newWorkerPool(name, WorkerPoolTypeDynamic),
	}
}

// GetWorkerType returns the worker type for this worker pool.
func (d *DynamicWorkerPool) GetWorkerType() WorkerType {
	return d.WorkerType
}

// Validate checks the state of the dynamic worker pool and returns an error if
// invalid.
func (d *DynamicWorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(d)
}

var _ IDynamicWorkerPool = &DynamicWorkerPool{}
