package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("hello!"))
		if err != nil {
			fmt.Printf("write error: %v \n", err.Error())
		}
	})
}
