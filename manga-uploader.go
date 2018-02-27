package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	regex             = ".*v[^\\d]*?(\\.?\\d+(?:\\.\\d+)*[a-zA-Z]?).*?c[^\\d]*?(\\.?\\d+(?:\\.\\d+)*[a-zA-Z]?).*?(?:\\[(.+)\\])?\\.(?:zip|cbz)$"
	mangaDirectory    = "G:/mangadex-uploads/"
	finishedDirectory = "G:/mangadex-uploads/done/"
	defaultGroup      = 2
	defaultLang       = 1
	sessionToken      = ""
	uploadURL         = "https://mangadex.com/ajax/actions.ajax.php?function=chapter_upload"
)

func main() {
	fmt.Println("Starting the application...")

	re := regexp.MustCompile(regex)
	filenameProcessed := re.FindStringSubmatch("v.0 c.1 [bakabt].zip")
	volume := filenameProcessed[1]
	chapter := filenameProcessed[2]
	group := filenameProcessed[3]
	// Test regex
	fmt.Println(volume, chapter, group)

	// Function for posting upload
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	request.Header.Set("Cookie", "mangadex="+sessionToken)
	// Needed to get MangaDex to accept XMLHttpRequest/POST requests
	request.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	fmt.Println("Terminating the application...")
}
