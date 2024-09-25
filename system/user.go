package system

import (
	"fmt"
	"os/user"
)

func GetUser() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("无法获取当前用户信息:", err)
		return
	}

	fmt.Println("用户名:", currentUser.Username)
	fmt.Println("用户ID:", currentUser.Uid)
	fmt.Println("主组ID:", currentUser.Gid)
	fmt.Println("用户家目录:", currentUser.HomeDir)
}
