package controllers

import (
	"fmt"
	"strings"
	"userManageSystem/models"

	"github.com/astaxie/beego"
)

//登录
func (this *AccountController) login() {

	//从前端获取数据信息
	/*var user models.UserInfo
	user.Mobile = this.Input().Get(models.DATA_KEY_MOBILE)
	beego.Info(user.Mobile)

	user.Password = this.Input().Get(models.DATA_KEY_PASSWORD)
	beego.Info(user.Password)
	*/
	//SetLoginChan(user)

	pwd := this.Input().Get(models.DATA_KEY_PASSWORD)
	mobile := this.Input().Get(models.DATA_KEY_MOBILE)
	beego.Info(mobile)
	beego.Info(pwd)

	user, err := models.GetUserByMobile(mobile)
	if err != nil {
		//this.Ctx.WriteString(err.Error())
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error()}
		this.ServeJSON()
		return
	}

	if strings.Compare(user.Password, pwd) != 0 {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": models.ErrPwd.Error()}
		this.ServeJSON()
		return
	}

	user.IsLogin = true
	models.UpdateUser(user)

	this.Data["json"] = map[string]interface{}{"status": 0, "msg": fmt.Sprintf("Login succeed,mobile is %s", user.Mobile)}
	this.ServeJSON()
	return
}

//注册
func (this *AccountController) register() {
	var user models.UserInfo

	user.Mobile = this.Input().Get(models.DATA_KEY_MOBILE)
	beego.Info(user.Mobile)

	if models.IsExistUser(user.Mobile) {
		//this.Ctx.WriteString(models.ErrUserAlreadyExists.Error())

		this.Data["json"] = map[string]interface{}{"status": 400, "msg": models.ErrUserAlreadyExists.Error()}
		this.ServeJSON()
		return
	}
	user.UserName = this.Input().Get(models.DATA_KEY_USER_NAME)
	beego.Info(user.UserName)
	user.NickName = this.Input().Get(models.DATA_KEY_NICK_NAME)
	beego.Info(user.NickName)
	user.Location = this.Input().Get(models.DATA_KEY_LOCATION)
	beego.Info(user.Location)
	user.Password = this.Input().Get(models.DATA_KEY_PASSWORD)
	beego.Info(user.Password)

	//SetRegisterChan(user)

	if errs := models.CheckValid(user); len(errs) > 0 {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": errs[0].Error()}
		this.ServeJSON()
		return
	}

	err := models.AddUser(user)
	if err != nil {
		//this.Ctx.WriteString(err.Error())
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 0, "msg": fmt.Sprintf("Create Account succeed,mobile is %s", user.Mobile)}
	this.ServeJSON()

	err = models.SaveUsers(models.Accounts)
	if err != nil {
		fmt.Printf("save users error:%s\n", err.Error())
	}
	return
}

func (this *AccountController) sendSMS() {
	mobile := this.GetString(models.DATA_KEY_MOBILE + "-from")
	beego.Info("aaa:" + mobile)
	if b, err := new(models.MobileValidator).Validate(mobile); !b {
		//this.Ctx.WriteString("mobile:" + err.Error())
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "mobile-from:" + err.Error()}
		this.ServeJSON()
		return
	}

	user, err := models.GetUserByMobile(mobile)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error()}
		this.ServeJSON()
		//this.Ctx.WriteString(err.Error())
		return
	}

	if !user.IsLogin {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": models.ErrNotLogin.Error()}
		this.ServeJSON()
		//this.Ctx.WriteString(models.ErrNotLogin.Error())
		return
	}

	destation := this.GetString(models.DATA_KEY_MOBILE + "-to")
	beego.Info("bbb:" + destation)
	if b, err := new(models.MobileValidator).Validate(destation); !b {
		//this.Ctx.WriteString("destationMobile:" + err.Error())
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "destatmoionMobile:" + err.Error()}
		this.ServeJSON()
		return
	}

	/*if _, err := models.GetUserByMobile(destation); err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}*/

	content := this.GetString("content")
	beego.Info("ccc:" + content)
	if len(content) == 0 {
		//this.Ctx.WriteString("please enter sms message")
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "please enter sms message"}
		this.ServeJSON()
		return
	}
	beego.Info("ddd 11")
	models.SaveSendMsg(mobile, destation, content)
	beego.Info("ddd")

	this.Data["json"] = map[string]interface{}{"status": 0, "msg": "send succeed"}
	this.ServeJSON()

	err = models.SaveSMS(models.MsgManage)
	if err != nil {
		fmt.Printf("save users error:%s\n", err.Error())
	}

	return
	/*
		case "receive":
			msg := models.GetSefMsgs(mobile)
			if msg.ReceiveMsgs != nil { //map[string][]string
				this.Ctx.WriteString("received msg is here")
			} else {
				this.Ctx.WriteString("not exist received msg")
			}

			if msg.SendMsgs != nil {
				this.Ctx.WriteString("sent msg is here")
			} else {
				this.Ctx.WriteString("not exist send msg")
			}*/

}

func (this *AccountController) deleteUser() {

	if strings.Compare(models.SYS_ACCOUNT, this.GetString("sysAccount")) != 0 {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "err account"}
		this.ServeJSON()
		return
	}

	if strings.Compare(models.SYS_PWD, this.GetString("sysPassword")) != 0 {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "err password"}
		this.ServeJSON()
		return
	}

	mobile := this.GetString("mobile")
	if b, err := new(models.MobileValidator).Validate(mobile); !b {
		//this.Ctx.WriteString("destationMobile:" + err.Error())
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": fmt.Sprintf("mobile:%s;%s", mobile, err.Error())}
		this.ServeJSON()
		return
	}

	models.RemoveUserByMobile(mobile)

	this.Data["json"] = map[string]interface{}{"status": 0, "msg": "delete succeed"}
	this.ServeJSON()

	err := models.SaveUsers(models.Accounts)
	if err != nil {
		fmt.Printf("save users error:%s\n", err.Error())
	}

	return
}

func (this *AccountController) unknownOption() {
	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "Unknown option"}
	this.ServeJSON()
	//SetResultChan("Unknown option")
	return
}

func accountProcess() {
	for {
		select {
		case user := <-RegisterChan:
			beego.Info("sssssssss")
			if models.IsExistUser(user.Mobile) {
				//this.Ctx.WriteString(models.ErrUserAlreadyExists.Error())
				SetResultChan(models.ErrUserAlreadyExists.Error())
			} else {
				if errs := models.CheckValid(user); len(errs) > 0 {
					SetResultChan(errs[0].Error())
				} else {
					err := models.AddUser(user)
					if err != nil {
						//this.Ctx.WriteString(err.Error())
						SetResultChan(err.Error())
					} else {
						SetResultChan("Register succeed")
					}
				}
			}

		case loginInfo := <-LoginChan:
			user, err := models.GetUserByMobile(loginInfo.Mobile)
			if err != nil {
				//this.Ctx.WriteString(err.Error())
				SetResultChan(err.Error())
			} else {
				if strings.Compare(user.Password, loginInfo.Password) != 0 {
					SetResultChan(models.ErrPwd.Error())
				} else {
					SetResultChan("Login succeed")
				}
			}
		}
	}
}

func init() {
	//go accountProcess()
}
