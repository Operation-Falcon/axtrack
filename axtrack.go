package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Deepanjalkumar/axtrack/banner"
)

func Extract_urls(domain, file string) {

	r, err := http.Get(fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=%s/*&output=txt&fl=original&collapse=urlkey", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	data := string(url_byte)

	filename, _ := os.Create(file)

	filename.WriteString(fmt.Sprintf("%s\n", data))

	fmt.Printf("%s\n", data)
}

func main() {

	domain := flag.String("domain", "", "domain name to extract all urls")

	output := flag.String("output", "", "output filename to save the results")

	flag.Parse()

	banner.Banner_design()

	Extract_urls(*domain, *output)
}
