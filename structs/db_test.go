package structs

import (
	//"crypto/md4"
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"log"
	"strconv"
	testing2 "testing"
	"time"
)

func TestDB(t *testing2.T) {
	start := time.Now()
	str1 := strconv.Itoa(int(time.Now().Unix()))
	atoi, _ := strconv.Atoi(str1)
	code, _ := encryption(atoi)
	var ans []byte
	for i := 0; i < len(code); i += 5 {
		ans = append(ans, code[i]-'c'+47)
	}
	log.Println(str1)
	log.Println(string(ans))
	log.Println(time.Now().Sub(start))
	//log.Println(encryption(int(time.Now().Unix())))
	//count, err := engine.Count(any(&User{}))
	//log.Println(count, err)
}

var cloud = "cloud"

func encryption(timeNow int) (string, string) {
	str := strconv.Itoa(timeNow)
	crypt := ""
	for i := range str {
		var l []byte
		for j := range cloud {
			l = append(l, cloud[j]+(str[i]-47))
		}
		crypt += string(l)
	}
	return crypt, hex.EncodeToString(md4.New().Sum([]byte(crypt)))
}
