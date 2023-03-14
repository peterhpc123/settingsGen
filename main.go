package main

import (
	"awesomeProject/cmd"
	_ "awesomeProject/cmd"
	_ "github.com/beevik/etree"
)

func main() {
	//1.构建cobra命令行工具
	cmd.Execute()
	//2.实现命令行逻辑
	//2.1 去github下载settings文件
	//2.2 访问远程仓库的apikey
	//3.修改settingsxml文件

}
