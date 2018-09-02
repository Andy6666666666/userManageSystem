package models

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

const FILE_USERS = "userInfo.dat"

type UserInfo struct {
	Mobile   string `validate:"mobile,min=11,max=11"`
	UserName string `validate:"nameOrPwd"`
	NickName string `validate:"nameOrPwd"`
	Location string `validate:"string,min=10,max=64"`
	Password string `validate:"nameOrPwd"`
	IsLogin  bool   `validate:""`
}

const (
	DATA_KEY_MOBILE    = "mobile"
	DATA_KEY_USER_NAME = "username"
	DATA_KEY_NICK_NAME = "nickname"
	DATA_KEY_LOCATION  = "location"
	DATA_KEY_PASSWORD  = "password"
	SYS_ACCOUNT        = "123456"
	SYS_PWD            = "123456"
)

var (
	ErrUserAlreadyExists = errors.New("The user you were trying to create already exists")
	ErrUserNotExists     = errors.New("The user not exist")
	ErrMobile            = errors.New("Invalid Mobile Number")
	ErrPwd               = errors.New("Error Password")
	ErrInvalidUserInfo   = errors.New("Invalid user info.")
	ErrNotLogin          = errors.New("please login")
)

type UserType map[string]UserInfo

var Accounts = make(UserType)
var lock = new(sync.RWMutex)

func IsExistUser(mobile string) bool {
	if strings.Count(mobile, "")-1 != 11 {
		return false
	}
	user := Accounts[mobile]
	if strings.Compare(user.Mobile, mobile) != 0 {
		return false
	}

	return true
}

func GetUserByMobile(mobile string) (UserInfo, error) {
	if !IsExistUser(mobile) {
		return UserInfo{Mobile: ""}, ErrUserNotExists
	}
	return Accounts[mobile], nil
}

func AddUser(user UserInfo) error {
	lock.Lock()
	defer lock.Unlock()
	if IsExistUser(user.Mobile) {
		return ErrUserAlreadyExists
	}

	Accounts[user.Mobile] = user

	return nil
}

func RemoveUserByMobile(mobile string) {
	lock.Lock()
	defer lock.Unlock()
	Accounts[mobile] = UserInfo{Mobile: ""}
}

func UpdateUser(user UserInfo) {
	lock.Lock()
	defer lock.Unlock()
	Accounts[user.Mobile] = user
}

func CheckValid(user UserInfo) []error {
	return validateStruct(user)
}

func SaveUsers(m UserType) error {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.OpenFile(FILE_USERS, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	fileWrite := bufio.NewWriter(file)

	_, err = fileWrite.WriteString(fmt.Sprintf("%v", m))
	if err != nil {
		return err
	}

	err = fileWrite.Flush()
	if err != nil {
		return err
	}

	return nil
}

func ReadUsers(m *UserType) error {
	file, err := os.OpenFile(FILE_SMS_MSGS, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Sscanf(string(buf), "%v", m)
	fmt.Printf("%v\n", m)
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
