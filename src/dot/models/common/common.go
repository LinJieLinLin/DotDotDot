package common

const (
	E0 = "传入数据有误！"
	E1 = "数据库连接失败！"
)

type Re struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
