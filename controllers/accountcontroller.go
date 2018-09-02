package controllers

import (
	"time"
	"userManageSystem/models"

	"github.com/astaxie/beego"
)

var (
	RegisterChan = make(chan models.UserInfo, 10)
	LoginChan    = make(chan models.UserInfo, 10)
	ResultChan   = make(chan string, 10)
)

func SetRegisterChan(user models.UserInfo) {
	RegisterChan <- user
}

func SetLoginChan(info models.UserInfo) {
	LoginChan <- info
}

func SetResultChan(result string) {
	beego.Info(result)
	ResultChan <- result
}

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Get() {
	this.TplName = "client.html"
}

func (this *AccountController) Post() {
	option := this.Input().Get("options")

	beego.Info(option)

	if option == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "option is null", "time": time.Now().Format("2006-12-12 12:12:12")}
		this.ServeJSON()
		return
	}

	switch option {
	case "login":
		this.login()
	case "register":
		this.register()
	case "sendSMS":
		this.sendSMS()
	case "deleteUser":
		this.deleteUser()
	default:
		this.unknownOption()
	}

	/*
		go func(this *AccountController) {
			for {
				select {
				case result := <-ResultChan:
					this.Data["json"] = map[string]interface{}{"status": 400, "msg": result, "time": time.Now().Format("2006-12-12 12:12:12")}
					this.ServeJSON()
				}
			}
		}(this)*/
}

/*
	var user models.UserInfo

	user.Mobile = c.GetString(models.DATA_KEY_MOBILE)
	if models.IsExistUser(user.Mobile) {
		c.Ctx.WriteString(models.ErrUserAlreadyExists.Error())
		return
	}
	user.UserName = c.GetString(models.DATA_KEY_USER_NAME)
	user.NickName = c.GetString(models.DATA_KEY_NICK_NAME)
	user.Location = c.GetString(models.DATA_KEY_LOCATION)
	user.Password = c.GetString(models.DATA_KEY_PASSWORD)

	err := models.AddUser(user)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("Create Account succeed,mobile is %s", user.Mobile))

*/
