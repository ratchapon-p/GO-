package main

import "fmt"

func main() {
	website := map[string]string{
		"Google":             "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Amazon Web Services"])
	website["LinkedIn"] = "https://linkedin.com"
	fmt.Println(website)

	delete(website, "Google")
	fmt.Println(website)


}