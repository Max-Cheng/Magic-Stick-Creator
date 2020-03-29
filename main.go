package main

import (
	"Magic-Stick-Creator/config"
	"Magic-Stick-Creator/core"
	"Magic-Stick-Creator/generate"
	"fmt"
	"os"
	"time"

)

func main() {
	var action string
	var version int = 1
	if core.Exists("config.plist") {
		var mod int
		var ac string
		fmt.Printf("监测到存在plist，是否进行三码迁移(若plist中没有3码谨慎操作)[Y/n]")
		fmt.Scan(&ac)
		if ac == "y" || ac == "Y" {
			fmt.Println("请选择你的plist版本")
			fmt.Println("1.Clover")
			fmt.Println("2.OC")
			fmt.Scan(&mod)
			if mod<=0||mod>2 {
				fmt.Println("选择错误")
			}else{
				core.ChangeSN(mod)
			}
		}
		fmt.Println("退出三码三码转移")
	}
	fmt.Println( "请选择你需要下载镜像的版本")
	fmt.Println("1.Last Version(最新版)")
	fmt.Println("2.Mojave")
	fmt.Println("3.High Sierra")
	fmt.Scan(&version);
	if version < 1 || version > 3 {
		fmt.Println("版本选择错误")
	}else{
		fmt.Printf( "需要先下载约500MB的Recovery镜像，是否开始? [Y/n]")
		fmt.Scan(&action);
		if action=="Y"||action=="y" {
			if core.Exists(config.Path) {
				os.Remove(config.Path)
			}
			os.Mkdir(config.Path,0755)
			fmt.Println("选择了"+config.Select[version-1].Vs+"版本...")
			core.Download_image(generate.Get_image_info(config.Select[version-1]))
			Add_Details()
			fmt.Println("镜像下载完成")
		}else {
			fmt.Println("中止安装")
			return
		}
	}
	fmt.Println("将在3秒后自动退出")
	time.Sleep(3*1e9)
}
func Add_Details()  {
	fileName := ".contentDetails"
	if core.Exists(fileName) {
		err:=os.Remove(fileName)
		if err != nil {panic(err)}
	}
	dstFile,err := os.Create(config.Path+"/"+fileName)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s:="macOS Boot From Recovery"
	dstFile.WriteString(s)
}