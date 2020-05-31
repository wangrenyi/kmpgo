package kubernetes

import (
	"flag"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/client-go/rest"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubernetesClient *KubernetesClient
	kubeconfig       *string
	syncOnce         sync.Once
)

type KubernetesClient struct {
	config    *rest.Config
	clientSet *kubernetes.Clientset
}

func NewKubernetesClient() *KubernetesClient {
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &KubernetesClient{
		config:    config,
		clientSet: clientSet,
	}
}

func GetKubernetesClient() (*KubernetesClient, error) {
	syncOnce.Do(func() {
		kubernetesClient = NewKubernetesClient()
	})

	return kubernetesClient, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
