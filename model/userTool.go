package model

import (
	"github.com/iEvan-lhr/apaa-for-go/structs"
	build "github.com/iEvan-lhr/encryption-apaa"
	tools "github.com/iEvan-lhr/exciting-tool"
	"time"
)

func (u *UserModel) generateUC(user *structs.User) bool {
	if ok, key := build.Decrypt(tools.Strings(user.DenKey), tools.Strings(user.Identity)); !ok {
		return false
	} else if key < time.Now().Unix() {
		user.DenKey, user.Identity = build.Encryption(time.Now().Add(15 * time.Minute).Unix())
	}
	return true
}

func (u *UserModel) checkUserLegitimate(user *structs.User) bool {
	if structs.Sure(&structs.DenKeyByRegister{DenKey: user.DenKey}) && structs.Check(user) {
		if len(user.Password) > 7 {
			// TODO 判断密码强弱
		} else {
			return false
		}
		return true
	} else {
		return false
	}
}
