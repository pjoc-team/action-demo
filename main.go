package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}).
		GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("hello!\n"))
			//c.String(http.StatusOK, "", "hello!\n")
		})
	err := r.Run(":8081")

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	_, err := writer.Write([]byte("hello!\n"))
	//	if err != nil {
	//		fmt.Printf("write error: %v \n", err.Error())
	//	}
	//})
	//err := http.ListenAndServe(":8081", nil)
	panic(err)
}
