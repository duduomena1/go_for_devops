package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseUrl := "192.168.1.25"

	// GET request

	getResp, err := http.Get(baseUrl + "/get")
	if err != nil {
		fmt.Println("Error while making GET request:", err)
		return
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)
	fmt.Println("GET Response:", string(getBody))

	// POST request

	postBody := []byte(`{"key": "value"}`)
	postResp, err := http.Post(baseUrl+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error on POST request:", err)
		return
	}
	defer postResp.Body.Close()
	postBodyResp, _ := io.ReadAll(postResp.Body)
	fmt.Println("POST Response:", string(postBodyResp))

	// PUT request

	client := &http.Client{}
	putBody := []byte(`{"key": "update value"}`)
	putReq, err := http.NewRequest("PUT", baseUrl+"/put", bytes.NewBuffer(putBody))
	putReq.Header.Set("Content-Type", "application/json")
	putResp, err := client.Do(putReq)
	if err != nil {
		fmt.Println("Error on PUT request:", err)
		return
	}
	defer putResp.Body.Close()
	putBodyResp, _ := io.ReadAll(putResp.Body)
	fmt.Println("PUT Response:", string(putBodyResp))

	// DELETE request

	deleteReq, err := http.NewRequest(http.MethodDelete, baseUrl+"/delete", nil)
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		fmt.Println("Error on DELETE request:", err)
		return
	}
	defer deleteResp.Body.Close()
	deleteBodyResp, _ := io.ReadAll(deleteResp.Body)
	fmt.Println("DELETE Response:", string(deleteBodyResp))
}
