package weferry

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// 执行sql语句
func PlanSQL(sql string) (*http.Response, error) {

	apiUrl := "http://106.55.251.45:10010/sql"
	request := RequestType{
		Db:  "MicroMsg.db",
		Sql: sql,
	}
	requestdata, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := sendHTTPRequest("POST", apiUrl, requestdata)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取群人数
func (f *Framework) GetChatRoomNumber(roomid string) int {
	sql := fmt.Sprintf("SELECT * FROM ChatRoom WHERE ChatRoomName='%s';", roomid)
	resp, err := PlanSQL(sql)
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()
	responsebody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var respdata ResponseType

	if err := json.Unmarshal(responsebody, &respdata); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}
	// 打印结果
	UserNameList := strings.Split(respdata.Data[0].UserNameList, "^G")
	// 根据群成员id查找群昵称
	// 打印列表长度

	return len(UserNameList)
}

// 全局函数：查询群成员昵称
var framework = &Framework{}

func GetChatRoomNick(userNameId string) string {
	return framework.GetChatRoomNick(userNameId)
}

// 查询群成员昵称
func (f *Framework) GetChatRoomNick(userNameId string) string {
	sql := fmt.Sprintf("SELECT UserName, NickName FROM Contact where UserName='%s';", userNameId)
	resp, err := PlanSQL(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responsebody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var respdata ResponseType

	if err := json.Unmarshal(responsebody, &respdata); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	return respdata.Data[0].NickName
}
