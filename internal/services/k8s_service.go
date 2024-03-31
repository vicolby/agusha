package services

import (
	"github.com/gimmefear/dswv3/pkg/k8s"
	"k8s.io/client-go/kubernetes"
)

type K8sService struct {
    client *kubernetes.Clientset
}

func NewK8sService() (*K8sService, error) {
    clientset, err := k8s.NewClientSet()
    if err != nil {
        return nil, err
    }
    return &K8sService{client: clientset}, nil
}
