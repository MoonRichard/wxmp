package main

import (
	"io"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
	"strconv"
	"mahonia"
)
var (
	baseUrl1 = `https://mp.weixin.qq.com/cgi-bin/appmsg?token=1898813492&lang=zh_CN&f=json&ajax=1&random=0.9753501583058535`
	baseUrl2 = `&action=list_ex&begin=`
	baseUrl3 = `&count=5&query=&fakeid=MzIzNjI0MDU4OQ%3D%3D&type=9`
)
func main() {
	var baseUrl string
	f,_ := os.OpenFile("./data/json.txt",os.O_CREATE|os.O_RDWR,0)
	defer f.Close()
	cookie := "noticeLoginFlag=1; eas_sid=D1p5T1Y678L9w4M0w0v1E0W3y6; o_cookie=904162909; pac_uid=1_904162909; pgv_pvi=7972885504; pgv_pvid=8208406600; pgv_si=s9722078208; pt2gguin=o0904162909; ptcz=94219b769a3b5a9b031c873f53089083ffc9e3969203c2d3c34de96aa97c583f; ptui_loginuin=904162909@qq.com; RK=fEX3YCH6G2; tvfe_boss_uuid=02bcc81f21ac3368; cert=BUoaLvKrXPMWJCzucTuJ3qpQ2ADkPSLQ; data_bizuin=3558625347; data_ticket=VoI70EXtVa530ijbWaRboCzYQF88+g5W8FozOsDivAGAhCsF7WhfV5dpnCW/5XEk; mm_lang=zh_CN; noticeLoginFlag=1; openid2ticket_oQ1Ac1H7nKsgNBcNebZSUztLtIhc=I6b2QGz2whfzBmTI911KdXQx44C40BCzdg8R8OOb8kw=; ticket=1bfce5091afd2da47c3a2e1cd4f32e16bc870442; ticket_id=gh_d39016766e0d; ua_id=x0Gb8H858vwaN0i2AAAAAGkmICu4_P7K6FWyMw-4Frc=; uuid=4f0560f6dff7ed2582818d1fd495158f; xid=6d2c9eca3f6286c575fd92ec35df23ec; slave_user=gh_d39016766e0d; slave_sid=TU85SGR6R24yNjh2NklfeVlVaFhwQ2hZSjlnR2JJbzBXVDN0T1owZ3dsZW1BY1JJNEc0ektlM3JmTW9rd1JjZ2o2ZnJiNXhtVE0yaGV0SUd5UUhuOHVVQXdJRGlEYkh5RmlTb2JMbWlQeVVhbHJGY1JmN1dudnZkNExYUWxWODFQMUZlbkZLVzhNS3BIdGlw; bizuin=3571627047"
	for i:=0;i<= 685; i+=5 {
		baseUrl = baseUrl1+baseUrl2+strconv.Itoa(i)+baseUrl3
		byte := get(baseUrl,cookie)
		content := mahonia.NewDecoder("utf-8").ConvertString(string(byte))
		io.WriteString(f,content)
		println(i)
	}

}

//模拟登录与解析
func get(url string,cookie string) []byte {
	client := http.Client{
		CheckRedirect:nil,
	}
	req,_ := http.NewRequest("GET",url,nil)
	//设置header
	req.Header.Set("Accept","application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding","gzip, deflate, br")
	req.Header.Set("Accept-Language","zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36")
	req.Header.Set("X-Requested-With","XMLHttpRequest")
	//填充cookie
	if(len(cookie) > 0){
		array:=strings.Split(cookie,"; ")
		for item:= range array{
			array2:=strings.Split(array[item],"=")
			if(len(array2)==2){
				cookieObj:= http.Cookie{}
				cookieObj.Name=array2[0]
				cookieObj.Value=array2[1]
				req.AddCookie(&cookieObj)
			}else{
				println("error,index out of range:"+array[item])
			}
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err:%s",err)
		return nil
	}
	defer resp.Body.Close()
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Printf("err,:%s",err)
			return nil
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}
	if(reader!=nil){
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("err:%s",err)
			return nil
		}
		return body

	}
	return nil
}