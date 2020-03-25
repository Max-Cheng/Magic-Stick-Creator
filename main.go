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
	var version int
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println( "║           macOS 在线安装U盘-制作工具             ║")
	fmt.Println( "║                         2020-03-22  by Co1a      ║")
	fmt.Println( "║                                                  ║")
	fmt.Println( "║      无需下载体积庞大的macOS镜像...              ║")
	fmt.Println( "║                体验稳如白果的在线安装...         ║")
	fmt.Println( "║                                                  ║")
	fmt.Println( "║  使用方法：                                      ║")
	fmt.Println( "║   1、准备Fat32格式的U盘，剩余空间大于500MB。     ║")
	fmt.Println( "║   2、将 运行NUC_Online_Installer、EFI-OC-        ║")
	fmt.Println( "║      xxxxxx.zip 拷贝至U盘根目录并解压EFI         ║")
	fmt.Println( "║   3、运行NUC_Online_Installer。                  ║")
	fmt.Println( "║   4、按照提示下载完成后，将U盘插上NUC。          ║")
	fmt.Println( "║   5、NUC开机，按住command/win键并持续点按R键。   ║")
	fmt.Println( "║   6、Cmd+R会激活Recovery模式启动，进入安装。     ║")
	fmt.Println( "║   7、Enjoy...                                    ║")
	fmt.Println( "║                                                  ║")
	fmt.Println( "║  如想了解更多 NUC8ixBE 豆子峡谷安装黑苹果知识，  ║")
	fmt.Println( "║          请查阅维奇的文章   http://u.nu/nuc8     ║")
	fmt.Println( "║                                                  ║")
	fmt.Println( "║   欢迎加入 Intel NUC Community Q群：341960876    ║")
	fmt.Println( "╚══════════════════════════════════════════════════╝")
	fmt.Println( "请选择你需要下载镜像的版本")
	fmt.Println("1.Last Version(最新版)")
	fmt.Println("2.Mojave")
	fmt.Println("3.High Sierra")
	fmt.Scanf("%d",&version);
	if version < 1 || version > 3 {
		fmt.Println("版本选择错误")
	}else{
		fmt.Printf( "需要先下载约500MB的Recovery镜像，是否开始? (Y/n)")
		fmt.Scanf("%s",&action);
		if action=="Y"||action=="y" {
			if core.Exists(config.Path) {
				os.Remove(config.Path)
			}
			os.Mkdir(config.Path,0755)
			fmt.Println("正在下载"+config.Select[version-1].Vs+"...")
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