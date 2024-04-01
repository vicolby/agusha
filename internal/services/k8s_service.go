package services

import (
	"context"
	"fmt"

	"github.com/gimmefear/dswv3/internal/domain"
	"github.com/gimmefear/dswv3/pkg/k8s"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K8sService struct {
	client *kubernetes.Clientset
}

type DeploymentConfig struct {
	name      string
	namespace string
	labels    map[string]string
}

func int32Ptr(i int32) *int32 { return &i }

func NewK8sService() (*K8sService, error) {
	clientset, err := k8s.NewClientSet()
	if err != nil {
		return nil, err
	}
	return &K8sService{client: clientset}, nil
}

func (s *K8sService) createNamespace(ctx context.Context, namespace string) error {
	nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
	_, err := s.client.CoreV1().Namespaces().Create(ctx, nsSpec, metav1.CreateOptions{})
	return err
}

func (s *K8sService) deleteNamespace(ctx context.Context, namespace string) error {
	err := s.client.CoreV1().Namespaces().Delete(ctx, namespace, metav1.DeleteOptions{})
	return err
}

func (s *K8sService) createDeployment(ctx context.Context, workspace domain.Workspace) error {

	config := DeploymentConfig{
		name:      workspace.Name,
		namespace: fmt.Sprintf("project-%d", workspace.ProjectID),
		labels:    map[string]string{"app": "nginx"},
	}

	s.createNamespace(ctx, config.namespace)

	deploymentsClient := s.client.AppsV1().Deployments(config.namespace)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: config.name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: config.labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: config.labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		return err
	}

	return nil
}
