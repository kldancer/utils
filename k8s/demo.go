package k8s

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func NewRestClient() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	// RESTClientFor takes a config and a group, and returns a RESTClient.
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	pod := v1.Pod{}
	// 其实就是拼接HTTP请求的URL
	if err := restClient.Get().Namespace("default").Resource("pods").Name("nginx").Do(context.TODO()).
		Into(&pod); err != nil {
		panic(err)
	}

}

func NewClientSet() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	// client集合
	// 根据API、VERSION对client进行分组
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podClient := clientset.CoreV1().Pods("default")
	if _, err := podClient.Get(context.TODO(), "nginx", metav1.GetOptions{}); err != nil {
		panic(err)
	}
}

func NewDynamicClient() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	// dynamicClient 操作任意K8s资源，包括CRD
	_, err = dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}

func NewDiscoveryClient() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	// discoveryClient 用于返现K8s提供的资源组、资源版本、资源信息。比如kubectl api-resources
	_, err = discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}
}

func NewInformer() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// informer
	factory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	_, err = factory.Core().V1().Pods().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    nil,
		UpdateFunc: nil,
		DeleteFunc: nil,
	})
	stopCh := make(chan struct{})
	factory.Start(stopCh)
}
