package generate

import (
	"math/rand"
	"net/http"
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
func get_image_info()  {
	
}