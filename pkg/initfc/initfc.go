package initfc

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func LoadFlightCtlConfigMap() map[string]string {
	// Create the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	configMap, err := clientset.CoreV1().ConfigMaps("flightctl-acm-addon").Get(context.TODO(), "flightctl-cm", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("ConfigMap Data: ", configMap.Data)
	fmt.Println("ConfigMap Data parts: ", configMap.Data)

	return configMap.Data
}
