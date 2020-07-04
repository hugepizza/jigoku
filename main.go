package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("tpl/jigoku.html")
	r.StaticFS("/tpl", http.Dir("./tpl"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "jigoku.html", "1")
	})
	r.GET("/send", func(c *gin.Context) {
		fmt.Println("===========")
		fmt.Println(c.Request.RemoteAddr + "  " + c.Query("message"))
	})
	r.Run(":80")
}
