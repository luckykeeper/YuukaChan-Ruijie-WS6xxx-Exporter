// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 登出
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/8/28 10:52
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package subFunction

import (
	"crypto/tls"
	"log"
	"net/http"
)

// 登出
func LogOut(SIDS, DeviceIP string) {
	for {
		logoutUrl := "https://" + DeviceIP + "/logout.do"

		cookie := &http.Cookie{Name: "SIDS", Value: SIDS, Path: "/"}

		//构建https请求，忽略证书错误
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("GET", logoutUrl, nil)
		req.AddCookie(cookie)

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			continue
		} else {
			log.Println("已经登出")
			return
		}
	}
}
