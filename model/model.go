// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 数据模型
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/8/28 10:52
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package model

// 获取AP列表-返回json
type ApList struct {
	OsTime float64     `json:"osTime"`
	Code   float64     `json:"code"`
	Msg    string      `json:"msg"`
	Data   *ApListData `json:"data"`
}

type ApListData struct {
	TotalCount float64        `json:"totalCount"`
	List       []ApListDetail `json:"list"`
}

type ApListDetail struct {
	OnlineTime         float64 `json:"onlineTime"`
	Ip                 string  `json:"ip"`
	Downflow_kbps      float64 `json:"downflow_kbps"`
	StaLimit           float64 `json:"staLimit"`
	SoftwareVersion    string  `json:"softwareVersion"`
	Flow_kbps          float64 `json:"flow_kbps"`
	VtapStatus         float64 `json:"vtapStatus"`
	Cpu_percent        float64 `json:"cpuPercent"`
	VacId              float64 `json:"vacId"`
	VapSupport         string  `json:"vapsupport"`
	ApGroup            string  `json:"apgroup"`
	HardwareVersion    string  `json:"hardwareVersion"`
	Location           string  `json:"location"`
	Model              string  `json:"model"`
	MasterApName       string  `json:"masterApName"`
	SubApMac           string  `json:"subApMac"`
	OfflineCount       float64 `json:"offlineCount"`
	SubApName          string  `json:"subApName"`
	Freememory_percent float64 `json:"freememory_percent"`
	MasterGroup        string  `json:"mastergroup"`
	Upflow_kbps        float64 `json:"upflow_kbps"`
	Stanum             float64 `json:"stanum"`
	Ctxid              float64 `json:"ctxid"`
	WorkMode           string  `json:"workMode"`
	ApName             string  `json:"apname"`
	Mac                string  `json:"mac"`
	State              string  `json:"state"`
	Acdes              string  `json:"acdes"`
}

// 获取在线用户-返回json
type UserList struct {
	OsTime float64       `json:"osTime"`
	Code   float64       `json:"code"`
	Msg    string        `json:"msg"`
	Data   *UserListData `json:"data"`
}

type UserListData struct {
	TotalCount float64          `json:"totalCount"`
	List       []UserListDetail `json:"list"`
}

type UserListDetail struct {
	Band          string  `json:"band"`
	Ipv4          string  `json:"ipv4"`
	Ipv4upRate    float64 `json:"ipv4upRate"`
	NetAuth       string  `json:"netAuth"`
	Rate          float64 `json:"rate"`
	Property80211 float64 `json:"802.11"`
	FlowNum       float64 `json:"flowNum"`
	Mac           string  `json:"mac"`
	Ipv6upRate    float64 `json:"ipv6upRate"`
	AcDec         string  `json:"acDec"`
	ApName        string  `json:"apname"`
	ApRadio       float64 `json:"apRadio"`
	Delay         float64 `json:"delay"`
	ApIp          string  `json:"apIp"`
	OnlineTimeval float64 `json:"onlineTimeval"`
	AssocAuth     string  `json:"assocAuth"`
	VacId         float64 `json:"vacId"`
	Vlan          int     `json:"vlan"`
	Ipv6          string  `json:"ipv6"`
	Ipv4downRate  float64 `json:"ipv4downRate"`
	PacketLoss    float64 `json:"packetLoss"`
	ClientType    string  `json:"clientType"`
	Ssid          string  `json:"ssid"`
	RSSI          float64 `json:"RSSI"`
	Ipv6downRate  float64 `json:"ipv6downRate"`
	// 新版 RGOS 新增的属性（20230828）
	AuthUsername string `json:"authUsername"`
}

// 获取盐值接口 | 20230906添加
type HmacInfo struct {
	Salt   string `json:"salt"`
	Iter   int    `json:"iter"`
	Digest string `json:"digest"`
	Keylen int    `json:"keylen"`
}
