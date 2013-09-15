package main

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"os"
	"strings"
	"fmt"
)

const (
	WOLFRAM_URL = "http://api.wolframalpha.com/v2/query"
	MAX_RESULTS = 5
)

var (
	appid = os.Getenv("q")
	postUrl = WOLFRAM_URL + 
		"?format=html"+
		"&appid=" + appid
)

type Pod struct {
	Markup string `xml:"markup"`
}

type QueryResult struct {
    XMLName xml.Name `xml:"queryresult"`
    Pods   []Pod  `xml:"pod"`
}

func wolframSearch(query string) string {
	resp, err := http.Get(postUrl + "&input=" + url.QueryEscape(query))
	
	fmt.Println(resp.Body)
	
	if err != nil {
		fmt.Printf("Error in HTTP GET: %s", err)
		return "error"
	}
	
	decoder := xml.NewDecoder(resp.Body)
	results := new(QueryResult)
	
	decoder.Decode(results)

	fmt.Println((*results).Pods[0].Markup)
	
	output := ""
	pods := (*results).Pods
	
	for i := range pods {
		if i > 5 {
			break
		}
		
		output += markup(pods[i].Markup)
		output += "<br>"
	}
		
	
	return output
}

func markup(data string) string {
	replacer := strings.NewReplacer("<![CDATA[", "", "]]>", "")
	return replacer.Replace(data)
}




