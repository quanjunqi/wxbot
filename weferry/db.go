package weferry

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"wxbot/engine/pkg/redisutil"
)

// 执行sql语句
func PlanSQL(sql string) (*http.Response, error) {

	apiUrl := "http://124.220.212.132:10010/sql"
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

// 下载图片
func SaveImage(id int64, dir, extra string) {
	apiUrl := "http://124.220.212.132:10010/save-image"
	request := RequestType{
		Dir:      dir,
		Extra:    extra,
		ID:       id,
		Timetout: 10,
	}
	requestdata, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := sendHTTPRequest("POST", apiUrl, requestdata)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 图片下载成功，将 ID 和图片地址存入 Redis

	// 正则表达式，匹配文件名（去掉扩展名）
	re := regexp.MustCompile(`([^/\\]+)\.dat$`)
	// 获取图片文件的名字
	match := re.FindStringSubmatch(extra)
	if len(match) > 1 {
		// match[1] 是捕获组，包含文件名
		imageURL := fmt.Sprintf("https://wechat-qjq.oss-cn-shenzhen.aliyuncs.com/%s.jpg", match[1])
		// 使用 Redis 工具包
		redisClient := redisutil.GetInstance("localhost:6379", "", 0)
		// 正确地将 int64 转换为字符串
		idStr := strconv.FormatInt(id, 10)
		err = redisClient.Set(idStr, imageURL)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Image ID %d saved with URL %s in Redis", id, imageURL)
	} else {
		fmt.Println("No match found.")
	}
}
