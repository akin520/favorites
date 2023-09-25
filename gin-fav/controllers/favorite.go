package controllers

import (
	"fmt"
	"gin-fav/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	fmt.Println("in index")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	var favs []*models.Favorite
	var pageP int
	var pageCount int
	var nextPage bool
	var err error
	favs, page, pageCount, nextPage, err = models.GetAllFavorite(page)
	log.Println(favs, page, pageCount, nextPage, err)
	if err != nil {
		log.Println(err)
	}
	if page-1 <= 1 {
		pageP = 1

	} else {
		pageP = page - 1

	}
	c.HTML(http.StatusOK, "favorite.html", gin.H{
		"favs":      favs,
		"page":      page,
		"pageCount": pageCount,
		"nextPage":  nextPage,
		"isIndex":   true,
		"PageP":     pageP,
		"PageN":     page + 1,
	})

}

func Add(c *gin.Context) {
	fmt.Println("in add")
	title := c.Query("title")
	url := c.Query("url")
	c.HTML(http.StatusOK, "add_fav.html", gin.H{
		"Url":   url,
		"Title": title,
	})
}

func AddPost(c *gin.Context) {
	fmt.Println("in add post")
	title := c.PostForm("title")
	url := c.PostForm("url")
	if len(title) > 0 && len(url) > 0 {
		err := models.AddFavorite(title, url)
		if err != nil {
			return
		}
		c.Redirect(302, "/")
		c.Abort()
	} else {
		c.HTML(http.StatusOK, "add_fav.html", gin.H{
			"Url":   url,
			"Title": title,
			"err":   "参数错误",
			"isErr": true,
		})
		c.Abort()
	}

}

func Search(c *gin.Context) {
	fmt.Println("in search")
	keyword := c.PostForm("keyword")
	favs, err := models.FindFavorite(keyword)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "favorite.html", gin.H{
		"favs": favs,
	})
}
