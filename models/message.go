package models

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const FILE_SMS_MSGS = "smsMsg.dat"

type Message map[string][]string

type UserMessages struct {
	SendMsgs    Message
	ReceiveMsgs Message
}

type Manage map[string]UserMessages

var MsgManage = make(Manage)

var msgLock = new(sync.RWMutex)

func SaveSendMsg(sefMobile, destationMobile, content string) {
	//var msg Message
	msgLock.Lock()
	defer msgLock.Unlock()
	if MsgManage[sefMobile].SendMsgs == nil {
		MsgManage[sefMobile] = UserMessages{SendMsgs: make(Message)}
	}

	if MsgManage[sefMobile].SendMsgs[destationMobile] == nil {
		MsgManage[sefMobile].SendMsgs[destationMobile] = make([]string, 1)
	}

	MsgManage[sefMobile].SendMsgs[destationMobile] = append(MsgManage[sefMobile].SendMsgs[destationMobile], content)
}

func SaveReceivedMsg(sefMobile, destationMobile, content string) {
	msgLock.Lock()
	defer msgLock.Unlock()
	if MsgManage[sefMobile].ReceiveMsgs == nil {
		MsgManage[sefMobile] = UserMessages{ReceiveMsgs: make(Message)}
	}

	if MsgManage[sefMobile].ReceiveMsgs[destationMobile] == nil {
		MsgManage[sefMobile].ReceiveMsgs[destationMobile] = make([]string, 1)
	}

	MsgManage[sefMobile].ReceiveMsgs[destationMobile] = append(MsgManage[sefMobile].ReceiveMsgs[destationMobile], content)
}

func GetSefMsgs(sefMobile string) UserMessages {
	return MsgManage[sefMobile]
}

func SaveSMS(m Manage) error {
	msgLock.Lock()
	defer msgLock.Unlock()
	file, err := os.OpenFile(FILE_SMS_MSGS, os.O_CREATE|os.O_WRONLY, 0755)
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

func ReadSMS(m *Manage) error {
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
