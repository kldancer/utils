package http

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://10.245.0.138:9013/ping/test-server-node-port.fabric-e2e.svc.cluster.local:30133"
	//url := "http://10.245.0.138:9013/ping/10.245.0.39:9013"
	//url := "http://10.245.0.138:9013/ping/10.245.1.177:9013"
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result := GetMethod(url)
			//result := DoMethod(url)
			//result := GetMethod2(url)
			fmt.Printf("result%d = %+v\n", i, result)
		}(i)
	}
	wg.Wait()
}
