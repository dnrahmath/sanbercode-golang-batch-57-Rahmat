package main

import (
	"fmt"
	"strconv"
)

// soal 1
func Soal1() {
	fmt.Println("--Soal 1--------------------------------------")

	for i := 0; i <= 20; i++ {
		strngNumber := strconv.Itoa(i)

		if i%3 == 0 && i%2 != 0 {
			fmt.Println(strngNumber + " - " + "I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(strngNumber + " - " + "Berkualitas")
		} else {
			fmt.Println(strngNumber + " - " + "Santai")
		}

	}
}

// soal 2
func Soal2() {
	fmt.Println("--Soal 2--------------------------------------")
	for vertical := 0; vertical < 7; vertical++ {
		for horizontal := 0; horizontal < vertical; horizontal++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

// soal 3
func Soal3() {
	fmt.Println("--Soal 3--------------------------------------")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:7])
}

// soal 4
func Soal4() {
	fmt.Println("--Soal 4--------------------------------------")
	var sayuran = []string{}

	var appendInput = []string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}

	for i := 0; i < len(appendInput); i++ {
		sayuran = append(sayuran, appendInput[i])
	}

	for i := 0; i < len(appendInput); i++ {
		strngNumber := strconv.Itoa(i)
		fmt.Println(strngNumber + ". " + sayuran[i])
	}

}

// soal 5
func Soal5() {
	fmt.Println("--Soal 5--------------------------------------")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	keysSatuan := make([]string, len(satuan))
	i := 0
	for k := range satuan {
		keysSatuan[i] = k
		i++
	}

	for i := 0; i < len(keysSatuan); i++ {
		strngValue := strconv.Itoa(satuan[keysSatuan[i]])
		fmt.Println(keysSatuan[i] + " = " + strngValue)
	}

	var value int = 1
	for i := 0; i < len(keysSatuan); i++ {
		value = value * satuan[keysSatuan[i]]
	}
	strngValue := strconv.Itoa(value)
	fmt.Println("Volume Balok = " + strngValue)
}

func main() {
	Soal1()
	Soal2()
	Soal3()
	Soal4()
	Soal5()
}
