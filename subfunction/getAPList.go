// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 获取 AP 列表
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/8/28 10:52
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package subFunction

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"yuukaChan-Ruijie-WS6xxx-Exporter/model"
)

// 获取AP列表
func GetApList(DeviceIP, SIDS, ApListStart, ApListEnd string) (APList model.ApList) {
	for {
		log.Println("开始获取AP列表！")
		apListUrl := "https://" + DeviceIP + "/web/init.cgi/ac.dashboard.ap_list/getApList"

		apListQuery := url.Values{}
		apListQuery.Add("Start", ApListStart)
		apListQuery.Add("End", ApListEnd)

		cookie := &http.Cookie{Name: "SIDS", Value: SIDS, Path: "/"}

		//构建https请求，忽略证书错误
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("POST", apListUrl, strings.NewReader(apListQuery.Encode()))

		if err != nil {
			log.Println(err)
			continue
		}
		//header
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.AddCookie(cookie)

		//发送请求
		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		defer res.Body.Close()

		//返回结果
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		// fmt.Println(string(body))
		var ApListDataReturn model.ApList
		json.Unmarshal(body, &ApListDataReturn)

		log.Println("数据采集 - 获取 AP 列表 - 成功完成")
		return ApListDataReturn
	}
}
