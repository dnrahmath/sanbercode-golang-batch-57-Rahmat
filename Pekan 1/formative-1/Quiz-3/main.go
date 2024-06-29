package main

import (
	"Quiz-3/controllers"
	"Quiz-3/database"
	"Quiz-3/middleware"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	} else {
		fmt.Println("success load file environtment")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	fmt.Println(psqlInfo)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	//CREATE DATABASE "quiz-3";
	database.DbMigrate(DB)

	defer DB.Close()

	// Router GIN
	router := gin.Default()
	router.Use(middleware.Auth())
	//------- No2 -------
	no2 := router.Group("/bangun-datar")
	{
		no2.GET("/segitiga-sama-sisi", func(c *gin.Context) {
			hitung := c.Query("hitung")
			alas := c.Query("alas")
			tinggi := c.Query("tinggi")
			controllers.Segitiga(c, hitung, alas, tinggi)
		})
		no2.GET("/persegi", func(c *gin.Context) {
			hitung := c.Query("hitung")
			sisi := c.Query("sisi")
			controllers.Persegi(c, hitung, sisi)
		})
		no2.GET("/persegi-panjang", func(c *gin.Context) {
			hitung := c.Query("hitung")
			panjang := c.Query("panjang")
			lebar := c.Query("lebar")
			controllers.PersegiPanjang(c, hitung, panjang, lebar)
		})
		no2.GET("/lingkaran", func(c *gin.Context) {
			hitung := c.Query("hitung")
			jariJari := c.Query("jariJari")
			controllers.Lingkaran(c, hitung, jariJari)
		})
	}
	//------- No3 -------
	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", controllers.InsertCategories)
	router.PUT("/categories/:id", controllers.UpdateCategories)
	router.DELETE("/categories/:id", controllers.DeleteCategories)
	router.GET("/categories/:id/books", func(c *gin.Context) {
		id := c.Param("id")
		categoryId, _ := strconv.Atoi(id)
		controllers.GetBookByIdCategory(c, categoryId)
	})
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", controllers.InsertBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:8080")
}
