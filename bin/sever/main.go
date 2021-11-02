package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/onedss/onewebssh"
)

func main() {
	r := gin.Default()

	//跨域设置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1", "http://localhost:8080"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	confing := &onewebssh.WebSSHConfig{
		Record:     true,
		RecPath:    "./rec/",
		RemoteAddr: "localhost:22",
		User:       "wida",
		Password:   "wida",
		AuthModel:  onewebssh.PASSWORD,
	}

	handle := onewebssh.NewWebSSH(confing)

	r.GET("/ws/:id", handle.ServeConn)
	r.GET("/recoder", handle.RecoderList)
	r.Static("/static", "./front/dist/")
	r.Static("/rec", "./rec/") //录像回看目录
	r.LoadHTMLFiles("./front/dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":8080")
}
