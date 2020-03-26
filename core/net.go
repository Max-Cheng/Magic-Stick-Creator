package core

import (
	"Magic-Stick-Creator/config"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	_ "gopkg.in/cheggaaa/pb.v1"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
func Download_image(UrlAndCookie map[string]string){
	HttpHandle(UrlAndCookie["CU"],"GET",map[string]string{"Cookie":"AssetToken="+UrlAndCookie["CT"]},3,nil)
	HttpHandle(UrlAndCookie["AU"],"GET", map[string]string{"Cookie":"AssetToken="+UrlAndCookie["AT"]},3,nil)
}

func HttpHandle(url,method string,cookie map[string]string,mod int,Body io.Reader) interface{}{
	//get_session mod:1,get_image_info:2,downloadmod:3
	client:=&http.Client{}
	reqest, err := http.NewRequest(method, url, Body)
	if err != nil {
		fmt.Println(err)
	}
	if len(cookie)!=0 {
		for i,v:=range cookie{
			reqest.Header.Add(i,v)
		}
	}
	reqest.Header.Add("Host", "osrecovery.apple.com")
	reqest.Header.Add("Connection", "close")
	reqest.Header.Add("User-Agent", "InternetRecovery/1.0")
	reqest.Header.Add("Content-Type", "text/plain")
	reqest.Header.Add("Connection", "close")
	reqest.Header.Add("Accept", "*/*")
	reqest.Header.Add("Cache-Control", "no-cache")
	reqest.Header.Add("Postman-Token", "d375c5bb-5a92-4a8b-8f9c-41083df18781")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate, br")
	reqest.Header.Add("Content-Length", "212")
	if err != nil {
		panic(err)
	}
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	switch mod {
	case 1:
		return response.Header.Get("Set-Cookie")[:83];break
	case 2:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		newbody:=(strings.Split(string(body),"\n"))
		result:= map[string]string{}
		for i:=0;i<len(newbody)-1;i++{
			result[newbody[i][:2]]=newbody[i][4:len(newbody[i])]
		}
		return result;break
	case 3:
		if Exists(config.Path+"/" + path.Base(url)) {
			err:=os.Remove(config.Path+"/" + path.Base(url))
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("正在下载"+path.Base(url)+"...")
		i, _ := strconv.Atoi(response.Header.Get("Content-Length"))
		sourceSiz := int64(i)
		source := response.Body
		bar := pb.New(int(sourceSiz)).SetUnits(pb.U_BYTES_DEC).SetRefreshRate(time.Millisecond * 10)
		bar.ShowSpeed = true
		bar.ShowTimeLeft = true
		bar.ShowFinalTime = true
		width,_:=pb.GetTerminalWidth()
		bar.SetWidth(width)
		bar.Start()
		out, err := os.Create(config.Path+"/"+path.Base(url))
		if err != nil {
			panic(err)
		}
		defer out.Close()
		writer := io.MultiWriter(out, bar)
		io.Copy(writer, source)
		bar.Finish()
		fmt.Println(path.Base(url)+"下载完成")
	}
	return ""
}
