package main

import (
	"fmt"
	"math"
)

// ===================================================================================================
// soal 1
var luasLingkaran float64
var kelilingLingkaran float64

func funcPointer(l *float64, k *float64, jj float64) {
	*l = math.Pi * jj * jj
	*k = 2 * math.Pi * jj
}

// ===================================================================================================
// soal 2
func introduce(pointer *string, nama, gender, pekerjaan, umur string) {
	if gender == "laki-laki" {
		*pointer = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", nama, pekerjaan, umur)
	} else if gender == "perempuan" {
		*pointer = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", nama, pekerjaan, umur)
	}
}

// ===================================================================================================
// soal 3
func addBuah(buah *[]string, nmBuah string) {
	*buah = append(*buah, nmBuah)
}

// ===================================================================================================
// soal 4
var dataFilm = []map[string]string{}

func tambahDataFilm(title, duration, genre, year string, dataFilmPointer *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilmPointer = append(*dataFilmPointer, film)
}

// ===================================================================================================
func main() {
	var jariJari float64 = 10.0

	fmt.Println("--Soal 1--------------------------------------")
	funcPointer(&luasLingkaran, &kelilingLingkaran, jariJari)
	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	fmt.Println("--Soal 2--------------------------------------")
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	fmt.Println("--Soal 3--------------------------------------")
	var buah = []string{}
	addBuah(&buah, "Jeruk")
	addBuah(&buah, "Semangka")
	addBuah(&buah, "Mangga")
	addBuah(&buah, "Strawberry")
	addBuah(&buah, "Durian")
	addBuah(&buah, "Manggis")
	addBuah(&buah, "Alpukat")
	for i, v := range buah {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	fmt.Println("--Soal 4--------------------------------------")
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, film["title"])
		fmt.Printf("   duration : %s\n", film["duration"])
		fmt.Printf("   genre : %s\n", film["genre"])
		fmt.Printf("   year : %s\n", film["year"])
		fmt.Println()
	}
}
