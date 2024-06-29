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

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	Books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": Books,
		}
	}

	c.JSON(http.StatusOK, result)
}

//===================================================================

func InsertBook(c *gin.Context) {
	// var ReqBook structs.ReqBook
	var Book structs.Book

	err := c.ShouldBindJSON(&Book)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBook(database.DbConnection, Book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

//===================================================================

func UpdateBook(c *gin.Context) {
	// var ReqBook structs.ReqBook
	var Book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&Book)
	if err != nil {
		panic(err)
	}

	Book.ID = int(id)

	err = repository.UpdateBook(database.DbConnection, Book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

//===================================================================

func DeleteBook(c *gin.Context) {
	var Book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	Book.ID = int(id)

	err := repository.DeleteBook(database.DbConnection, Book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Book",
	})
}

//===================================================================
