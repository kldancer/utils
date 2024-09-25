package system

import (
	"fmt"
	"os"
)

func GetEnv() {
	a := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	b := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")

	fmt.Println(a)
	fmt.Println(b)

}

func main() {
	GetEnv()
}
