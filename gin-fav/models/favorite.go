package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math"
)

type Favorite struct {
	gorm.Model
	Title string `orm:"unique"`
	Url   string
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func openDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	CheckError(err)
	return db
}

func init() {
	db := openDB()
	db.AutoMigrate(&Favorite{})
	//for i := 0; i < 10; i++ {
	//	title := fmt.Sprintf("title %d", i)
	//	url := fmt.Sprintf("http://www.baidu.com/%d", i)
	//	AddFavorite(title, url)
	//}
}

func AddFavorite(title string, url string) error {
	db := openDB()
	fav := Favorite{Title: title, Url: url}
	result := db.Create(&fav)
	return result.Error
}

func FindFavorite(keyword string) ([]*Favorite, error) {
	db := openDB()
	var favs []*Favorite
	//key := fmt.Sprintf("\%%s\%", keyword)
	log.Println(keyword)
	//result := db.Model(&Favorite{}).Where("title LIKE ?", "%"+keyword+"%").Find(favs)
	result := db.Raw("select * from favorites where title like '%" + keyword + "%'").Scan(&favs)
	return favs, result.Error
}

func GetAllFavorite(page int) ([]*Favorite, int, int, bool, error) {
	db := openDB()
	var favs []*Favorite
	var count int64
	db.Model(&Favorite{}).Count(&count)
	fmt.Println(count)
	size := 10
	pageCount := math.Ceil((float64(count) / float64(size)))
	result := db.Offset((page - 1) * size).Limit(size).Find(&favs)
	var nextPage bool
	if page < int(pageCount) {
		nextPage = true
	}
	log.Println("GetAllFavorite:", favs, page, int(pageCount), nextPage)
	return favs, page, int(pageCount), nextPage, result.Error
}
