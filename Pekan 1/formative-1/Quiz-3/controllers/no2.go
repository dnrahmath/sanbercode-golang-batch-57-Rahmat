package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//===================================================================

func Segitiga(c *gin.Context, hitung string, alas string, tinggi string) {
	if hitung == "Keliling" {

	} else if hitung == "Luas" {

	} else {
		//tidak diizinkan
	}
	result := gin.H{
		"hitung": hitung,
		"alas":   alas,
		"tinggi": tinggi,
	}
	c.JSON(http.StatusOK, result)
}

//===================================================================

func Persegi(c *gin.Context, hitung string, sisi string) {
	if hitung == "Keliling" {

	} else if hitung == "Luas" {

	} else {
		//tidak diizinkan
	}
	result := gin.H{
		"hitung": hitung,
		"sisi":   sisi,
	}
	c.JSON(http.StatusOK, result)
}

//===================================================================

func PersegiPanjang(c *gin.Context, hitung string, panjang string, lebar string) {
	if hitung == "Keliling" {

	} else if hitung == "Luas" {

	} else {
		//tidak diizinkan
	}
	result := gin.H{
		"hitung":  hitung,
		"panjang": panjang,
		"lebar":   lebar,
	}
	c.JSON(http.StatusOK, result)
}

//===================================================================

func Lingkaran(c *gin.Context, hitung string, jariJari string) {
	if hitung == "Keliling" {

	} else if hitung == "Luas" {

	} else {
		//tidak diizinkan
	}
	result := gin.H{
		"hitung":   hitung,
		"jariJari": jariJari,
	}
	c.JSON(http.StatusOK, result)
}

//===================================================================
