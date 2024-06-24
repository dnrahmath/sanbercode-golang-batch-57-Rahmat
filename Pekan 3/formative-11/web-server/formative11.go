package soal4

import (
	"fmt"
	"math"
	"net/http"
)

func hitung(w http.ResponseWriter, r *http.Request) {
	jariJari := 7.0
	tinggi := 10.0

	// Hitung
	volume := math.Pi * math.Pow(jariJari, 2) * tinggi
	luasAlas := math.Pi * math.Pow(jariJari, 2)
	kelilingAlas := 2 * math.Pi * jariJari

	fmt.Fprintf(w, "jariJari: %.2f, tinggi: %.2f, volume: %.2f, luas alas: %.2f, keliling alas: %.2f",
		jariJari, tinggi, volume, luasAlas, kelilingAlas)
}

func FuncSoal4() {
	http.HandleFunc("/index", hitung)

	fmt.Println("starting web server at http://localhost:8080/index")

	http.ListenAndServe(":8080", nil)
}
