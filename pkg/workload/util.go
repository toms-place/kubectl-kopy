package workload

import (
	"strings"

	"github.com/websi96/kubectl-kopy/pkg/api"
)

func Join(workloads []api.Workload, sep string) string {
	allNames := []string{}

	for _, wk := range workloads {
		allNames = append(allNames, wk.GetName())
	}
	return strings.Join(allNames, sep)
}
