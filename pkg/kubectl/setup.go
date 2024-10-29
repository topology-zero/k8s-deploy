package kubectl

import (
	"k8s-deploy/config"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var K8sClient *kubernetes.Clientset

func Setup() {
	var err error
	var restConfig *rest.Config
	if config.Env == "prod" {
		restConfig, err = rest.InClusterConfig()
	} else {
		restConfig, err = clientcmd.BuildConfigFromFlags("", homedir.HomeDir()+"/.kube/config")
	}
	if err != nil {
		logrus.Fatal(err)
	}

	K8sClient, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		logrus.Fatal(err)
	}
}
