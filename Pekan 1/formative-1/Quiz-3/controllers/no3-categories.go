package controllers

import (
	"Quiz-3/database"
	"Quiz-3/repository"
	"Quiz-3/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//===================================================================

func GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	Categoriess, err := repository.GetAllCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": Categoriess,
		}
	}

	c.JSON(http.StatusOK, result)
}

//===================================================================

func InsertCategories(c *gin.Context) {
	var Categories structs.Category

	err := c.ShouldBindJSON(&Categories)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategories(database.DbConnection, Categories)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Categories",
	})
}

//===================================================================

func UpdateCategories(c *gin.Context) {
	var Categories structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&Categories)
	if err != nil {
		panic(err)
	}

	Categories.ID = int(id)

	err = repository.UpdateCategories(database.DbConnection, Categories)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Categories",
	})
}

//===================================================================

func DeleteCategories(c *gin.Context) {
	var Categories structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	Categories.ID = int(id)

	err := repository.DeleteCategories(database.DbConnection, Categories)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Categories",
	})
}

//===================================================================
