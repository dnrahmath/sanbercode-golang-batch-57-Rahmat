package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

//===================================================================

func Segitiga(c *gin.Context, hitung string, alas string, tinggi string) {
	if hitung == "Keliling" || hitung == "keliling" {
		a, _ := strconv.ParseFloat(alas, 64)
		t, _ := strconv.ParseFloat(tinggi, 64)
		hasil := a + t + math.Sqrt(a*a+t*t)
		result := gin.H{
			"hitung": hitung,
			"alas":   alas,
			"tinggi": tinggi,
			"hasil":  hasil,
		}
		c.JSON(http.StatusOK, result)
	} else if hitung == "Luas" || hitung == "luas" {
		a, _ := strconv.ParseFloat(alas, 64)
		t, _ := strconv.ParseFloat(tinggi, 64)
		hasil := 0.5 * a * t
		result := gin.H{
			"hitung": hitung,
			"alas":   alas,
			"tinggi": tinggi,
			"hasil":  hasil,
		}
		c.JSON(http.StatusOK, result)
	} else {
		//tidak diizinkan
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operasi tidak diizinkan"})
	}
}

//===================================================================

func Persegi(c *gin.Context, hitung string, sisi string) {
	if hitung == "Keliling" || hitung == "keliling" {
		s, _ := strconv.ParseFloat(sisi, 64)
		hasil := 4 * s
		result := gin.H{
			"hitung": hitung,
			"sisi":   sisi,
			"hasil":  hasil,
		}
		c.JSON(http.StatusOK, result)
	} else if hitung == "Luas" || hitung == "luas" {
		s, _ := strconv.ParseFloat(sisi, 64)
		hasil := s * s
		result := gin.H{
			"hitung": hitung,
			"sisi":   sisi,
			"hasil":  hasil,
		}
		c.JSON(http.StatusOK, result)
	} else {
		//tidak diizinkan
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operasi tidak diizinkan"})
	}
}

//===================================================================

func PersegiPanjang(c *gin.Context, hitung string, panjang string, lebar string) {
	if hitung == "Keliling" || hitung == "keliling" {
		p, _ := strconv.ParseFloat(panjang, 64)
		l, _ := strconv.ParseFloat(lebar, 64)
		hasil := 2 * (p + l)
		result := gin.H{
			"hitung":  hitung,
			"panjang": panjang,
			"lebar":   lebar,
			"hasil":   hasil,
		}
		c.JSON(http.StatusOK, result)
	} else if hitung == "Luas" || hitung == "luas" {
		p, _ := strconv.ParseFloat(panjang, 64)
		l, _ := strconv.ParseFloat(lebar, 64)
		hasil := p * l
		result := gin.H{
			"hitung":  hitung,
			"panjang": panjang,
			"lebar":   lebar,
			"hasil":   hasil,
		}
		c.JSON(http.StatusOK, result)
	} else {
		//tidak diizinkan
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operasi tidak diizinkan"})
	}
}

//===================================================================

func Lingkaran(c *gin.Context, hitung string, jariJari string) {
	if hitung == "Keliling" || hitung == "keliling" {
		r, _ := strconv.ParseFloat(jariJari, 64)
		hasil := 2 * math.Pi * r
		result := gin.H{
			"hitung":   hitung,
			"jariJari": jariJari,
			"hasil":    hasil,
		}
		c.JSON(http.StatusOK, result)
	} else if hitung == "Luas" || hitung == "luas" {
		r, _ := strconv.ParseFloat(jariJari, 64)
		hasil := math.Pi * r * r
		result := gin.H{
			"hitung":   hitung,
			"jariJari": jariJari,
			"hasil":    hasil,
		}
		c.JSON(http.StatusOK, result)
	} else {
		//tidak diizinkan
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operasi tidak diizinkan"})
	}
}

//===================================================================
