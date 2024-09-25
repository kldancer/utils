package cloud

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	log "github.com/sirupsen/logrus"
)

func Test() {
	ecsClient, err := ecs.NewClientWithAccessKey("cn-wulan-ste3-d01", "JNcyZjijSoVUdnC8", "Q09POF4wTMfJ6vlAJolMpcj9ZTtOub")
	if err != nil {
		log.Errorf("failed to create ecs client: %v", err)
		return
	}

	ecsClient.GetConfig().WithScheme("HTTPS")
	ecsClient.SetHTTPSInsecure(true)
	ecsClient.Domain = "ecs.cloud.ste3.com"

	vpcClient, err := vpc.NewClientWithAccessKey("cn-wulan-ste3-d01", "JNcyZjijSoVUdnC8", "Q09POF4wTMfJ6vlAJolMpcj9ZTtOub")
	if err != nil {
		log.Errorf("failed to create vpc client: %v", err)
	}

	vpcClient.GetConfig().WithScheme("HTTPS")
	vpcClient.SetHTTPSInsecure(true)
	vpcClient.Domain = "vpc.cloud.ste3.com"

	describeNetworkInterfacesRequest(ecsClient)
	describeInstances(ecsClient)
	describeInstanceTypes(ecsClient)
	describeVpcAttribute(vpcClient, "cn-wulan-ste3-d01", "vpc-8x1qgq0n")

}

func describeNetworkInterfacesRequest(ecsClient *ecs.Client) {
	log.Infoln()
	log.Infof("describeNetworkInterfacesRequest")
	req := ecs.CreateDescribeNetworkInterfacesRequest()
	req.InstanceId = "i-a4c018p4xjouqjwlqbrn"
	//req.Type = "Primary"
	resp, err := ecsClient.DescribeNetworkInterfaces(req)
	if err != nil {
		log.Errorf("failed to describe network interfaces: %v", err)
		return
	}

	log.Infof("req: %+v", req)
	log.Infof("resp: %+v", resp)

	for _, netInterface := range resp.NetworkInterfaceSets.NetworkInterfaceSet {
		fmt.Printf("netInterface.SecurityGroupIds.SecurityGroupId: %v\n", netInterface.SecurityGroupIds.SecurityGroupId)
	}

	log.Info("end")
}

func describeVpcAttribute(vpcClient *vpc.Client, RegionId, VpcId string) {
	log.Infoln()
	log.Infof("describeVpcAttribute")
	req := vpc.CreateDescribeVpcAttributeRequest()
	req.RegionId = RegionId
	req.VpcId = VpcId

	log.Infof("req: %+v", req)
	resp, err := vpcClient.DescribeVpcAttribute(req)
	if err != nil {
		log.Errorf("failed to describe vpc attribute: %v", err)
	}

	log.Infof("resp: %+v", resp)

	log.Infof("end")

}

func describeInstances(ecsClient *ecs.Client) {
	log.Infoln()
	log.Infof("describeInstances")
	req := ecs.CreateDescribeInstancesRequest()
	//req.InstanceIds = "[\"i-a4c018p4xjouqjwlqbrn\"]"

	log.Infof("req: %+v", req)
	resp, err := ecsClient.DescribeInstances(req)
	if err != nil {
		log.Errorf("failed to describe instances: %v", err)
	}

	log.Infof("resp: %+v", resp)
	log.Infof("end")

}

func describeInstanceTypes(ecsClient *ecs.Client) {
	log.Infoln()
	log.Infof("describeInstanceTypes")
	req := ecs.CreateDescribeInstanceTypesRequest()
	req.InstanceTypes = &[]string{"ecs.n4v2.4xlarge"}
	resp, err := ecsClient.DescribeInstanceTypes(req)
	if err != nil {
		log.Errorf("failed to describe instance types: %v", err)
	}
	log.Infof("resp: %+v", resp)
	log.Infof("end")

}
