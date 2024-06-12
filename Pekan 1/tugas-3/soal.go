package main

import (
	"fmt"
	"strconv"
)

func Soal1() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	intPanjangPersegiPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
	intlebarPersegiPanjang, _ := strconv.Atoi(lebarPersegiPanjang)
	luasPersegiPanjang = intPanjangPersegiPanjang * intlebarPersegiPanjang
	kelilingPersegiPanjang = (intPanjangPersegiPanjang + intlebarPersegiPanjang) * 2

	intalasSegitiga, _ := strconv.Atoi(alasSegitiga)
	inttinggiSegitiga, _ := strconv.Atoi(tinggiSegitiga)
	luasSegitiga = (intalasSegitiga * inttinggiSegitiga) / 2

	fmt.Println("--Soal 1--------------------------------------")
	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)
}

func Soal2() {
	var nilaiJohn = 80
	var nilaiDoe = 50

	bandingkan := func(x int) string {
		if x >= 80 {
			return "A"
		} else if x >= 70 && x < 80 {
			return "B"
		} else if x >= 60 && x < 70 {
			return "C"
		} else if x >= 50 && x < 60 {
			return "D"
		} else if x < 50 {
			return "E"
		}
		return ""
	}

	fmt.Println("--Soal 2--------------------------------------")
	fmt.Println("nilaiJohn : " + bandingkan(nilaiJohn))
	fmt.Println("nilaiDoe : " + bandingkan(nilaiDoe))
}

func Soal3() {
	var tanggal = 26
	var bulan = 9
	var tahun = 2000

	atur := func(x int) string {
		switch x {
		case 1:
			return "Januari"
		case 2:
			return "Februari"
		case 3:
			return "Maret"
		case 4:
			return "April"
		case 5:
			return "Mei"
		case 6:
			return "Juni"
		case 7:
			return "Juli"
		case 8:
			return "Agustus"
		case 9:
			return "September"
		case 10:
			return "Oktober"
		case 11:
			return "November"
		case 12:
			return "Desember"
		}
		return ""
	}

	strTanggal := strconv.Itoa(tanggal)
	strTahun := strconv.Itoa(tahun)

	fmt.Println("--Soal 3--------------------------------------")
	fmt.Println(strTanggal + " " + atur(bulan) + " " + strTahun)

	Soal4(tahun)
}

func Soal4(tahun int) {
	bandingkan := func(x int) string {
		if x >= 1944 && x < 1964 {
			return "Baby boome"
		} else if x >= 1965 && x < 1979 {
			return "Generasi X"
		} else if x >= 1980 && x < 1994 {
			return "Generasi Y"
		} else if x >= 1995 && x < 2015 {
			return "Generasi Z"
		}
		return "gdfgffgd"
	}
	fmt.Println("--Soal 4--------------------------------------")
	fmt.Println(bandingkan(tahun))

}

func main() {
	Soal1()
	Soal2()
	Soal3()
}
