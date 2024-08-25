package main

import (
	"fmt"
)

type mahasiswa struct {
	nama    string
	umur    int
	kelamin string
}

func main() {
	var totalMahasiswa int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&totalMahasiswa)

	mahasiswaList := inputMahasiswa(totalMahasiswa)
	listMahasiswa(mahasiswaList)

}

func listMahasiswa(mahasiswaList []mahasiswa) {
	fmt.Println("------------------------------------------------")
	fmt.Printf("%-15s %-15s %-15s\n", "Nama", "Umur", "Jenis Kelamin")
	fmt.Println("------------------------------------------------")
	for _, m := range mahasiswaList {
		fmt.Printf("%-15s %-15d %-15s\n", m.nama, m.umur, m.kelamin)
	}
	fmt.Println("------------------------------------------------")
}

func inputMahasiswa(totalMahasiswa int) []mahasiswa {
	var mahasiswaList []mahasiswa
	for i := 0; i < totalMahasiswa; i++ {
		var m mahasiswa
		fmt.Println("Data Mahasiswa ke-", i+1)
		fmt.Print("Masukkan nama mahasiswa: ")
		fmt.Scan(&m.nama)
		fmt.Print("Masukkan umur mahasiswa: ")
		fmt.Scan(&m.umur)
		// fmt.Print("Masukkan kelamin mahasiswa (L/P): ")
		m.kelamin = inputJenisKelamin()

		mahasiswaList = append(mahasiswaList, m)
	}

	return mahasiswaList
}

func inputJenisKelamin() string {
	var kelamin string
	for {
		fmt.Print("Masukkan jenis kelamin (L/P): ")
		fmt.Scan(&kelamin)
		if kelamin == "l" || kelamin == "L" || kelamin == "p" || kelamin == "P" {
			break
		} else {
			fmt.Println("Jenis kelamin tidak valid. Silakan masukkan L atau P.")
		}
	}
	if kelamin == "l" || kelamin == "L" {
		return "Laki-laki"
	} else {
		return "Perempuan"
	}
}
