package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

var logger service.Logger

func PreventCheckRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("stopped redirects")
}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	log.Println("开始运行run")
	sleepSeconds := 10 * time.Second
	loopCount := 0
	for {
		//if loopCount > 24 {
		//	f, err := os.Create("D:/GoAutoLoginZzuliStudent.txt")
		//	if err == nil {
		//		_ = f.Close()
		//	}
		//}
		jar, _ := cookiejar.New(nil)
		httpClient := &http.Client{
			CheckRedirect: PreventCheckRedirect,
			Jar:           jar,
		}
		//"http://1.1.1.1/?isReback=1"
		request, _ := http.NewRequest("GET", "http://www.msftconnecttest.com/redirect", nil)
		response, err := httpClient.Do(request)
		if err != nil {
			if !strings.Contains(fmt.Sprintf("%v", err), "stopped redirects") {
				log.Println("<errorMsg>无需认证，成功连接</errorMsg>")
				loopCount++
				time.Sleep(sleepSeconds)
				continue
			}
		}
		_ = response.Body.Close()
		redirectUrl := strings.Join(response.Header["Location"], "")
		if !strings.Contains(redirectUrl, "go.microsoft.com") {
			urlParsed, _ := url.Parse(redirectUrl)
			queryParam, _ := url.ParseQuery(urlParsed.RawQuery)

			loginUrl := "http://10.168.6.10:801/eportal/"
			params := url.Values{
				"c":           {"ACSetting"},
				"a":           {"Login"},
				"protocol":    {"http:"},
				"hostname":    {"10.168.6.10"},
				"iTermType":   {"1"},
				"wlanuserip":  {strings.Join(queryParam["wlanuserip"], "")},
				"wlanacip":    {strings.Join(queryParam["wlanacip"], "")},
				"mac":         {"00-00-00-00-00-00"},
				"ip":          {strings.Join(queryParam["wlanuserip"], "")},
				"enAdvert":    {"0"},
				"queryACIP":   {"0"},
				"loginMethod": {"1"},
			}
			loginUrl = loginUrl + "?" + params.Encode()

			username := os.Args[2]
			password := os.Args[3]
			accountType := os.Args[4]
			postForm := url.Values{}
			postForm.Add("DDDDD", fmt.Sprintf(",0,%s%s", username, accountType))
			postForm.Add("upass", fmt.Sprintf("%s", password))
			loginResponse, _ := httpClient.PostForm(loginUrl, postForm)
			_ = loginResponse.Body.Close()
			urlParsed, _ = url.Parse(strings.Join(loginResponse.Header["Location"], ""))
			queryParam, _ = url.ParseQuery(urlParsed.RawQuery)
			decodeBytes, _ := base64.StdEncoding.DecodeString(strings.Join(queryParam["ErrorMsg"], ""))
			errorMsg := string(decodeBytes)

			log.Printf("<errorMsg>%s</errorMsg>\n", errorMsg)
		} else {
			log.Println("<errorMsg>无需认证，成功连接</errorMsg>")
		}
		log.Printf("<loopCount>%d</loopCount>", loopCount)

		loopCount++
		time.Sleep(sleepSeconds)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func init() {
	f, err := os.Create("D:/GoAutoLoginZzuliStudent.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	var args []string
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			args = append(args, "run")
			var username, password, accountType string
			fmt.Print("请输入用户名：")
			_, _ = fmt.Scanln(&username)
			args = append(args, username)
			fmt.Print("请输入密码：")
			_, _ = fmt.Scanln(&password)
			args = append(args, password)
			fmt.Print("请输入账户类型：")
			_, _ = fmt.Scanln(&accountType)
			switch accountType {
			case "校园网":
				accountType = "@zzulis"
			case "校园移动":
				accountType = "@cmcc"
			case "校园联通":
				accountType = "@unicom"
			case "校园单宽":
				accountType = "@other"
			default:
				fmt.Println("您输入的账户类型不正确……")
				return
			}
			args = append(args, accountType)
		}
	}

	svcConfig := &service.Config{
		Name:        "GoAutoLoginZzuliStudent",
		DisplayName: "GoAutoLoginZzuliStudent",
		Description: "This is an Go service.",
		Arguments:   args,
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)

	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "start":
			_ = s.Start()
			fmt.Println("服务启动成功")
			return
		case "stop":
			_ = s.Stop()
			fmt.Println("服务停止成功")
			return
		case "restart":
			_ = s.Restart()
			fmt.Println("服务重启成功")
			return
		case "status":
			status, _ := s.Status()
			switch status {
			case service.StatusUnknown:
				fmt.Println("服务未安装")
			case service.StatusRunning:
				fmt.Println("服务运行中")
			case service.StatusStopped:
				fmt.Println("服务已停止")
			}
			return
		case "install":
			_ = s.Install()
			fmt.Println("服务安装成功")
			return
		case "uninstall":
			_ = s.Uninstall()
			fmt.Println("服务卸载成功")
			return
		case "run":
			if len(os.Args) < 5 {
				fmt.Println("参数不足")
				return
			}
		default:
			fmt.Println("参数有误")
			return
		}
	}
	fmt.Println("开始运行……")
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
