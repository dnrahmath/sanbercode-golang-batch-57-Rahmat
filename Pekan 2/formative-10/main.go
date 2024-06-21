package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// ===================================================================================================
// soal 1
func Mess(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

func Soal1() {
	kalimat := "Golang Backend Development"
	tahun := 2021
	defer Mess(kalimat, tahun)
	fmt.Println("Program berjalan...")
}

// ===================================================================================================
// soal 2
func kelilingSegitigaSamaSisi(sisi int, outMess bool) (string, error) {
	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if !outMess {
			panic(err)
		}
		return "", err
	}

	keliling := 3 * sisi
	if outMess {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}

	return fmt.Sprintf("%d", keliling), nil
}

func Soal2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

// ===================================================================================================
// soal 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}

func cetakAngka(angka *int) {
	fmt.Printf("Total angka: %d\n", *angka)
}

// ===================================================================================================
// soal 4
func addPhone(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}
func showPhonesNo4(phones *[]string) {
	sort.Strings(*phones)

	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}
func Soal4() {
	// Deklarasi variabel phones
	var phones = []string{}

	// Menambahkan data ke dalam variabel phones
	addPhone(&phones, "Xiaomi")
	addPhone(&phones, "Asus")
	addPhone(&phones, "IPhone")
	addPhone(&phones, "Samsung")
	addPhone(&phones, "Oppo")
	addPhone(&phones, "Realme")
	addPhone(&phones, "Vivo")

	// Menampilkan data phones dengan urutan dan jeda satu detik
	showPhonesNo4(&phones)
}

// ===================================================================================================
// soal 5
func showPhonesNo5(phones []string, wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}
func Soal5() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup
	wg.Add(1)
	go showPhonesNo5(phones, &wg)
	wg.Wait()
}

// ===================================================================================================
// soal 6
func getMovies(moviesChannel chan string, movies ...string) {
	defer close(moviesChannel)
	for _, movie := range movies {
		moviesChannel <- movie
	}
}
func Soal6() {
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	for value := range moviesChannel {
		fmt.Println(value)
	}
}

// ===================================================================================================
func main() {
	fmt.Println("--Soal 1--------------------------------------")
	Soal1()

	fmt.Println("--Soal 2--------------------------------------")
	Soal2()

	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
	defer fmt.Println("--Soal 3--------------------------------------")

	fmt.Println("--Soal 4--------------------------------------")
	Soal4()

	fmt.Println("--Soal 5--------------------------------------")
	Soal5()

	fmt.Println("--Soal 6--------------------------------------")
	Soal6()
}
