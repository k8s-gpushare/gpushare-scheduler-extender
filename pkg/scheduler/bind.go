package scheduler

import (
	"github.com/AliyunContainerService/gpushare-scheduler-extender/pkg/cache"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kube-scheduler/extender/v1"
)

// Bind is responsible for binding node and pod
type Bind struct {
	Name  string
	Func  func(podName string, podNamespace string, podUID types.UID, node string, cache *cache.SchedulerCache) error
	cache *cache.SchedulerCache
}

// Handler handles the Bind request
func (b Bind) Handler(args v1.ExtenderBindingArgs) *v1.ExtenderBindingResult {
	err := b.Func(args.PodName, args.PodNamespace, args.PodUID, args.Node, b.cache)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return &v1.ExtenderBindingResult{
		Error: errMsg,
	}
}
