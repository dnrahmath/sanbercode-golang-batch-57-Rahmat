package main

import (
	"Quiz-3/controllers"
	"Quiz-3/database"
	"Quiz-3/middleware"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "practice"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// psqlInfo := fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// psqlInfo := fmt.Sprintf("host=#{host} port=#{port} user=#{user} password=#{password} dbname=#{dbname} sslmode=disable")

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	} else {
		fmt.Println("success load file environtment")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
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

	database.DbMigrate(DB)

	defer DB.Close()

	// Router GIN
	router := gin.Default()
	router.Use(middleware.Auth())
	//------- No2 -------
	router.GET("/segitiga-sama-sisi", func(c *gin.Context) {
		hitung := c.Query("hitung")
		alas := c.Query("alas")
		tinggi := c.Query("tinggi")
		controllers.Segitiga(c, hitung, alas, tinggi)
	})
	router.GET("/persegi", func(c *gin.Context) {
		hitung := c.Query("hitung")
		sisi := c.Query("sisi")
		controllers.Persegi(c, hitung, sisi)
	})
	router.GET("/persegi-panjang ", func(c *gin.Context) {
		hitung := c.Query("hitung")
		panjang := c.Query("panjang")
		lebar := c.Query("lebar")
		controllers.PersegiPanjang(c, hitung, panjang, lebar)
	})
	router.GET("/lingkaran", func(c *gin.Context) {
		hitung := c.Query("hitung")
		jariJari := c.Query("jariJari")
		controllers.Lingkaran(c, hitung, jariJari)
	})
	//------- No3 -------
	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", controllers.InsertCategories)
	router.PUT("/categories/:id", controllers.UpdateCategories)
	router.DELETE("/categories/:id", controllers.DeleteCategories)
	router.GET("/categories/:id/books", controllers.DeleteCategories)
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", controllers.InsertBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:8080")
}
