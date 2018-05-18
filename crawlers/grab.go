package crawlers

import (
	"net/http"
	"github.com/mozillazg/request"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

/**
 * 抓取数据操作
 */


 var (
 	//设置cookies
 	cookies = make(map[string]string)
 	//设置请求参数
	params = make(map[string]string)
	//请求对象
	httpClient = new(http.Client)
	//设置请求头,初始化一些参数
	headers = map[string]string{
		 "Accept":           "multipart/form-data",
		 "Accept-Encoding":  "gzip, deflate",
		 "Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		 "Content-Length":   "219",
		 "Content-type":     "application/x-www-form-urlencoded",
		 "User-Agent":       "Chrome/65.0.0",
	 }
 )


 /**
  * 设置请求头
  */
 func SetHeader(header map[string]string){
	 for key, value := range header {
		 headers[key] = value
	 }
 }

 /**
  * 设置cookies
  */
 func SetCookies(cookie map[string]string) {
	 for key, value := range cookie {
		 cookies[key] = value
	 }
 }
 /**
  * 设置请求参数
  */
 func SetParams(param map[string]string){
	 for key, value := range param {
		 params[key] = value
	 }
 }
 /**
  * 获取数据
  */
 func GetData(url string)(data string,err error){
	 req := request.NewRequest(httpClient)
	 req.Cookies = cookies
	 req.Params = params
	 req.Headers = headers
 	 resp, error := req.Get(url)
	 defer resp.Body.Close()
	 if error!=nil {
		 fmt.Println(error)
		 return "",error
	 }
	 doc, err := goquery.NewDocumentFromReader(resp.Body)
	 if err!=nil {
		 fmt.Println(error)
		 return "",error
	 }

	 var buf bytes.Buffer
     doc.Each(func(i int, selection *goquery.Selection) {
         buf.WriteString(selection.Text())
	 })
     return buf.String(),nil
 }

 /**
  * 获取数据
  */
 func PostData(url string)(data string,err error){
	 req := request.NewRequest(httpClient)
	 req.Cookies = cookies
	 req.Params = params
	 req.Headers = headers
	 resp, error := req.Post(url)
	 defer resp.Body.Close()
	 if error!=nil {
		 fmt.Println(error)
		 return "",error
	 }

	 doc, err := goquery.NewDocumentFromReader(resp.Body)
	 if err!=nil {
		 fmt.Println(error)
		 return "",error
	 }

	 var buf bytes.Buffer
	 doc.Each(func(i int, selection *goquery.Selection) {
	 	//fmt.Println(mahonia.NewEncoder("GBK").ConvertString(selection.Text()))
		 buf.WriteString(selection.Text())
	 })
	 return buf.String(),nil
 }

