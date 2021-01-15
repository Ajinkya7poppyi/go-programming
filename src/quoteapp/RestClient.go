package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, er := http.Get("https://favqs.com/api/qotd.json")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
