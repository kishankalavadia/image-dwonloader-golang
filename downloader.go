// author : Kishan Kalavadia
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var filename, imagename string
	fmt.Println("Enter image-name for save: ")
	fmt.Scanf("%s\n", &imagename)
	fmt.Println("Enter filename of image-links: ")
	fmt.Println("note: separate link using comma(,) in txt file.")
	fmt.Scanf("%s\n", &filename)
	fileByteData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileStringData := string(fileByteData)
	re := regexp.MustCompile(`\r?\n`)
	fileStringData = re.ReplaceAllString(fileStringData, "")

	urls := strings.Split(fileStringData, ",")
	// don't worry about errors
	for i := 0; i < len(urls); i++ {
		response, e := http.Get(urls[i])
		if e != nil {
			log.Fatal(e)
		}

		defer response.Body.Close()
		//open a file for writing
		os.Mkdir("./"+imagename, os.ModePerm)
		file, err := os.Create("./" + imagename + "/" + imagename + "-" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}
		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		fmt.Printf("Success %d!\n", i)
	}
}
