package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://github.com/miku/zek
// for i in $(ls -1); do zek -c -n ${i%.xml} "./$i" >> "bbb.go"; done

type Namesilo struct {
	XMLName xml.Name `xml:"namesilo"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text      string `xml:",chardata"`
		Operation string `xml:"operation"`
		Ip        string `xml:"ip"`
	} `xml:"request"`
	Reply struct {
		Text   string `xml:",chardata"`
		Code   string `xml:"code"`
		Detail string `xml:"detail"`
	} `xml:"reply"`
}

func main() {
	resp, err := http.DefaultClient.Get("https://www.namesilo.com/api_reference.php")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("./temp.html", bytes, 0666)

	file, err := os.Open("/home/ldez/sources/go/src/github.com/ldez/jaelon/namesiloscrapper/bla.html")

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	countPre := 0

	doc.Find("#innerContentContainerLeft pre").Each(func(i int, s *goquery.Selection) {
		namesilo := Namesilo{}
		err := xml.Unmarshal([]byte(s.Text()), &namesilo)
		if err != nil {
			log.Fatal(err)
		}

		countPre += 1

		ioutil.WriteFile(fmt.Sprintf("./samples/%s.xml", namesilo.Request.Operation), []byte(s.Text()), 0666)
	})

	countSample := 0

	doc.Find("#innerContentContainerLeft .apiRequestExamples").Each(func(i int, s *goquery.Selection) {
		uri := strings.ReplaceAll(strings.TrimSpace(strings.ReplaceAll(s.Text(), "\n", "")), " ", "")

		parse, err := url.Parse(uri)
		if err != nil {
			log.Fatal(err)
		}

		opName := path.Base(parse.Path)
		countSample += 1
		fmt.Println(opName)
	})

	countParams := 0

	doc.Find("#innerContentContainerLeft ul").Each(func(i int, s *goquery.Selection) {
		if s.Find(".highlight").Length() > 0 {
			return
		}

		if s.Find("li > b").Length() > 0 {
			countParams += 1
			fmt.Println("---------------")
			fmt.Println(s.Html())
			fmt.Println("---------------")
		}
	})

	fmt.Println(countPre, countSample, countParams)
}
