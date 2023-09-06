// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 程序入口
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/9/06 13:20
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	subFunction "yuukaChan-Ruijie-WS6xxx-Exporter/subfunction"

	"github.com/VictoriaMetrics/metrics"
	"github.com/urfave/cli/v2"
	"gopkg.in/ini.v1"
)

var (
	DeviceIP      string
	DeviceUser    string
	DeviceAuth    string
	ApListStart   string
	ApListEnd     string
	UserListStart string
	UserListEnd   string
	SIDS          string
)

// 基础变量设置
var (
	YuukaExporterWebPort, YuukaExporterBasicUrl, YuukaExporterBasicUsername, YuukaExporterBasicPassword string
)

// CLI
func YuukaExporterCLI() {
	YuukaExporter := &cli.App{
		Name: "YuukaExporter",
		Usage: "YuukaExporter - 锐捷RG-WS6xxx Series Exporter" +
			"\nPowered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site>" +
			"\n————————————————————————————————————————" +
			"\n注意：使用前需要先填写同目录下 config.ini !",
		Version: "1.0.1_build20230906",
		Commands: []*cli.Command{
			// 爬取数据，启动 Exporter
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "启动 YuukaExporter",
				Action: func(cCtx *cli.Context) error {
					yuukaExporter()
					return nil
				},
			},
		},
		Copyright: "Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site> | https://github.com/luckykeeper",
	}

	if err := YuukaExporter.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// 程序入口
func main() {
	readConfig()
	YuukaExporterCLI()
}

// 程序入口
func yuukaExporter() {
	log.Println("YuukaChan RG-WS6xxx Series Exporter Started!")
	log.Println("Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site>")
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		log.Println("________________________________")
		log.Println("收到采集请求，启动数据采集业务流")
		log.Println("数据采集 - 登录 AC - 开始")
		SIDS := subFunction.GetSIDS(DeviceIP, DeviceUser, DeviceAuth)

		log.Println("数据采集 - 获取 AP 列表 - 开始")
		apListData := subFunction.GetApList(DeviceIP, SIDS, ApListStart, ApListEnd)
		log.Println("指标发布 - AP 列表 - 开始")
		subFunction.UpdateMetric("AC_RGOS_SysTime",
			map[string]string{"AC_IP": DeviceIP},
			apListData.OsTime-28800)
		subFunction.UpdateMetric("AC_Managed_APCount",
			map[string]string{"AC_IP": DeviceIP},
			apListData.Data.TotalCount)

		// 自定义指标 - AP 总下行流量 / AP 总上行流量（单位均为 kbps ， 换算 m/s 除以 86400）
		var (
			apTotalDownKbps, apTotalUpKbps float64
		)
		apTotalDownKbps = 0
		apTotalUpKbps = 0
		for _, apList := range apListData.Data.List {
			apTotalDownKbps = apTotalDownKbps + apList.Downflow_kbps
			apTotalUpKbps = apTotalUpKbps + apList.Upflow_kbps
			subFunction.UpdateMetric("AP_OnlineTime",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.OnlineTime)

			subFunction.UpdateMetric("AP_Downflow_kbps",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Downflow_kbps)

			subFunction.UpdateMetric("AP_StaLimit",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.StaLimit)

			subFunction.UpdateMetric("AP_Flow_kbps",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Flow_kbps)

			subFunction.UpdateMetric("AP_VtapStatus",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.VtapStatus)

			subFunction.UpdateMetric("AP_Cpu_percent",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Cpu_percent)

			subFunction.UpdateMetric("AP_VacId",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.VacId)

			subFunction.UpdateMetric("AP_OfflineCount",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.OfflineCount)

			subFunction.UpdateMetric("AP_Freememory_percent",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Freememory_percent)

			subFunction.UpdateMetric("AP_Upflow_kbps",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Upflow_kbps)

			subFunction.UpdateMetric("AP_Stanum",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Stanum)

			subFunction.UpdateMetric("AP_Ctxid",
				map[string]string{"AC_IP": DeviceIP,
					"AP_IP":              apList.Ip,
					"AP_SoftwareVersion": apList.SoftwareVersion,
					"AP_VapSupport":      apList.VapSupport,
					"AP_ApGroup":         apList.ApGroup,
					"AP_HardwareVersion": apList.HardwareVersion,
					"AP_Location":        apList.Location,
					"AP_Model":           apList.Model,
					"AP_MasterApName":    apList.MasterApName,
					"AP_SubApMac":        apList.SubApMac,
					"AP_SubApName":       apList.SubApName,
					"AP_Mastergroup":     apList.MasterGroup,
					"AP_WorkMode":        apList.WorkMode,
					"AP_ApName":          apList.ApName,
					"AP_Mac":             apList.Mac,
					"AP_State":           apList.State,
					"AP_Acdes":           apList.Acdes},
				apList.Ctxid)
		}
		// 自定义指标 - AP 总下行流量 / AP 总上行流量（单位均为 kbps ， 换算 m/s 除以 86400）
		subFunction.UpdateMetric("AC_Managed_APTotalDownKbps",
			map[string]string{"AC_IP": DeviceIP},
			apTotalDownKbps)

		subFunction.UpdateMetric("AC_Managed_APTotalUpKbps",
			map[string]string{"AC_IP": DeviceIP},
			apTotalUpKbps)
		// AP 总流量
		subFunction.UpdateMetric("AC_Managed_APTotalKbps",
			map[string]string{"AC_IP": DeviceIP},
			apTotalUpKbps+apTotalDownKbps)

		log.Println("指标发布 - AP 列表 - 完成")

		log.Println("数据采集 - 获取在线用户信息 - 开始")
		userListData := subFunction.GetOnileUserList(DeviceIP, SIDS, UserListStart, UserListEnd)
		log.Println("指标发布 - 在线用户信息 - 开始")
		subFunction.UpdateMetric("AC_Managed_OnlineUserCount",
			map[string]string{"AC_IP": DeviceIP},
			userListData.Data.TotalCount)

		// 自定义指标 - 2.4G 用户总数 / 5G 用户总数
		var (
			User2GTotal, User5GTotal float64
		)
		User2GTotal = 0
		User5GTotal = 0
		for _, userList := range userListData.Data.List {
			if userList.Band == "2G" {
				User2GTotal++
			} else if userList.Band == "5G" {
				User5GTotal++
			}

			subFunction.UpdateMetric("User_Ipv4upRate",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Ipv4upRate)

			subFunction.UpdateMetric("User_Rate",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Rate)

			subFunction.UpdateMetric("User_Property80211",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Property80211)

			subFunction.UpdateMetric("User_FlowNum",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.FlowNum)

			subFunction.UpdateMetric("User_Ipv6upRate",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Ipv6upRate)

			subFunction.UpdateMetric("User_ApRadio",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.ApRadio)

			subFunction.UpdateMetric("User_Delay",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Delay)

			subFunction.UpdateMetric("User_OnlineTimeval",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.OnlineTimeval)

			subFunction.UpdateMetric("User_VacId",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.VacId)

			subFunction.UpdateMetric("User_Ipv4downRate",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Ipv4downRate)

			subFunction.UpdateMetric("User_PacketLoss",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.PacketLoss)

			subFunction.UpdateMetric("User_RSSI",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.RSSI)

			subFunction.UpdateMetric("User_Ipv6downRate",
				map[string]string{"AC_IP": DeviceIP,
					"User_Band":         userList.Band,
					"User_IPV4":         userList.Ipv4,
					"User_NetAuth":      userList.NetAuth,
					"User_Mac":          userList.Mac,
					"User_AcDec":        userList.AcDec,
					"User_ApName":       userList.ApName,
					"User_ApIp":         userList.ApIp,
					"User_AssocAuth":    userList.AssocAuth,
					"User_Vlan":         strconv.Itoa(userList.Vlan),
					"User_Ipv6":         userList.Ipv6,
					"User_ClientType":   userList.ClientType,
					"User_Ssid":         userList.Ssid,
					"User_AuthUsername": userList.AuthUsername},
				userList.Ipv6downRate)
		}
		// 自定义指标 - 2.4G 用户总数 / 5G 用户总数
		subFunction.UpdateMetric("AC_Managed_User2GTotal",
			map[string]string{"AC_IP": DeviceIP},
			User2GTotal)

		subFunction.UpdateMetric("AC_Managed_User5GTotal",
			map[string]string{"AC_IP": DeviceIP},
			User5GTotal)

		log.Println("指标发布 - 在线用户信息 - 完成")

		// 登录完成后必须退出，否则多次查询后会 Sessionid num reach max ，不再允许登录
		// 关于加密部分，可以确认加密的盐值一定，加密接口 /hmac_info.do?user=<username>&_=1693211894520
		// 登出接口示例 GET /logout.do?_=1693213958891 参数是当前时间戳
		subFunction.LogOut(SIDS, DeviceIP)

		// 发布示例
		// subfunction.UpdateMetric("test",
		// 	map[string]string{"testIP": "127.0.0.1"},
		// 	1)
		metrics.WritePrometheus(w, true)
	})
	log.Fatal(http.ListenAndServe(YuukaExporterWebPort, nil))

}

// read config.ini
func readConfig() {
	configFile, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("读取配置文件 config.ini 失败，失败原因为: %v", err)
		os.Exit(1)
	}

	YuukaExporterWebPort = configFile.Section("YuukaExporter").Key("webPort").String()
	YuukaExporterBasicUrl = configFile.Section("YuukaExporter").Key("basicUrl").String()
	YuukaExporterBasicUsername = configFile.Section("YuukaExporter").Key("basicUser").String()
	YuukaExporterBasicPassword = configFile.Section("YuukaExporter").Key("basicPassword").String()

	DeviceIP = configFile.Section("login").Key("ip").String()
	DeviceUser = configFile.Section("login").Key("user").String()
	DeviceAuth = configFile.Section("login").Key("auth").String()
	ApListStart = configFile.Section("apList").Key("start").String()
	ApListEnd = configFile.Section("apList").Key("end").String()
	UserListStart = configFile.Section("userList").Key("start").String()
	UserListEnd = configFile.Section("userList").Key("end").String()
}
