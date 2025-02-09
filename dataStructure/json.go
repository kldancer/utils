package dataStructure

import (
	"context"
	"fmt"
	"github.com/pojozhang/sugar"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type JoinQosParam struct {
	RefNets []RefNet `json:"refNets"`
}
type RefNet struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	VmOwner   string `json:"vmOwner"`
}

func JsonMe() {
	unJoinQosParam := &JoinQosParam{
		RefNets: []RefNet{
			{Name: "net1", Namespace: "ns1", VmOwner: "vm1"},
			{Name: "net2", Namespace: "ns2", VmOwner: "vm2"},
		},
	}

	a := sugar.Json{Payload: *unJoinQosParam}
	fmt.Printf("a: %+v\n", a)
	fmt.Printf("*unJoinQosParam: %+v\n", *unJoinQosParam)
}

func SugarMe() {
	clusterId := "a1a4a7f4"
	qosId := "qos-x3iccqsg"

	url := fmt.Sprintf("http://10.210.20.152:11690/qos/%s/%s/unJoinQos", clusterId, qosId)
	log.Infof("unJoinQos.url: %s", url)

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJhZG1pbiIsIm1ldGhvZCI6IlRJQ0tFVCIsInVzZXJfaWQiOiIxZGY3MTk1N2Q3NzA0MzY0OTNhODRmMjRkOTM0NDg1MyIsInBhc3N3b3JkRXhwaXJlc0F0IjoiIiwiaXNzIjoiaHR0cHM6Ly8xMC4yMTAuMjAuMTU1IiwiZXhwIjoxNzM0NjEyNjk3LCJwcm9qZWN0TmFtZSI6InJvb3QiLCJpYXQiOjE3MzQ1OTgyOTcsInByb2plY3RJZCI6IjkwNTJhZTg1NzkyMTQzZmY5NTVjMzVjNTJlOGU0MWJkIiwidXNlcm5hbWUiOiJhZG1pbiJ9.gMqKSndFHAIWqM_K6GRHwcgUql2rZPuDZZA73OOF7zY"
	userInfo := "[{\"userId\":\"1df71957d770436493a84f24d9344853\",\"roleScope\":\"system\",\"roleName\":\"admin\",\"roleId\":\"bc446e28f2314ded942e6c38be0cc85a\",\"internal\":false,\"userName\":\"admin\",\"parentId\":null,\"domainId\":\"default\",\"projectsAndRoles\":[{\"role_name\":\"admin\",\"org_default\":true,\"project_id\":\"9052ae85792143ff955c35c52e8e41bd\",\"role_id\":\"bc446e28f2314ded942e6c38be0cc85a\",\"role_type\":\"system\",\"role_alias\":null,\"project_label\":\"root\",\"actor_id\":\"1df71957d770436493a84f24d9344853\",\"project_name\":\"root\"}],\"projectLabel\":\"root\",\"projectName\":\"root\",\"projectId\":\"9052ae85792143ff955c35c52e8e41bd\"}]"

	unJoinQosParam := &JoinQosParam{
		RefNets: []RefNet{
			{Name: "subnet-tcjyljcd", Namespace: "vm-87c173d0", VmOwner: "i-i1wkv9l9"},
		},
	}

	//jsonData, err := json.Marshal(unJoinQosParam)
	//if err != nil {
	//	fmt.Println("Error marshaling to JSON:", err)
	//	return
	//}

	response := sugar.Post(context.Background(), url, sugar.Json{Payload: unJoinQosParam}, sugar.Header{"X-Auth-Token": token, "userinfo": userInfo})
	if response.StatusCode == http.StatusOK {
		log.Infof("cluster %s unJoin qos %s success", clusterId, qosId)
	} else {
		err := fmt.Errorf("unjoin qos %s to cluster %s failed", qosId, clusterId)
		panic(err)
	}
}
