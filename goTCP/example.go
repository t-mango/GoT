package goTCP

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TStart() {

	service := NewService(6000)
	go service.Start()

	reader := bufio.NewReader(os.Stdin)
	for {
		var cmd, args, parms string
		data, _, _ := reader.ReadLine()
		temp := string(data)

		list := strings.Split(temp, " ")
		fmt.Println(len(list))
		if len(list) == 2 {
			args = list[1]

		}
		if len(list) == 3 {
			args = list[1]
			parms = list[2]
		}
		cmd = list[0]

		switch cmd {
		case "list":
			service.GetList()
		case "send":
			if len(args) > 0 {
				if len(parms) > 0 {
					service.Send(args, parms)
					break
				}
				fmt.Println("请输入发送内容")
				break
			}
			fmt.Println("请输入key")
		case "sendAll":
			service.SendAll(args)
		case "showMsg":
			if len(args) > 0 {
				if len(parms) > 0 {
					service.ShowMsg(args, parms)
					break
				}
				fmt.Println("请输入发送内容")
				break
			}
			fmt.Println("请输入key")

		default:
			fmt.Println("命令错误")
		}
	}

}
