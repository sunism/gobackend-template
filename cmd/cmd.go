package main

import (
	"fmt"
	"s3video/core"
	"s3video/gloabl"
)
//主函数，启动程序
func main(){
	gloabl.GVA_VP = core.Viper()
	core.InitLog()

	fmt.Printf("hello world")
}