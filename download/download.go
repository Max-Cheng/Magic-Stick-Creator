package download

import (
	"Magic-Stick-Creator/core"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func Download(url, cookie string) {
	if core.Exists("com.apple.recovery.boot/" + path.Base(url)) {
		err:=os.Remove("com.apple.recovery.boot/" + path.Base(url))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("正在下载"+path.Base(url)+"...")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Host = "oscdn.apple.com"
	req.Header.Set("Cookie", "AssetToken="+cookie)
	req.Header.Set("User-Agent", "InternetRecovery/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	out, err := os.Create("com.apple.recovery.boot/"+path.Base(url))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(path.Base(url)+"下载完成")
}