package controller

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func CreateClient(kubeConfigPath string) (*kubernetes.Clientset ,error){
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		klog.Errorf("load kubeConfig from %s failed, err: %s", kubeConfigPath, err.Error())
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Errorf("create k8s client failed, err: %s", err.Error())
		return nil, err
	}
	return clientSet, nil
}