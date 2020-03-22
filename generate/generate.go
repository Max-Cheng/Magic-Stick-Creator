package generate

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
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

func Get_image_info() map[string]string {
	url := "http://osrecovery.apple.com/InstallationPayload/RecoveryImage"
	method := "POST"
	payload := strings.NewReader("cid=EE3029997629CF05\nk=EBDDC2B8CCE2493D7E9380F82C04999495C0724ED220D414AA3AB41E0DBD9D9A\nsn=00000000000000000\nfg=E723289F4F30BB3FED84594C9BED7EB7C52C614359A526D317D99FDF459A066B\nos=default\nbid=Mac-7BA5B2D9E42DDD94")
	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Host", "osrecovery.apple.com")
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "InternetRecovery/1.0")
	req.Header.Add("Cookie", Get_session())
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "d375c5bb-5a92-4a8b-8f9c-41083df18781")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Length", "212")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	newbody:=(strings.Split(string(body),"\n"))
	result:= map[string]string{}
	for i:=0;i<len(newbody)-1;i++{
		result[newbody[i][:2]]=newbody[i][4:len(newbody[i])]
	}
	return result
}