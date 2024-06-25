package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// ===================================================================================================

// NilaiMahasiswa struct
type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	IndeksNilai string `json:"indexnilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var (
	nilaiMhsList []NilaiMahasiswa
	mutex        sync.Mutex
	idCounter    uint = 1
)

// ===================================================================================================

// Function untuk menentukan IndeksNilai berdasarkan Nilai
func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

// ===================================================================================================

// GetNilaiMhs handler
func getNilaiMhs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	mutex.Lock()
	dataNilaiMhs, err := json.Marshal(nilaiMhsList)
	mutex.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataNilaiMhs)
}

// =================================================

// PostNilaiMhs handler
func postNilaiMhs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newNilai NilaiMahasiswa

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newNilai); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newNilai.Nama == "" || newNilai.MataKuliah == "" || newNilai.Nilai > 100 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	newNilai.ID = idCounter
	idCounter++

	newNilai.IndeksNilai = getIndeksNilai(newNilai.Nilai)

	mutex.Lock()
	nilaiMhsList = append(nilaiMhsList, newNilai)
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNilai)
}

// ===================================================================================================

// Auth middleware
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username atau Password tidak sesuai"))
	})
}

// ===================================================================================================

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getNilaiMhs(w, r)
	} else if r.Method == http.MethodPost {
		postNilaiMhs(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "POST" {
	// 	}
	// })
	http.Handle("/", auth(http.HandlerFunc(mainHandler)))

	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
