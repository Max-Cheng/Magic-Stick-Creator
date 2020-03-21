package generate

import (
	"math/rand"
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