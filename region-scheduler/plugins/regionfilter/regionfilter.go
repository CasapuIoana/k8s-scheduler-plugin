package regionfilter

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var _ framework.FilterPlugin = &Plugin{}

const (
	PluginName = "RegionFilterPlugin"
)

type Plugin struct {
	handle framework.Handle
}

func (p *Plugin) Name() string {
	return PluginName
}

func (p *Plugin) Filter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	podRegion, ok := pod.GetLabels()["custom-label"]
	if !ok {
		return framework.NewStatus(framework.Success, "Pod does not specify a region")
	}

	nodeRegion, ok := nodeInfo.Node().GetLabels()["custom-label"]
	if !ok || podRegion != nodeRegion {
		return framework.NewStatus(framework.Unschedulable, "Node is in a different region")
	}

	return nil
}

func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &Plugin{}, nil
}
