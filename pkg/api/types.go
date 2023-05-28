package api

import v1 "k8s.io/api/core/v1"

type Reason string

const (
	NO_RESOURCE_REFFERENCE     Reason = "No Reference"
	WORKLOAD_HAS_ZERO_REPLICAS Reason = "Zero Replica"
)

type Workload interface {
	GetName() string
	GetVolumeClaimNames() []string
	// IsEmpty tells that the workload still has pod running
	IsEmpty() bool
}

type VolumeClaim struct {
	v1.PersistentVolumeClaim
	Reason    Reason
	Workloads []Workload
}

type Volume struct {
	v1.PersistentVolume
	Reason    Reason
	Workloads []Workload
}
