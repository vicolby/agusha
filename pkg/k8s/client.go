package k8s

import (
    "path/filepath"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/util/homedir"
    "k8s.io/client-go/tools/clientcmd"
)

func NewClientSet() (*kubernetes.Clientset, error) {
    var config *rest.Config
    var err error

    if home := homedir.HomeDir(); home != "" {
        kubeconfig := filepath.Join(home, ".kube", "config")
        config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
    } else {
        config, err = rest.InClusterConfig()
    }
    if err != nil {
        return nil, err
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }

    return clientset, nil
}

