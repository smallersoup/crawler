package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

/**
爬取网络资源函数
*/
func Fetch(url string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest is err ", err)
		return nil, fmt.Errorf("NewRequest is err %v\n", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	//返送请求获取返回结果
	resp, err := client.Do(req)

	//直接用http.Get(url)进行获取信息，爬取时可能返回403，禁止访问
	//resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error: http Get, err is %v\n", err)
	}

	//关闭response body
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: StatusCode is %d\n", resp.StatusCode)
	}

	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(bodyReader, determineEncoding(bodyReader).NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

/**
确认编码格式
*/
func determineEncoding(r *bufio.Reader) encoding.Encoding {

	//这里的r读取完得保证resp.Body还可读
	body, err := r.Peek(1024)

	//如果解析编码类型时遇到错误,返回UTF-8
	if err != nil {
		log.Printf("determineEncoding error is %v", err)
		return unicode.UTF8
	}

	//这里简化,不取是否确认
	e, _, _ := charset.DetermineEncoding(body, "")
	return e
}
