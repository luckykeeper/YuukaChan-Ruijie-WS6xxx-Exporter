// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 登录并获取SIDS
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/8/28 10:52
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package subFunction

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 登录并获取SIDS
func GetSIDS(DeviceIP string, DeviceUser, DeviceAuth string) (SIDS string) {
	for {
		log.Println("开始登录流程！")
		// 定义登录相关参数
		loginUrl := "https://" + DeviceIP + "/login.do"

		// 下面注释掉的是旧版的 authData ，说明请参考 config.ini
		// authData := `auth=` + DeviceAuth
		authData := "user=" + DeviceUser + "&key=" + DeviceAuth
		authDataByte := []byte(authData) // 需要注意的是，因为锐捷AC需要的不是标准形式的 JSON 字符串，
		// 所以不能用json序列化之后转[]byte而是直接转[]byte
		// fmt.Println(authData)

		loginParam := bytes.NewBuffer(authDataByte)
		// fmt.Println(loginParam)

		//构建https请求，忽略证书错误
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("POST", loginUrl, loginParam)

		if err != nil {
			log.Println(err)
			continue
		}
		//header
		req.Header.Add("Content-Type", "application/json")

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
		cookie := res.Header.Get("Set-Cookie")
		// fmt.Println(string(body))
		bodyResult := strings.Split((strings.Split(string(body), "<return-code>")[1]), "</return-code>")[0]
		if bodyResultInt, _ := strconv.Atoi(bodyResult); bodyResultInt == 0 {
			// log.Println("登录成功！")
			SIDS = (strings.Split(strings.Split(cookie, "SIDS=")[1], "; SameSite"))[0]
			// log.Println("SIDS=", SIDS)
			// log.Println("登录流程顺利完成！")
			log.Println("数据采集 - 登录 AC - 成功完成")
			return
		} else {
			// continue
			log.Fatalln("登录流程认证失败，检查config.ini中auth是否正确！")
		}
		return
	}
}
