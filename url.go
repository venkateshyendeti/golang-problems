package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawurl := "https://www.google.com/search?q=postgresql+tutorial&oq=&gs_lcrp=EgZjaHJvbWUqCQgAEEUYOxjCAzIJCAAQRRg7GMIDMgkIARBFGDsYwgMyCQgCEEUYOxjCAzIJCAMQRRg7GMIDMgkIBBBFGDsYwgMyC"
	parsedurl, err := url.Parse(rawurl)
	if err != nil {
		fmt.Println("error parsed url:", err)
		return
	}
	fmt.Println("scheme", parsedurl.Scheme)
	fmt.Println("Host:", parsedurl.Host)
	fmt.Println("path:", parsedurl.Path)
	fmt.Println("raw query:", parsedurl.RawQuery)
	queryparms := parsedurl.Query()
	fmt.Println("query parameters :")
	for key, value := range queryparms {
		fmt.Printf("%s:%s\n", key, value)
	}
	queryparms.Set("q", "golang tutorial")
	queryparms.Add("page", "1")
	parsedurl.RawQuery = queryparms.Encode()
	newurl := parsedurl.String()
	fmt.Println("new url", newurl)
}
