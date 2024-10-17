package weferry

type RequestType struct {
	Receiver string `json:"receiver"`
	Msg      string `json:"msg"`
	Aters    string `json:"aters"`
	Path     string `json:"path"`
	Db       string `json:"db"`
	Sql      string `json:"sql"`
	Extra    string `json:"extra"`
	Dir      string `json:"dir"`
	ID       int64  `json:"id"`
	Timetout int64  `json:"timeout"`
	Roomid   string `json:"roomid"`
	Wxid     string `json:"wxid"`
}

type ResponseType struct {
	Status int    `json:"status"`
	Error  any    `json:"error"`
	Data   []Data `json:"data"`
}
type Data struct {
	UserName        string `json:"UserName"`
	NickName        string `json:"NickName"`
	UserNameList    string `json:"UserNameList"`
	Reserved2       string `json:"Reserved2"`
	DisplayNameList string `json:"DisplayNameList"`
	Reserved1       any    `json:"Reserved1"`
	Reserved5       any    `json:"Reserved5"`
	Reserved7       any    `json:"Reserved7"`
	Reserved4       any    `json:"Reserved4"`
	RoomData        string `json:"RoomData"`
	Reserved3       any    `json:"Reserved3"`
	SelfDisplayName string `json:"SelfDisplayName"`
	Reserved8       any    `json:"Reserved8"`
	Owner           any    `json:"Owner"`
	IsShowName      int    `json:"IsShowName"`
	ChatRoomFlag    int    `json:"ChatRoomFlag"`
	Reserved6       any    `json:"Reserved6"`
	ChatRoomName    string `json:"ChatRoomName"`
}
