package main

import (
	"context"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type PodInfo struct {
	PodName    string
	Containers []string
	Namespace  string
}

// GetClient returns a new client session with the Kubernetes cluster
func GetClient(kubeConfig *rest.Config) *kubernetes.Clientset {
	client, err := kubernetes.NewForConfig(kubeConfig)
	CheckError(err, "Could not generate a Kubernetes client")

	return client
}

// GetKubeconfig returns the cluster configuration
func GetKubeconfig() *rest.Config {
	kubeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	CheckError(err, "Could not obtain or read the local .kube/config file")

	return kubeConfig
}

// GetPods returns a list of pods
func GetPods(client *kubernetes.Clientset) *v1.PodList {
	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	CheckError(err, "Could not retrieve pods from all namespaces")

	return pods
}
