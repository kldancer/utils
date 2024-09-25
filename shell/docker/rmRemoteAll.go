package docker

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"k8s.io/apimachinery/pkg/util/json"
)

func RmRemoteAll() {
	//imageName := getArgument("image_name")
	//tagName := getArgument("tag_name")
	imageName := "dev/kube-gpu"
	tagName := ""

	fmt.Printf("image_name:     %s\n", imageName)
	if tagName != "" {
		fmt.Printf("tag_name:       %s\n", tagName)
	}
	fmt.Println("********")

	var tags []string
	var err error

	if tagName != "" {
		tags = []string{tagName}
	} else {
		tags, err = getTags(imageName)
		if err != nil {
			fmt.Println("Error getting tags:", err)
			os.Exit(1)
		}
	}

	for _, tag := range tags {
		fmt.Println(" ")
		fmt.Println("--------------------------")

		digest, err := getDigest(imageName, tag)
		if err != nil {
			fmt.Println("Error getting digest:", err)
			continue
		}

		url := fmt.Sprintf("http://deploy.bocloud.k8s:40443/v2/%s/manifests/%s", imageName, digest)
		fmt.Printf("tag:            %s\n", tag)
		fmt.Printf("digest:         %s\n", digest)
		fmt.Printf("url:            %s\n", url)

		err = deleteImage(url)
		if err != nil {
			fmt.Println("Error deleting image:", err)
		}
	}

	fmt.Println("********")
	fmt.Printf("finish rm %s\n", imageName)
}

func getArgument(argName string) string {
	args := os.Args[1:]
	for i := 0; i < len(args)-1; i++ {
		if args[i] == fmt.Sprintf("%s", argName) {
			return args[i+1]
		}
	}
	return ""
}

func getTags(imageName string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("http://deploy.bocloud.k8s:40443/v2/%s/tags/list", imageName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tagsResponse struct {
		name string
		Tags []string `json:"tags"`
	}
	if err := json.Unmarshal(body, &tagsResponse); err != nil {
		return nil, err
	}

	return tagsResponse.Tags, nil
}

func getDigest(imageName, tag string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("HEAD", fmt.Sprintf("http://deploy.bocloud.k8s:40443/v2/%s/manifests/%s", imageName, tag), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Read and discard the progress information
	io.Copy(ioutil.Discard, resp.Body)

	digestHeader := resp.Header.Get("Docker-Content-Digest")
	return digestHeader, nil
}

func deleteImage(url string) error {

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	return nil
}
