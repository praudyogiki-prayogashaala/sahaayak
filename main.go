package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	query *string
)

func init() {
	query = flag.String("wiki", "", "query string")
	flag.Parse()
}	

func main() {
	if *query == "" {
		fmt.Println("no input provided, sahaayak going down")
		return
	}

	addr := "http://en.wikipedia.org/wiki/" + *query

	res, _ := http.Get(addr)
	out, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	text := string(out[:])

	start := strings.Index(text, "<p>")
	end := strings.Index(text, "</p>")

	summary := text[start:end]

	result := ""
	valid := 0

	for i := 0; i < len(summary); i++ {
		c := string(summary[i:i+1])
		if strings.Compare(c,"<") == 0 || strings.Compare(c,"[") == 0{
			valid += 1
		} else if strings.Compare(c,">") == 0 || strings.Compare(c,"]") == 0{
			valid -= 1
		} else if valid==0 {
			result = result + c
		}
	}

	fmt.Println(result)
}
