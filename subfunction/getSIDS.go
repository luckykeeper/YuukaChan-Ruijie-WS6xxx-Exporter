// 优香酱锐捷 AC Exporter ，适配 WS6xxx 系列设备 - 登录并获取SIDS
// @CreateTime : 2023/8/28 10:52
// @LastModified : 2023/9/06 13:20
// @Author : Luckykeeper
// @Contact : luckykeeper@luckykeeper.site | https://github.com/luckykeeper | https://luckykeeper.site
// @ProgramEntry: yuukaExporter.go
// @Project : yuukaChan-Ruijie-WS6xxx-Exporter

package subFunction

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"yuukaChan-Ruijie-WS6xxx-Exporter/model"

	"github.com/goccy/go-json"
	"golang.org/x/crypto/pbkdf2"
)

// 登录并获取SIDS
func GetSIDS(DeviceIP string, DeviceUser, DeviceAuth string) (SIDS string) {
	for {
		log.Println("开始登录流程！")

		// Step 1: 请求 hmac_info 接口，取得加密用盐值
		hmacInfoUrl := "https://" + DeviceIP + "/hmac_info.do?user=" + DeviceUser
		method := "GET"

		//构建https请求，忽略证书错误
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		hmacInfoClient := &http.Client{Transport: tr}
		hmacInfoReq, err := http.NewRequest(method, hmacInfoUrl, nil)

		if err != nil {
			log.Println(err)
			continue
		}
		hmacInfoRes, err := hmacInfoClient.Do(hmacInfoReq)
		if err != nil {
			log.Println(err)
			continue
		}
		defer hmacInfoRes.Body.Close()

		hmacInfoBody, err := io.ReadAll(hmacInfoRes.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		var hmacInfo model.HmacInfo
		json.Unmarshal(hmacInfoBody, &hmacInfo)

		// 锐捷加密算法
		ruijiePassword := genRuijiePassword(DeviceAuth, hmacInfo)

		// 定义登录相关参数
		loginUrl := "https://" + DeviceIP + "/login.do"

		// 下面注释掉的是旧版的 authData ，说明请参考 config.ini
		// authData := `auth=` + DeviceAuth
		authData := "user=" + DeviceUser + "&key=" + ruijiePassword
		authDataByte := []byte(authData) // 需要注意的是，因为锐捷AC需要的不是标准形式的 JSON 字符串，
		// 所以不能用json序列化之后转[]byte而是直接转[]byte
		// fmt.Println(authData)

		loginParam := bytes.NewBuffer(authDataByte)
		// fmt.Println(loginParam)

		//构建https请求，忽略证书错误
		// tr := &http.Transport{
		// 	TLSClientConfig: &tls.Config{
		// 		InsecureSkipVerify: true,
		// 	},
		// }

		loginClient := &http.Client{Transport: tr}
		loginReq, err := http.NewRequest("POST", loginUrl, loginParam)

		if err != nil {
			log.Println(err)
			continue
		}
		//header
		loginReq.Header.Add("Content-Type", "application/json")

		//发送请求
		loginRes, err := loginClient.Do(loginReq)
		if err != nil {
			log.Println(err)
			continue
		}
		defer loginRes.Body.Close()

		//返回结果
		loginBody, err := io.ReadAll(loginRes.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		cookie := loginRes.Header.Get("Set-Cookie")
		// fmt.Println(string(body))
		bodyResult := strings.Split((strings.Split(string(loginBody), "<return-code>")[1]), "</return-code>")[0]
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

// 锐捷加密算法
func genRuijiePassword(password string, hmacInfo model.HmacInfo) string {
	hash := pbkdf2.Key([]byte(password), []byte(hmacInfo.Salt), hmacInfo.Iter, hmacInfo.Keylen, sha256.New)
	return hex.EncodeToString(hash)
}
