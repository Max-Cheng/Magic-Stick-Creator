package generate

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

func Genedate_id(x int)string{
	valid_chars := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	id:=[]byte{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < x; i++ {
		id=append(id,valid_chars[rand.Intn(len(valid_chars))])
	}
	return string(id)
}
func Get_session() string {
	url := "http://osrecovery.apple.com/"
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Set("Host","osrecovery.apple.com")
	reqest.Header.Set("Connection","close")
	reqest.Header.Set("User-Agent","InternetRecovery/1.0")
	if err != nil {
		panic(err)
	}
	response, err := (&http.Client{}).Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//This Function maybe wrong if can't get session edit here⬇️
	return response.Header.Get("Set-Cookie")[:83]
}

func Get_image_info()  {
	url:="http://osrecovery.apple.com/InstallationPayload/RecoveryImage"
	data:=url2.Values{}
	data.Add("cid",Genedate_id(16))
	data.Add("k",Genedate_id(64))
	data.Add("sn","00000000000000000")
	data.Add("fg",Genedate_id(64))
	data.Add("os","default")
	data.Add("bid","Mac-"+Genedate_id(16))
	reqest,err:=http.NewRequest("POST",url,strings.NewReader(data.Encode()))
	if err!=nil{panic(err)}
	reqest.Header.Set("Host","osrecovery.apple.com")
	reqest.Header.Set("User-Agent","InternetRecovery/1.0")
	reqest.Header.Set("Cookie",Get_session())
	reqest.Header.Set("Expect","")
	reqest.Header.Set("Content-Type","text/plain")
	reqest.Header.Set("Connection","close")


	response, err := (&http.Client{}).Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body_byte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body_byte))
}