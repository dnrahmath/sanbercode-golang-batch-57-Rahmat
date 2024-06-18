package main

import (
	"fmt"
)

// ===================================================================================================
// soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func periksaBuah(val bool) string {
	if val {
		return "Ada"
	} else {
		return "Tidak"
	}
}

func Soal1() {
	var buah1 = buah{"Nanas", "Kuning", false, 9000}
	var buah2 = buah{"Jeruk", "Oranye", true, 8000}
	var buah3 = buah{"Semangka", "Hijau & Merah", true, 10000}
	var buah4 = buah{"Pisang", "Kuning", false, 5000}

	fmt.Println("===================================================================")
	fmt.Println("Nama\t\tWarna\t\tAda Bijinya\tHarga")
	fmt.Println("===================================================================")
	fmt.Printf("%s\t\t%s\t\t%s\t\t%d\n", buah1.nama, buah1.warna, periksaBuah(buah1.adaBijinya), buah1.harga)
	fmt.Printf("%s\t\t%s\t\t%s\t\t%d\n", buah2.nama, buah2.warna, periksaBuah(buah2.adaBijinya), buah2.harga)
	fmt.Printf("%s\t%s\t%s\t\t%d\n", buah3.nama, buah3.warna, periksaBuah(buah3.adaBijinya), buah3.harga)
	fmt.Printf("%s\t\t%s\t\t%s\t\t%d\n", buah4.nama, buah4.warna, periksaBuah(buah4.adaBijinya), buah4.harga)
	fmt.Println("===================================================================")
}

// ===================================================================================================
// soal 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

//---

type persegi struct {
	sisi int
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

//---

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

//---

func Soal2() {
	var segitiga1 = segitiga{alas: 5, tinggi: 3}
	var persegi1 = persegi{sisi: 4}
	var persegiPanjang1 = persegiPanjang{panjang: 6, lebar: 2}

	fmt.Println("Luas segitiga:", segitiga1.luas())
	fmt.Println("Luas persegi:", persegi1.luas())
	fmt.Println("Luas persegi panjang:", persegiPanjang1.luas())
}

// ===================================================================================================
// soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) tambahWarna(warna string) {
	p.colors = append(p.colors, warna)
}

func Soal3() {
	p := phone{
		name:  "iPhone",
		brand: "Apple",
		year:  2023,
		colors: []string{
			"Silver",
			"Space Gray",
		},
	}
	p.tambahWarna("Gold")

	fmt.Println("Nama:", p.name)
	fmt.Println("Brand:", p.brand)
	fmt.Println("Tahun:", p.year)
	fmt.Println("Warna:", p.colors)
}

// ===================================================================================================
// soal 4
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

var dataFilm []movie

func tambahDataFilm(film movie, dataFilmPointer *[]movie) {
	*dataFilmPointer = append(*dataFilmPointer, film)
}

func Soal4() {
	tambahDataFilm(movie{
		title:    "LOTR",
		duration: 120, // dalam menit
		genre:    "action",
		year:     1999,
	}, &dataFilm)

	tambahDataFilm(movie{
		title:    "Avenger",
		duration: 120,
		genre:    "action",
		year:     2019,
	}, &dataFilm)

	tambahDataFilm(movie{
		title:    "Spiderman",
		duration: 120,
		genre:    "action",
		year:     2004,
	}, &dataFilm)

	tambahDataFilm(movie{
		title:    "Juon",
		duration: 120,
		genre:    "horror",
		year:     2004,
	}, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1, film.title)
		fmt.Printf("   Duration: %d minutes\n", film.duration)
		fmt.Printf("   Genre: %s\n", film.genre)
		fmt.Printf("   Year: %d\n", film.year)
		fmt.Println()
	}
}

// ===================================================================================================
func main() {
	fmt.Println("--Soal 1--------------------------------------")
	Soal1()

	fmt.Println("--Soal 2--------------------------------------")
	Soal2()

	fmt.Println("--Soal 3--------------------------------------")
	Soal3()

	fmt.Println("--Soal 4--------------------------------------")
	Soal4()
}
