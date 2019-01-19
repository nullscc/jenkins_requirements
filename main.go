package main

import (
	"fmt"
	"github.com/nullscc/jenkins_requirements/routers"
	"github.com/nullscc/jenkins_requirements/utils"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		routers.Run(":8000")
	} else if os.Args[1] == "post" {
		_, err := utils.PostSha1(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("成功")
	} else if os.Args[1] == "get" {
		sha1, err := utils.GetSha1(os.Args[2], os.Args[3], os.Args[4])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(sha1)

	} else {
		fmt.Println("参数错误")
	}
}
