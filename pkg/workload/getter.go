package workload

import (
	"context"

	"github.com/websi96/kubectl-kopy/pkg/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// func GetAllPVs(clientSet *kubernetes.Clientset) (volumes []api.VolumeClaim, err error) {
// 	volumeList, err := clientSet.CoreV1().PersistentVolumes().List(metaV1.ListOptions{})
// 	if err != nil {
// 		return
// 	}
// 	volumes = []api.VolumeClaim{}
// 	for _, v := range volumeList.Items {
// 		volumes = append(volumes, v)
// 	}
// 	return
// }

func GetAllDeployment(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}
	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, deploymentWorkload(w))
	}
	return
}

func GetAllDaemonSet(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().DaemonSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, daemonsetWorkload(w))
	}
	return
}

func GetAllStatefulset(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().StatefulSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, statefulsetWorkload(w))
	}
	return
}

func GetAllJobs(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.BatchV1().Jobs(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, jobWorkload(w))
	}
	return
}

func GetAllPods(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, podWorkload(w))
	}
	return
}
