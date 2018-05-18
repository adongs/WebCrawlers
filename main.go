package main

import (
	"WebCrawlers/crawlers"
	"time"
	"github.com/tidwall/gjson"
	"fmt"
)

func zhaopinggou(){
	crawlers.SetHeader(map[string]string{
		"Accept":           "multipart/form-data",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Content-Length":   "219",
		"Content-type":     "application/x-www-form-urlencoded",
		"Host":             "www.zhaopingou.com",
		"Origin":           "http://www.zhaopingou.com",
		"Proxy-Connection": "keep-alive",
		"Referer":          "http://www.zhaopingou.com/search?key=java",
		"User-Agent":       "Chrome/65.0.0",
	})
	crawlers.SetCookies(map[string]string{
		"JSESSIONID":     "F3D0CFA651F5335993AA18A8B8C85B90",
		"fangWenIp":      "123.125.5.5",
		"fanwenTime1":    time.Now().Format("2006-01-02 15:04:05"),
		"fangWenNumber1": "5",
	})
	crawlers.SetParams(map[string]string{
		"cityId":          "-1",
		"addressQu":       "",
		"strKey":          "",
		"workYears":       "",
		"degreesTypes":    "",
		"positionNature":  "",
		"companyNature":   "",
		"companyScaleId":  "",
		"companySeedtime": "",
		"monthType":       "",
		"monthStr":        "",
		"sPayMonth":       "",
		"ePayMonth":       "",
		"pageSize":        "0",
		"pageNo":          "1",
		"clientNo":        "",
		"userToken":       "",
		"clientType":      "2",
	})

	data,error := crawlers.PostData("http://www.zhaopingou.com/zhaopingou_interface/c_search_find_position?timestamp=")
	if error==nil {
		value :=gjson.Get(data,"positionReleaseList")
		for _, name := range value.Array() {
			println(name.String())
		}

	}
}

/**
 * 载入相关的规则和初始化应用程序
 */
func main() {


	crawlers.SetHeader(map[string]string{
		"Accept": "*/*",
		"Accept-Encoding": "",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Connection": "keep-alive",
		"Host": "recommender-api-ms.juejin.im",
		"Origin": "https://juejin.im",
		"Referer": "https://juejin.im/",
	})
	crawlers.SetCookies(map[string]string{

	})
	crawlers.SetParams(map[string]string{
		"suid": "IBvuBaZYmFQUVnRumqFu",
		"ab": "welcome_3",
		"src": "web",
	})

	data,error := crawlers.PostData("https://recommender-api-ms.juejin.im/v1/get_recommended_entry?suid=IBvuBaZYmFQUVnRumqFu&ab=welcome_3&src=web")
	if error==nil {
		fmt.Println(data)

	}




















	/*
	if err !=nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	//query := elastic.NewTermQuery("Sex","25")
	//query := elastic.NewMatchQuery("Name","于东")
	//query := elastic.NewMatchQuery("title","centos升级").Operator("or")
	//query := elastic.NewMultiMatchQuery("centos","title","abstract")
	query := elastic.NewCommonTermsQuery("title","编译").CutoffFrequency(0.0001).LowFreqOperator("or")
	//result,err_r :=client.Search().
	result,err_r :=client.Search().
		Index("website").
		Query(query).
		Do(ctx)
	if err_r!=nil {
		panic(err_r)
	}
	for _,data :=range result.Hits.Hits {
		source,_ :=json.Marshal(data.Source)
		fmt.Println(string(source))
	}
*/




/*	exists,err_index := client.IndexExists("user").Do(context.Background())
	client.IndexPutSettings("")
	if err_index!=nil {
		fmt.Println(err_index)
	}
	if !exists {
		fmt.Println("index不存在")
	}*/

/*type User struct {
	Name string
    Sex int
    Message string

}*/

/*	user:= User{Name:"于东",Sex:25,Message:"信息"}
	_,err_crete:=client.Index().
		                Index("user").
		                Type("user").
		                Id("1").
		                BodyJson(user).
		                Do(context.Background())
	if err_crete!=nil {
		fmt.Println(err_crete)
	}*/

	//termq := elastic.NewTermQuery("Name","于东")
 /*   result,err_result :=client.Get().Index("user").Type("user").
    Id("1").Do(context.Background())
	if err_result!=nil {
		panic(err_result)
	}
	if result.Found {
		fmt.Println(result)
	}*/
	//fmt.Printf("query took %d milliseconds \n",result.TookInMillis)
	//fmt.Println(result.Hits.TotalHits)

  /*  ctx := context.Background()
    //query := elastic.NewTermQuery("Sex","25")
	//query := elastic.NewMatchQuery("Name","于东")
	query := elastic.NewTermQuery("title","vmware")
    result,err_r :=client.Search().
    	            Index("website").
    	            Type("blog").
    	            Query(query).
    	            Pretty(true).
    	            Do(ctx)
	if err_r!=nil {
		panic(err_r)
	}*/
/*	fmt.Println(result)
   var users User
	for _,itme := range result.Each(reflect.TypeOf(users)) {
		t := itme.(User)
		fmt.Printf("%s  %s",t.Message,t.Name)
	}*/

}
