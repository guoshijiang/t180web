package types

//获取团队个人信息结构体
type TeamPersonInfo struct {
	PersonAddr string `json:"addr"`
	PersonName string `json:"name"`
	PersonImg string `json:"img"`
	PersonContent string `json:"content"`
	PersonGy string `json:"gy"`
	AddTime int `json:"addtime"`
}