/*
@Time : 2019/7/8 11:19 
@Author : Lukebryan
@File : wechat_api.go
@Software: GoLand
*/
package api

import (
	"activationcode/config"
	"activationcode/util"
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"github.com/spf13/cast"
	"log"
	"strings"
)

func GetLoginMap(wxidArr []string,userID string) map[string]int {
	dataMap := make(map[string]interface{})
	dataMap["wxid_list"] = strings.Join(wxidArr,",")
	heardMap := make(map[string]string)
	heardMap["Authorization"] = utils.GentToken(userID,"")
	data, err := utils.PostFormRequest(config.Sysconfig.WechatServerAddr+"/account/batchcheckloginstatus",dataMap,heardMap)

	if err != nil {
		log.Println("BindWechatGroup List Error1: ",err)
	}

	maps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data),&maps)
	if err != nil {
		log.Println("BindWechatGroup List Error2: ",err)
	}

	dataMaps := make(map[string]int)
	b,_ := json.Marshal(maps["Data"])
	err = json.Unmarshal(b,&dataMaps)
	if err != nil {
		log.Println("BindWechatGroup List Error3: ",err)
	}
	return dataMaps
}

func LoginApi() (string,error) {
	dataMap := make(map[string]interface{})
	dataMap["user_name"] = config.Sysconfig.APIUserName
	dataMap["password"] = config.Sysconfig.APIPassword
	heardMap := make(map[string]string)
	resp, err := utils.PostFormRequest(config.Sysconfig.WechatServerAddr+"/user/login", dataMap, heardMap)
	if err != nil {
		log.Println("api user login error: ", err)
		return "",err
	}
	maps := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &maps)
	if err != nil {
		log.Println("api user login json.Unmarshal Error: ", err)
		return "",err
	}
	if cast.ToString(maps["Code"]) == "-1" {
		log.Println(cast.ToString(maps["Msg"]))
		return "",errors.New(cast.ToString(maps["Msg"]))
	}
	return cast.ToString(maps["Data"]),nil
}

func SyncContacts(userID,wechatID string,currentContactSeq,currentChatroomContactSeq int) {

	dataMap := make(map[string]interface{})
	dataMap["currentContactSeq"] = currentContactSeq
	dataMap["currentChatroomContactSeq"] = currentChatroomContactSeq
	heardMap := make(map[string]string)
	heardMap["Authorization"] = utils.GentToken(userID,wechatID)
	resp, err := utils.PostFormRequest(config.Sysconfig.WechatServerAddr+"/account/synccontacts",dataMap,heardMap)
	if err != nil {
		log.Println("synccontacts error: ",err)
		return
	}
	maps := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &maps)
	if err != nil {
		log.Println("synccontacts json.Unmarshal Error: ", err)
		return
	}

	b, err := json.Marshal(maps["Data"])
	if err != nil {
		log.Println("synccontacts json.Marshal Error: ", err)
		return
	}
	//{"Code":0,"Data":{"LoopFlag":false,"FailedContacts":null,"CurrentContactSeq":0,"CurrentChatroomContactSeq":0},"Msg":"同步成功"}
	finalMap := make(map[string]interface{})
	err = json.Unmarshal(b, &finalMap)
	if err != nil {
		log.Println("synccontacts json.Unmarshal Error: ", err)
		return
	}
	if !cast.ToBool(finalMap["LoopFlag"]) {

		return
	}else {
		SyncContacts(userID,wechatID,cast.ToInt(finalMap["CurrentContactSeq"]),cast.ToInt(finalMap["CurrentChatroomContactSeq"]))
	}
}

//发消息
func SendMessage(fromWxId, toWxID, userID, content string) bool {

	dataMap := make(map[string]interface{})
	dataMap["to_wxid"] = toWxID
	dataMap["content"] = content
	heardMap := make(map[string]string)
	heardMap["Authorization"] = utils.GentToken(userID, fromWxId)
	data, err := utils.PostFormRequest(config.Sysconfig.WechatServerAddr+"/message/sendmessage", dataMap, heardMap)
	if err != nil {
		log.Println("sendmessage error: ", err)
		return false
	}
	maps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &maps)
	if err != nil {
		log.Println("sendmessage json.Unmarshal Error: ", err)
		return false
	}

	if cast.ToString(maps["Code"]) == "-1" {
		return false
	}

	return true
}

//添加好友到标签
func AddFriendsWithLabel(userID,wechatID,label string) bool {

	dataMap := make(map[string]interface{})
	dataMap["label"] = label
	dataMap["wxid"] = wechatID
	heardMap := make(map[string]string)
	heardMap["Authorization"] = utils.GentToken(userID, wechatID)
	data, err := utils.PostFormRequest(config.Sysconfig.WechatServerAddr+"/account/modifycontactlabellist", dataMap, heardMap)
	if err != nil {
		log.Println("modifycontactlabellist error: ", err)
		return false
	}
	maps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &maps)
	if err != nil {
		log.Println("modifycontactlabellist json.Unmarshal Error: ", err)
		return false
	}

	if cast.ToString(maps["Code"]) == "-1" {
		return false
	}

	return true
}

//获取标签列表好友
func GetFriendsWithLabel(userID,wechatID,label string) bool {

	dataMap := make(map[string]interface{})
	dataMap["label"] = label
	heardMap := make(map[string]string)
	heardMap["Authorization"] = utils.GentToken(userID, wechatID)
	data, err := utils.PostFormRequest(config.Sysconfig.WechatServerAddr+"/account/getcontactlabellist", dataMap, heardMap)
	if err != nil {
		log.Println("getcontactlabellist error: ", err)
		return false
	}
	maps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &maps)
	if err != nil {
		log.Println("getcontactlabellist json.Unmarshal Error: ", err)
		return false
	}

	if cast.ToString(maps["Code"]) == "-1" {
		return false
	}

	return true
}