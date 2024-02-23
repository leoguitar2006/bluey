package main

import "github.com/gin-gonic/gin" // Sempre que não der um nome no import, pega o último nome da path

func main() {
	r := gin.Default() //:=
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
