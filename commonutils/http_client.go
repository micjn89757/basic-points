package commonutils

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetPractice get请求示例
func GetPractice() (string, error) {
	resp, err := http.Get("https://www.liwenzhou.com")

	if err != nil {
		log.Printf("get failed, err:%v\n", err)
		return "", err
	}

	defer resp.Body.Close() // 使用完resp后必须关闭resp的主体

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read from resp.Body failed, err:%v\n", err)
		return "", err
	}

	log.Printf("%v\n", string(body))
	return string(body), nil
}

// GetValuesPractice 带参数的Get请求示例
func GetValuesPractice() (string, error) {
	apiUrl := "https://www.asus.com.cn/laptops/for-gaming/tuf-gaming/asus-tuf-gaming-f16-2024/helpdesk_bios/"
	// URL param
	data := url.Values{} // GET请求的参数需要使用Go语言内置的net/url处理
	data.Set("model2Name", "FX607JVR")
	u, err := url.ParseRequestURI(apiUrl)

	if err != nil {
		return "", err
	}

	log.Printf("%v\n", u.String())

	resp, err := http.Get(u.String())

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// PostPractice post请求示例
func PostPractice() (string, error) {
	url := "http://127.0.0.1:4523/m1/4024071-0-default/pet"

	// post请求 content type类型
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"
	// json

	contentType := "application/json"
	data := `{"name": "Hello Kitty", "status": "sold"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))

	if err != nil {
		log.Printf("post failed, err:%v\n", err)
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("get resp failed, err:%v\n", err)
		return "", err
	}

	return string(b), nil
}
