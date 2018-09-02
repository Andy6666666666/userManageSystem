package main

import (
	"userManageSystem/models"
	_ "userManageSystem/routers"

	"github.com/astaxie/beego"
)

func main() {
	if b, _ := models.PathExists(models.FILE_SMS_MSGS); b {
		if err := models.ReadSMS(&models.MsgManage); err != nil {
			beego.Info("read sms:" + err.Error())
		} else {
			beego.Info("read sms ok")
		}
	}

	if b, _ := models.PathExists(models.FILE_USERS); b {
		if err := models.ReadUsers(&models.Accounts); err != nil {
			beego.Info("read users:" + err.Error())
		} else {
			beego.Info("read users ok")
		}
	}

	beego.Run()
}
