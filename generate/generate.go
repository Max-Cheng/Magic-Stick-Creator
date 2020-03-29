package generate

import (
	"Magic-Stick-Creator/config"
	"Magic-Stick-Creator/core"
	"fmt"
	"math/rand"
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
func Get_image_info(ver config.Vertion)map[string]string{
	url := "http://osrecovery.apple.com/InstallationPayload/RecoveryImage"
	method := "POST"
	version:="cid="+Genedate_id(16)+"\n"+"k="+Genedate_id(64)+"\n"+"sn="+ver.SN+"\n"+"fg="+Genedate_id(64)+"\n"+"os=default\n"+"bid="+ver.BID
	payload := strings.NewReader(version)
	return core.HttpHandle(url,method, map[string]string{"Cookie":Get_session()},2,payload).(map[string]string)
}
func Get_session()string{
	return fmt.Sprintf("%s",core.HttpHandle("http://osrecovery.apple.com/","GET",map[string]string{},1,nil))
}