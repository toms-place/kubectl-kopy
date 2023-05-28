package plugin

import (
	"context"

	"github.com/websi96/kubectl-kopy/pkg/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetVolumes(clientSet *kubernetes.Clientset) (volumes []*api.Volume, err error) {
	list, err := clientSet.CoreV1().PersistentVolumes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}
	for _, pv := range list.Items {
		volumes = append(volumes, &api.Volume{
			PersistentVolume: pv,
			Reason:           api.NO_RESOURCE_REFFERENCE,
		})
	}
	return
}

func GetVolumeClaims(clientSet *kubernetes.Clientset, namespace string) (volumes []*api.VolumeClaim, err error) {
	list, err := clientSet.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return
	}
	for _, pvc := range list.Items {
		volumes = append(volumes, &api.VolumeClaim{
			PersistentVolumeClaim: pvc,
			Reason:                api.NO_RESOURCE_REFFERENCE,
		})
	}
	return
}
