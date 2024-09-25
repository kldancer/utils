package k8s

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

const token = "ZXlKaGJHY2lPaUpTVXpJMU5pSXNJbXRwWkNJNklraGlTblJLVkUxS2EyWjJUSGRyUkc4d2RuZFdWMlZmTkhGQ2RYZFhUMFJxTldReVl6ZzFiM2g2YzBraWZRLmV5SnBjM01pT2lKcmRXSmxjbTVsZEdWekwzTmxjblpwWTJWaFkyTnZkVzUwSWl3aWEzVmlaWEp1WlhSbGN5NXBieTl6WlhKMmFXTmxZV05qYjNWdWRDOXVZVzFsYzNCaFkyVWlPaUprWldaaGRXeDBJaXdpYTNWaVpYSnVaWFJsY3k1cGJ5OXpaWEoyYVdObFlXTmpiM1Z1ZEM5elpXTnlaWFF1Ym1GdFpTSTZJbVJsWm1GMWJIUXRkRzlyWlc0dE5qUm9aSElpTENKcmRXSmxjbTVsZEdWekxtbHZMM05sY25acFkyVmhZMk52ZFc1MEwzTmxjblpwWTJVdFlXTmpiM1Z1ZEM1dVlXMWxJam9pWkdWbVlYVnNkQ0lzSW10MVltVnlibVYwWlhNdWFXOHZjMlZ5ZG1salpXRmpZMjkxYm5RdmMyVnlkbWxqWlMxaFkyTnZkVzUwTG5WcFpDSTZJakUzWm1FNFpEZzNMV1prTVRjdE5EQTVOaTA1TmpFNUxUTXpPR1ZoWlRJd04yWmlOeUlzSW5OMVlpSTZJbk41YzNSbGJUcHpaWEoyYVdObFlXTmpiM1Z1ZERwa1pXWmhkV3gwT21SbFptRjFiSFFpZlEuWW8wSE5PRk1DVDcybGNXZW9aWkFERnprU2pHN2Y3V2ZwSjh3MDU2d1NNUWozcEtVaUs1OTZENHJYZW5RNDl4cDZwRGlXQm9iN05EMUZ6eFhrZkVJMmZRMXhEZFZHUi02YkU5UFdYWEJZQ2JackhReU44ME8zd0FPX2ZtVXZOR3cySVpiNU1CbmFnVGNIaUdYUllxMU5yakxHRHBMT29PMFlTakVWSUhrb3I0b1owZzR3dnVDNGJSYk9CQml4MVpCc3M4bVFvZW5tNXB4WHlqdzQyYWtlMHowSXNWRHJ6VzZkQjZJbG1HYW5xeTFpQm93a01pY2plbVBIV3FUMEUzYk5UbjZoNkhibnVVZnVtaDNnZW5fTlZOUWFGWGdnc0tyNmwwc0kwQlZ1YjdxaUpTR0E0N2E1Y3BOVURUODZpWEpyY20tdG9kUnl3blZyRnVZT2JuUlVB"

func Informer() {
	//kubeconfig, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	//if err != nil {
	//	fmt.Printf("Error creating kubeconfig: %v\n", err)
	//	os.Exit(1)
	//}

	// 使用ServiceAccount Token初始化客户端
	// 如果你在Kubernetes集群内运行，可以直接使用InClusterConfig
	//config, err := rest.InClusterConfig()

	// 如果你在集群外运行，可以使用以下方式设置ServiceAccount Token
	config := &rest.Config{
		Host:        "https://172.28.8.138:6443",
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true, // 如果API服务器使用自签名证书，需要将Insecure设置为true
		},
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating clientset: %v\n", err)
		os.Exit(1)
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	configMapInformer := informerFactory.Core().V1().ConfigMaps().Informer()
	podInformer := informerFactory.Core().V1().Pods().Informer()

	configMapInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Infof("ConfigMap added: %s", obj.(*coreV1.ConfigMap).Name)

		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			log.Infof("ConfigMap updated %s", newObj.(*coreV1.ConfigMap).Name)
		},
		DeleteFunc: func(obj interface{}) {
			log.Infof("ConfigMap deleted %s", obj.(*coreV1.ConfigMap).Name)

		},
	})

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Infof("Pod added: %s", obj.(*coreV1.Pod).Name)
			if obj.(*coreV1.Pod).Name == "pod1" {
				log.Infof("aaaaa")
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			log.Infof("Pod updated %s", newObj.(*coreV1.Pod).Name)
			if newObj.(*coreV1.Pod).Name == "pod1" {
				log.Infof("uuuuuu")
			}
		},
		DeleteFunc: func(obj interface{}) {
			log.Infof("Pod deleted %s", obj.(*coreV1.Pod).Name)
			if obj.(*coreV1.Pod).Name == "pod1" {
				log.Infof("dddddd")
			}
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	// 启动 informer
	go informerFactory.Start(stopCh)

	// 等待程序退出
	<-wait.NeverStop
}
