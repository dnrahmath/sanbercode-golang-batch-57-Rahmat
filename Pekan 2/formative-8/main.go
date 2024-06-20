package main

import (
	"fmt"
	"math"
	"strings"
)

// ===================================================================================================
// soal 1
type segitigaSamaSisi struct{ alas, tinggi int }

type persegiPanjang struct{ panjang, lebar int }

type tabung struct{ jariJari, tinggi float64 }

type balok struct{ panjang, lebar, tinggi int }

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// ================ bangunDatar ================
// segitigaSamaSisi
func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

// persegiPanjang
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// ================ bangunRuang ================
// tabung
func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

// balok
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

// ================ ================

func Soal1() {
	// ================ bangunDatar ================
	var bangunDatar hitungBangunDatar

	segitiga := segitigaSamaSisi{alas: 10, tinggi: 7}
	persegi := persegiPanjang{panjang: 8, lebar: 5}

	// segitigaSamaSisi
	bangunDatar = segitiga
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Println("Luas:", bangunDatar.luas())
	fmt.Println("Keliling:", bangunDatar.keliling())

	// persegiPanjang
	bangunDatar = persegi
	fmt.Println("\nPersegi Panjang:")
	fmt.Println("Luas:", bangunDatar.luas())
	fmt.Println("Keliling:", bangunDatar.keliling())

	// ================ bangunRuang ================
	var bangunRuang hitungBangunRuang

	tbg := tabung{jariJari: 5, tinggi: 10}
	blk := balok{panjang: 4, lebar: 3, tinggi: 2}

	// tabung
	bangunRuang = tbg
	fmt.Println("\nTabung:")
	fmt.Println("Volume:", bangunRuang.volume())
	fmt.Println("Luas Permukaan:", bangunRuang.luasPermukaan())

	// balok
	bangunRuang = blk
	fmt.Println("\nBalok:")
	fmt.Println("Volume:", bangunRuang.volume())
	fmt.Println("Luas Permukaan:", bangunRuang.luasPermukaan())
}

// ===================================================================================================
// soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type DevInfo interface {
	OutputInf() string
}

func (p phone) OutputInf() string {
	colors := strings.Join(p.colors, ", ")
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.name, p.brand, p.year, colors)
}

func Soal2() {
	myPhone := phone{
		name:  "Samsung Galaxy Note 20",
		brand: "Samsung",
		year:  2020,
		colors: []string{
			"Mystic Bronze",
			"Mystic White",
			"Mystic Black",
		},
	}

	var devinfo DevInfo = myPhone
	fmt.Println(devinfo.OutputInf())
}

// ===================================================================================================
// soal 3
func luasPersegi(inpuAngkaSisi int, outputString bool) interface{} {
	if inpuAngkaSisi == 0 {
		if outputString {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}
	luas := inpuAngkaSisi * inpuAngkaSisi
	if outputString {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", inpuAngkaSisi, luas)
	} else {
		return luas
	}
}
func Soal3() {
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
}

// ===================================================================================================
// soal 4
func Soal4() {
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	// Type Assertion
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	// Menggabungkan kedua slice
	semuaAngka := append(angkaPertama, angkaKedua...)

	// Menjumlahkan semua angka dan membentuk string dari penjumlahan
	total := 0
	var angkaString []string
	for _, angka := range semuaAngka {
		total += angka
		angkaString = append(angkaString, fmt.Sprintf("%d", angka))
	}

	hasil := fmt.Sprintf(
		"%s%s = %d",
		prefix.(string),
		strings.Join(angkaString, " + "),
		total,
	)
	fmt.Println(hasil)
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
