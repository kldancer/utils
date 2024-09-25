package k8s

import (
	"context"
	"fmt"
	"os"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const token2 = "eyJhbGciOiJSUzI1NiIsImtpZCI6Ii1CUnpZSjY0N0NGYjVNY0RmODBYOHk5U2pwNlVyYzh3bnRzdVpIbHh6MWcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJtYW5hZ2VtZW50LWFkbWluLXRva2VuLWxmY2t3Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Im1hbmFnZW1lbnQtYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIyZjQ3MmM4ZS1hNTc5LTRkNWEtOTBjYy1lMTk4ZGZlMjMzMTAiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06bWFuYWdlbWVudC1hZG1pbiJ9.AJirJx4YKEAGbjXAvhGTPCn-SDraVTCnQimvWA1rIK2excPeNfq-L6Xs7fOh3weZiNpl88m_2XMg23qtAHqRHf52ZO-YYL-t7rqWe6WyX8tC3-G84arQT_Gtci3oVnOXnJ7wPej65uMLVdITmBy4wcMFRVrxZs8195kYXfaYoQAiqP_C4pNgJpxp865akYtPcrPaDSsJwfBOmIDas-Jf0csRD2U6vbzl38vGDgxhU13_k36f-3QFyWnKBZRBpB9VJj4EBFx3WOVI-igA3ibMfCde8lB4EnuNocwY-cddXPFvigB3EOgNXW3TVnxm81JyNvqM3PIGp0_zhBed2aviQg"

func Informer2() {
	config := &rest.Config{
		Host:        "https://10.241.1.140:6443",
		BearerToken: token2,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true, // 如果API服务器使用自签名证书，需要将Insecure设置为true
		},
	}

	c, err := client.New(config, client.Options{})
	if err != nil {
		fmt.Printf("Error creating clientset: %v\n", err)
		os.Exit(1)
	}

	svc, err := GetService(c, "katib-ui", "kubeflow")
	if err != nil {
		fmt.Printf("Error getting service: %v\n", err)
		os.Exit(1)
	}

	nodePort := svc.Spec.Ports[0].NodePort

	fmt.Println("nodePort: ", nodePort)

}

func GetService(k client.Client, name, ns string) (*apiv1.Service, error) {
	service := &apiv1.Service{}
	if err := k.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: ns}, service); err != nil {
		return nil, err
	}
	return service, nil
}
