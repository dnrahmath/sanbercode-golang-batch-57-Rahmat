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

func GetBookByIdCategory(c *gin.Context, categoryId int) {
	var (
		result gin.H
	)

	Books, err := repository.GetBooksByCategoryId(database.DbConnection, categoryId)

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

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	Books, err := repository.GetAllBooks(database.DbConnection)

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
	var ReqBook structs.ReqBook

	err := c.ShouldBindJSON(&ReqBook)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBook(database.DbConnection, ReqBook.Convert())
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

//===================================================================

func UpdateBook(c *gin.Context) {
	var ReqBook structs.ReqBook
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&ReqBook)
	if err != nil {
		panic(err)
	}

	Book := ReqBook.Convert()

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
