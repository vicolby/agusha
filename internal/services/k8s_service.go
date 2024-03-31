package services

import (
    "context"
	"github.com/gimmefear/dswv3/pkg/k8s"
	"k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    apiv1 "k8s.io/api/core/v1"
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

func (s *K8sService) createNamespace(namespace string) error {
    nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
    _, err := s.client.CoreV1().Namespaces().Create(context.TODO(), nsSpec, metav1.CreateOptions{})
    return err
}

