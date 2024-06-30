// Source: https://gosamples.dev/join-url-elements/

package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	hostname := "https://gosamples.dev/"
	path1 := "//join-url-elements//"
	path2 := "/foo/"
	path3 := "//bar"

	result, err := url.JoinPath(hostname, path1, path2, path3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
