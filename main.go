package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("hello!\n"))
		if err != nil {
			fmt.Printf("write error: %v \n", err.Error())
		}
	})
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
