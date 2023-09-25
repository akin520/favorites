package main

import (
	"fmt"
	"gin-fav/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/", controllers.Index)
	r.POST("/fav_search", controllers.Search)
	r.GET("/fav_add", controllers.Add)
	r.POST("/fav_add", controllers.AddPost)
	r.Run()
}
