package model

import (
	"encoding/hex"
	"github.com/iEvan-lhr/apaa-for-go/structs"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"golang.org/x/crypto/md4"
	"time"
)

func (u *UserModel) generateUC(user *structs.User) bool {
	if ok, key := decrypt(tools.Strings(user.DenKey), tools.Strings(user.Identity)); !ok {
		return false
	} else if key < time.Now().Unix() {
		user.DenKey, user.Identity = encryption(time.Now().Add(15 * time.Minute).Unix())
	}
	return true
}

func decrypt(str1, str2 *tools.String) (bool, int64) {
	if !str2.Check(checkHax(str1.Bytes())) {
		return false, 0
	} else {
		ans := &tools.String{}
		for i := 0; i < len(str1.Bytes()); i += 5 {
			ans.Append(str1.Bytes()[i] - 'c' + 47)
		}
		timeLabel, err := ans.Atoi()
		anything.ErrorDontExit(err)
		return true, int64(timeLabel)
	}
}

var cloud = "cloud"

func encryption(timeNow int64) (string, string) {
	str := tools.Itoa(timeNow)
	crypt := &tools.String{}
	for i := range str.Bytes() {
		var l []byte
		for j := range cloud {
			l = append(l, cloud[j]+(str.Bytes()[i]-47))
		}
		crypt.Append(l)
	}
	return crypt.String(), encodeHax(md4.New().Sum(crypt.Bytes()))
}

func checkHax(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}
func encodeHax(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}
