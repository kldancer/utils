package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GetMethod(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func DoMethod(rawURL string) string {
	// 解析 URL 字符串
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return ""
	}

	httpCli := http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := httpCli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func GetMethod2(rawURL string) string {
	// 解析 URL 字符串
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	httpCli := http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := httpCli.Get(parsedURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
