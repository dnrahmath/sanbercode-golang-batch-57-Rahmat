package main

import (
	"fmt"
)

// ===================================================================================================
// soal 1
func luasPersegiPanjang(p int, l int) int {
	value := p * l
	return value
}

func kelilingPersegiPanjang(p int, l int) int {
	value := (p + l) * 2
	return value
}

func volumeBalok(p int, l int, t int) int {
	value := p * l * t
	return value
}

func Soal1() {
	fmt.Println("--Soal 1--------------------------------------")

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
}

// ===================================================================================================
// soal 2
func introduce(nm string, jk string, prof string, umr string) string {
	var panggilan string
	if jk == "laki-laki" {
		panggilan = "Pak"
	} else if jk == "perempuan" {
		panggilan = "Bu"
	} else {
		panggilan = ""
	}
	return panggilan + " " + nm + " adalah seorang " + prof + " yang berusia " + umr + " tahun"
}

func Soal2() {
	fmt.Println("--Soal 2--------------------------------------")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

// ===================================================================================================
// soal 3
func buahFavorit(nmPerson string, buah ...string) string {
	var listBuah string
	for i := 0; i < len(buah); i++ {
		if i == len(buah)-1 {
			listBuah = listBuah + buah[i]
		} else {
			listBuah = listBuah + buah[i] + ", "
		}
	}

	return "halo nama saya " + nmPerson + " dan buah favorit saya adalah " + listBuah
}
func Soal3() {
	fmt.Println("--Soal 3--------------------------------------")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
}

// ===================================================================================================
// soal 4
var dataFilm = []map[string]string{}

func tambahDataFilm(genre string, jam string, tahun string, title string) []map[string]string {
	film := map[string]string{
		"genre": genre,
		"jam":   jam,
		"tahun": tahun,
		"title": title,
	}

	dataFilm = append(dataFilm, film)

	return dataFilm
}

func Soal4() {
	fmt.Println("--Soal 4--------------------------------------")
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func main() {
	Soal1()
	Soal2()
	Soal3()
	Soal4()
}
