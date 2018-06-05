package main

import (
	"os"
	"io/ioutil"
	"regexp"
	"strconv"
	"io"
	"fmt"
)

var(
	regAttr = `"title":".+?"`
	regHref = `"link":".+?"`
)

func main() {
	fJson,_ := os.OpenFile("./data/json.txt",os.O_RDONLY, 0666)
	fTxt,_ := os.OpenFile("./data/info.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0667)
	body,_ := ioutil.ReadAll(fJson)
	defer fJson.Close()
	regTitle := regexp.MustCompile(regAttr)
	regHref := regexp.MustCompile(regHref)
	titleList := regTitle.FindAllString(string(body),-1)
	hrefList := regHref.FindAllString(string(body),-1)
	if len(titleList)==len(hrefList) {
		n := len(titleList)
		for i := 0;i < n ;i++  {
			str := strconv.Itoa(i+1)+"-"+titleList[i][9:len(titleList[i])-1]+hrefList[i][8:len(hrefList[i])-1]+"\n"
			io.WriteString(fTxt,str)
			fmt.Println(i+1)
		}
	}else {
		fmt.Printf("解析错误\n")
	}


}