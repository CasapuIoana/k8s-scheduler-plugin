package main

import (
	"fmt"
	"os"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"regionscheduler/plugins/regionfilter"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(regionfilter.PluginName, regionfilter.New),
	)
	
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
