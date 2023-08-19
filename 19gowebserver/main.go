package main

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strings"
)

func main() {
	// Make sure to start the server first! (webserver folder)
	fmt.Println("Web server in Golang!")
	PerformGetRequest("http://localhost:8000/get")
	PerformPostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")
}

func PerformGetRequest(url string) {
	fmt.Println("Performing GET request...")

	response, err := http.Get(url)
	checkErr(err, "resp")
	defer response.Body.Close() // close the connection, never forget this!

	// fmt.Println("Status code:", response.StatusCode)
	// fmt.Println("Content length:", response.ContentLength)

	// respBytes, err := io.ReadAll(response.Body)
	// checkErr(err, "respBytes")
	// fmt.Println("Body:", string(respBytes))

	var responseString strings.Builder
	respBytes, err := io.ReadAll(response.Body)
	responseString.Write(respBytes)

	fmt.Println("Body:", responseString.String())
}

func PerformPostJsonRequest(url string) {
	fmt.Println("Performing POST json request...")

	requestBody := strings.NewReader(`
		{
			"firstname": "Saksham",
			"lastname": "Karnawat"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	checkErr(err, "resp")

	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	checkErr(err, "respBytes")

	fmt.Println("Body:", string(respBytes))
}

func PerformPostFormRequest(url string) {
	fmt.Println("Performing POST form request...")

	data := neturl.Values{}
	data.Add("firstname", "Saksham")
	data.Add("lastname", "Karnawat")

	response, err := http.PostForm(url, data)
	checkErr(err, "resp")

	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)

	fmt.Println("Body:", string(respBytes))
}

func checkErr(err error, errmsg string) {
	if err != nil {
		fmt.Println("Error:", errmsg)
		panic(err)
	}
}
