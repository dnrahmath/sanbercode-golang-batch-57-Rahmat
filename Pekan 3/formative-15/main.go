package main

import (
	"database/sql"
	"fmt"
	"formative-15/controllers"
	"formative-15/database"
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
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run("localhost:8080")
}
